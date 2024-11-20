package repositories

import (
	"context"
	"employee-service/models"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
)

type EmployeeRepositoryInterface interface {
	AddEmployee(ctx context.Context, employee models.Employee) (int32, error)
	DeleteEmployee(ctx context.Context, id int32) error
	ShowCompanyEmployees(ctx context.Context, companyId int32, department models.Department) ([]models.Employee, error)
	UpdateEmployee(ctx context.Context, employee models.Employee) error
}

type EmployeeRepository struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func departmentExists(ctx context.Context, tx pgx.Tx, department models.Department) (bool, error) {
	var departExists bool

	err := tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM departments WHERE name = $1 AND phone = $2)",
		department.Name, department.Phone).Scan(&departExists)
	if err != nil {
		err = fmt.Errorf("employee_repo: department exists: query row department: %w", err)
		return false, err
	}
	return departExists, nil
}

func createOrGetDepartmentId(ctx context.Context, tx pgx.Tx, departExists bool, department models.Department) (int32, error) {
	var departId int32
	if !departExists {
		err := tx.QueryRow(ctx, "INSERT INTO departments (name, phone) VALUES ($1, $2) RETURNING id",
			department.Name, department.Phone).Scan(&departId)
		if err != nil {
			err = fmt.Errorf("employee_repo: department exists: insert department: %w", err)
			return 0, err
		}
	} else {
		err := tx.QueryRow(ctx, "SELECT id FROM departments WHERE name = $1 AND phone = $2",
			department.Name, department.Phone).Scan(&departId)
		if err != nil {
			err = fmt.Errorf("department exists: query row department: %w", err)
			return 0, err
		}
	}
	return departId, nil
}

func (r *EmployeeRepository) AddEmployee(ctx context.Context, employee models.Employee) (int32, error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		err = fmt.Errorf("employee_repo: add_employee: acquire connection: %w", err)
		return 0, err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("employee_repo: add_employee: begin transaction: %w", err)
		return 0, err
	}
	defer tx.Rollback(ctx)

	var departmentId int32
	departExists, err := departmentExists(ctx, tx, employee.Department)
	if err != nil {
		return 0, fmt.Errorf("employee_repo: add_employee: %w", err)
	}

	departmentId, err = createOrGetDepartmentId(ctx, tx, departExists, employee.Department)
	if err != nil {
		return 0, fmt.Errorf("employee_repo: add_employee: %w", err)
	}

	var passportId int32

	err = tx.QueryRow(ctx, "INSERT INTO passports (type, number) VALUES ($1, $2) RETURNING id",
		employee.Passport.Type, employee.Passport.Number).Scan(&passportId)
	if err != nil {
		err = fmt.Errorf("employee_repo: add_employee: insert passport: %w", err)
		return 0, err
	}

	insertQuery := `
		INSERT INTO employees (name, surname, phone, company_id, passport_id, department_id) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	err = tx.QueryRow(ctx, insertQuery, employee.Name, employee.Surname, employee.Phone, employee.CompanyId,
		passportId, departmentId).Scan(&passportId)
	if err != nil {
		err = fmt.Errorf("employee_repo: add_employee: insert employee: %w", err)
		return 0, err
	}

	if err = tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("employee_repo: add_employee: commit transaction: %w", err)
	}

	return passportId, nil
}

func (r *EmployeeRepository) DeleteEmployee(ctx context.Context, id int32) error {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		err = fmt.Errorf("employee_repo: delete_employee: acquire connection: %w", err)
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("employee_repo: delete_employee: begin transaction: %w", err)
		return err
	}
	defer tx.Rollback(ctx)

	var passportId, departmentId int32

	err = tx.QueryRow(ctx, "SELECT passport_id, department_id FROM employees WHERE id = $1", id).
		Scan(&passportId, &departmentId)
	if err != nil {
		err = fmt.Errorf("employee_repo: delete_employee: pass and department id: %w", err)
		return err
	}

	_, err = tx.Exec(ctx, "DELETE FROM employees WHERE id = $1", id)
	if err != nil {
		err = fmt.Errorf("employee_repo: delete_employee: delete employee: %w", err)
		return err
	}

	_, err = tx.Exec(ctx, "DELETE FROM passports WHERE id = $1", passportId)
	if err != nil {
		err = fmt.Errorf("employee_repo: delete_employee: delete passport: %w", err)
		return err
	}

	var leftDepartment bool
	err = tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM employees WHERE department_id = $1)", departmentId).
		Scan(&leftDepartment)
	if err != nil {
		err = fmt.Errorf("employee_repo: delete_employee: query row department: %w", err)
		return err
	}

	if !leftDepartment {
		_, err = tx.Exec(ctx, "DELETE FROM departments WHERE id = $1", departmentId)
		if err != nil {
			err = fmt.Errorf("employee_repo: delete_employee: delete department: %w", err)
			return err
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("employee_repo: delete_employee: commit error: %w", err)
	}

	return nil
}

func (r *EmployeeRepository) ShowCompanyEmployees(ctx context.Context, companyId int32, department models.Department) ([]models.Employee, error) {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		err = fmt.Errorf("employee_repo: show_department_employee: acquire connection: %w", err)
		return nil, err
	}
	defer conn.Release()

	query := `
		SELECT e.id, e.name, e.surname, e.phone, e.company_id,
		       p.type, p.number,
		       d.name, d.phone
		FROM employees AS e
		JOIN departments AS d ON e.department_id = d.id AND e.company_id = $1 %s
		JOIN passports AS p ON e.passport_id = p.id`

	args := []interface{}{}
	args = append(args, companyId)

	switch {
	case department.Name == "" && department.Phone == "":
		query = fmt.Sprintf(query, "")
	case department.Name != "" && department.Phone == "":
		query = fmt.Sprintf(query, "AND d.name = $2")
		args = append(args, department.Name)
	case department.Name == "" && department.Phone != "":
		query = fmt.Sprintf(query, "AND d.phone = $2")
		args = append(args, department.Phone)
	default:
		query = fmt.Sprintf(query, "AND d.name = $2 AND d.phone = $3")
		args = append(args, department.Name)
		args = append(args, department.Phone)
	}

	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		err = fmt.Errorf("employee_repo: show_department_employee: query: %w", err)
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var employee models.Employee
		var passport models.Passport
		var department2 models.Department
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Surname, &employee.Phone, &employee.CompanyId,
			&passport.Type, &passport.Number, &department2.Name, &department2.Phone)

		if err != nil {
			err = fmt.Errorf("employee_repo: show_department_employee: scan: %w", err)
			return nil, err
		}

		employee.Passport = passport
		employee.Department = department2

		employees = append(employees, employee)
	}

	return employees, nil
}

func cleanUnusedDepartments(ctx context.Context, tx pgx.Tx, departmentId int32) error {
	var countOldDepartment int
	err := tx.QueryRow(ctx, "SELECT COUNT(*) FROM employees WHERE department_id = $1", departmentId).
		Scan(&countOldDepartment)
	if err != nil {
		err = fmt.Errorf("clean unused depart: select department: %w", err)
		return err
	}
	if countOldDepartment == 0 {
		_, err = tx.Exec(ctx, "DELETE FROM departments WHERE id = $1", departmentId)
		if err != nil {
			err = fmt.Errorf("clean unused depart: delete department: %w", err)
			return err
		}
	}
	return nil
}

func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, employee models.Employee) error {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		err = fmt.Errorf("employee_repo: update_employee: acquire connection: %w", err)
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("employee_repo: update_employee: begin transaction: %w", err)
		return err
	}
	defer tx.Rollback(ctx)

	var departmentId int32

	err = tx.QueryRow(ctx, "SELECT department_id FROM employees WHERE id = $1", employee.Id).
		Scan(&departmentId)
	if err != nil {
		err = fmt.Errorf("employee_repo: update_employee: pass and depart ids query: %w", err)
		return err
	}

	if employee.Name != "" || employee.Surname != "" || employee.Phone != "" || employee.CompanyId != 0 {
		err = updateEmployeeData(ctx, tx, employee)
		if err != nil {
			return fmt.Errorf("employee repo: update employee: update empl data: %w", err)
		}
	}

	if employee.Passport.Type != "" || employee.Passport.Number != "" {
		err = updatePassport(ctx, tx, employee)
		if err != nil {
			return fmt.Errorf("employee repo: update employee: update pass data: %w", err)
		}
	}

	if employee.Department.Phone != "" || employee.Department.Name != "" {
		newDepartmentId, err := updateDepartment(ctx, tx, employee.Id, departmentId, employee.Department)
		if err != nil {
			return fmt.Errorf("employee repo: update employee: update depart data: %w", err)
		}

		_, err = tx.Exec(ctx, "UPDATE employees SET department_id = $1 WHERE id = $2", newDepartmentId, employee.Id)
		if err != nil {
			err = fmt.Errorf("update employee: update department id: %w", err)
			return err
		}

		err = cleanUnusedDepartments(ctx, tx, departmentId)
		if err != nil {
			err = fmt.Errorf("update employee: clean unused department: %w", err)
			return err
		}

	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("employee_repo: update_employee: commit transaction: %w", err)
	}

	return nil
}

func updateEmployeeData(ctx context.Context, tx pgx.Tx, employee models.Employee) error {
	updateEmployeeQuery := "UPDATE employees SET"
	employeeArgs := []interface{}{}
	fields := make([]string, 0)
	index := 1

	if employee.Name != "" {
		fields = append(fields, fmt.Sprintf("name = $%v", index))
		employeeArgs = append(employeeArgs, employee.Name)
		index++
	}
	if employee.Surname != "" {
		fields = append(fields, fmt.Sprintf("surname = $%v", index))
		employeeArgs = append(employeeArgs, employee.Surname)
		index++
	}
	if employee.Phone != "" {
		fields = append(fields, fmt.Sprintf("phone = $%v", index))
		employeeArgs = append(employeeArgs, employee.Phone)
		index++
	}
	if employee.CompanyId != 0 {
		fields = append(fields, fmt.Sprintf("company_id = $%v", index))
		employeeArgs = append(employeeArgs, employee.CompanyId)
		index++
	}
	updateEmployeeQuery += " " + strings.Join(fields, ", ") + fmt.Sprintf(" WHERE id = $%v", index)
	employeeArgs = append(employeeArgs, employee.Id)

	_, err := tx.Exec(ctx, updateEmployeeQuery, employeeArgs...)
	if err != nil {
		return fmt.Errorf("update employee data: %w", err)
	}
	return nil
}

func updatePassport(ctx context.Context, tx pgx.Tx, employee models.Employee) error {
	updatePassportQuery := "UPDATE passports SET"
	passportArgs := []interface{}{}
	fields := make([]string, 0)
	index := 1

	if employee.Passport.Type != "" {
		fields = append(fields, fmt.Sprintf("type = $%v", index))
		passportArgs = append(passportArgs, employee.Passport.Type)
		index++
	}
	if employee.Passport.Number != "" {
		fields = append(fields, fmt.Sprintf("number = $%v", index))
		passportArgs = append(passportArgs, employee.Passport.Number)
		index++
	}
	updatePassportQuery += " " + strings.Join(fields, ", ") + fmt.Sprintf(" WHERE id = $%v", index)
	passportArgs = append(passportArgs, employee.Id)

	_, err := tx.Exec(ctx, updatePassportQuery, passportArgs...)
	if err != nil {
		return fmt.Errorf("employee_repo: update_employee: update passport: %w", err)
	}
	return nil
}

func updateDepartment(ctx context.Context, tx pgx.Tx, id, departmentId int32, department models.Department) (int32, error) {
	var newDepartmentId int32
	switch {
	case department.Name != "" && department.Phone != "":
		departExists, err := departmentExists(ctx, tx, department)
		if err != nil {
			err = fmt.Errorf("update_department: check department exists: %w", err)
			return 0, err
		}

		newDepartmentId, err = createOrGetDepartmentId(ctx, tx, departExists, department)
		if err != nil {
			err = fmt.Errorf("update_department: create department: %w", err)
			return 0, err
		}

		_, err = tx.Exec(ctx, "UPDATE employees SET department_id = $1 WHERE id = $2", newDepartmentId, id)
		if err != nil {
			err = fmt.Errorf("update_department: update department id: %w", err)
			return 0, err
		}
		return newDepartmentId, nil

	case department.Name != "":
		var departmentPhone string
		err := tx.QueryRow(ctx, "SELECT phone FROM departments WHERE id = $1", departmentId).Scan(&departmentPhone)
		if err != nil {
			err = fmt.Errorf("update_department: select department: %w", err)
			return 0, err
		}

		departExists, err := departmentExists(ctx, tx, models.Department{Name: department.Name, Phone: departmentPhone})
		if err != nil {
			err = fmt.Errorf("update_department: check department exists: %w", err)
			return 0, err
		}

		newDepartmentId, err = createOrGetDepartmentId(ctx, tx, departExists, models.Department{Name: department.Name, Phone: departmentPhone})
		if err != nil {
			err = fmt.Errorf("update_department: create department: %w", err)
			return 0, err
		}
		return newDepartmentId, nil

	case department.Phone != "":
		var departmentName string
		err := tx.QueryRow(ctx, "SELECT name FROM departments WHERE id = $1", departmentId).Scan(&departmentName)
		if err != nil {
			err = fmt.Errorf("update_department: select department: %w", err)
			return 0, err
		}

		departExists, err := departmentExists(ctx, tx, models.Department{Name: departmentName, Phone: department.Phone})
		if err != nil {
			err = fmt.Errorf("update_department: check department exists: %w", err)
			return 0, err
		}

		newDepartmentId, err = createOrGetDepartmentId(ctx, tx, departExists, models.Department{Name: departmentName, Phone: department.Phone})
		if err != nil {
			err = fmt.Errorf("update_department: create department: %w", err)
			return 0, err
		}
		return newDepartmentId, nil
	}
	return 0, fmt.Errorf("update_department: no one case")
}

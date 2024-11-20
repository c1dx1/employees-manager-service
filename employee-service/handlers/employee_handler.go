package handlers

import (
	"context"
	"employee-service/models"
	"employee-service/proto"
	"employee-service/repositories"
	"fmt"
	"log"
)

type EmployeeHandlerInterface interface {
	AddEmployee(ctx context.Context, req *proto.AddEmployeeRequest) (*proto.AddEmployeeResponse, error)
	DeleteEmployee(ctx context.Context, req *proto.DeleteEmployeeRequest) (*proto.DeleteEmployeeResponse, error)
	ShowCompanyEmployees(ctx context.Context, req *proto.CompanyEmployeesRequest) (*proto.EmployeesResponse, error)
	UpdateEmployee(ctx context.Context, req *proto.UpdateEmployeeRequest) (*proto.UpdateEmployeeResponse, error)
}

type EmployeeHandler struct {
	repo repositories.EmployeeRepository
	proto.UnimplementedEmployeeServiceServer
}

func NewEmployeeHandler(repo repositories.EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{repo: repo}
}

func (h *EmployeeHandler) AddEmployee(ctx context.Context, req *proto.AddEmployeeRequest) (*proto.AddEmployeeResponse, error) {
	employee := models.Employee{
		Name:      req.Name,
		Surname:   req.Surname,
		Phone:     req.Phone,
		CompanyId: req.CompanyId,
	}
	if req.Passport != nil {
		employee.Passport = models.Passport{
			Type:   req.Passport.Type,
			Number: req.Passport.Number,
		}
	} else {
		employee.Passport = models.Passport{
			Type:   "",
			Number: "",
		}
	}
	if req.Department == nil {
		employee.Department = models.Department{
			Name:  "",
			Phone: "",
		}
	} else {
		employee.Department = models.Department{
			Name:  req.Department.Name,
			Phone: req.Department.Phone,
		}
	}

	id, err := h.repo.AddEmployee(ctx, employee)
	if err != nil {
		log.Printf("employee_handler: repo add employee: %w", err)
		return nil, fmt.Errorf("employee_handler: repo add employee: %w", err)
	}

	return &proto.AddEmployeeResponse{Id: id}, nil
}

func (h *EmployeeHandler) DeleteEmployee(ctx context.Context, req *proto.DeleteEmployeeRequest) (*proto.DeleteEmployeeResponse, error) {
	if err := h.repo.DeleteEmployee(ctx, req.Id); err != nil {
		err = fmt.Errorf("employee_handler: repo delete employee: %w", err)
		log.Printf("%w", err)
		return &proto.DeleteEmployeeResponse{Success: "Fail"}, err
	}

	return &proto.DeleteEmployeeResponse{Success: "Success"}, nil
}

func (h *EmployeeHandler) ShowCompanyEmployees(ctx context.Context, req *proto.CompanyEmployeesRequest) (*proto.EmployeesResponse, error) {
	var department models.Department
	if req.Department == nil {
		department = models.Department{
			Name:  "",
			Phone: "",
		}
	}
	employees, err := h.repo.ShowCompanyEmployees(
		ctx, req.CompanyId, department)

	if err != nil {
		err = fmt.Errorf("employee_handler: repo show comp employees: %w", err)
		log.Printf("%w", err)
		return &proto.EmployeesResponse{}, err
	}

	var resp = &proto.EmployeesResponse{}
	for _, employee := range employees {
		protoEmployee := &proto.Employee{
			Id:        employee.Id,
			Name:      employee.Name,
			Surname:   employee.Surname,
			Phone:     employee.Phone,
			CompanyId: employee.CompanyId,
			Passport: &proto.Employee_Passport{
				Type:   employee.Passport.Type,
				Number: employee.Passport.Number,
			},
			Department: &proto.Employee_Department{
				Name:  employee.Department.Name,
				Phone: employee.Department.Phone,
			},
		}
		resp.Employees = append(resp.Employees, protoEmployee)
	}

	return resp, nil
}

func (h *EmployeeHandler) UpdateEmployee(ctx context.Context, req *proto.UpdateEmployeeRequest) (*proto.UpdateEmployeeResponse, error) {
	var employee models.Employee
	employee.Id = req.Id
	employee.Name = req.Name
	employee.Surname = req.Surname
	employee.Phone = req.Phone
	employee.CompanyId = req.CompanyId
	if req.Passport == nil {
		employee.Passport = models.Passport{
			Type:   "",
			Number: "",
		}
	} else {
		employee.Passport = models.Passport{
			Type:   req.Passport.Type,
			Number: req.Passport.Number,
		}
	}
	if req.Department == nil {
		employee.Department = models.Department{
			Name:  "",
			Phone: "",
		}
	} else {
		employee.Department = models.Department{
			Name:  req.Department.Name,
			Phone: req.Department.Phone,
		}
	}

	err := h.repo.UpdateEmployee(ctx, employee)
	if err != nil {
		err = fmt.Errorf("employee_handler: update empl:repo err: %w", err)
		log.Printf("%w", err)
		return &proto.UpdateEmployeeResponse{Success: "Fail"}, err
	}

	return &proto.UpdateEmployeeResponse{Success: "Success"}, nil
}

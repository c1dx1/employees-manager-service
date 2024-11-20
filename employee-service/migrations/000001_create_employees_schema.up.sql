CREATE TABLE departments
(
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(255),
    phone VARCHAR(50)
);
CREATE TABLE passports
(
    id     SERIAL PRIMARY KEY,
    type   VARCHAR(50),
    number VARCHAR(50)
);
CREATE TABLE employees
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(255),
    surname       VARCHAR(255),
    phone         VARCHAR(50),
    company_id    INT,
    department_id INT,
    passport_id   INT,
    FOREIGN KEY (department_id) REFERENCES departments (id),
    FOREIGN KEY (passport_id) REFERENCES passports (id)
);

INSERT INTO departments (name, phone)
VALUES ('HR Department', '123-456-789'),
       ('IT Department', '987-654-321'),
       ('Finance Department', '456-789-123'),
       ('Marketing Department', '321-654-987'),
       ('Sales Department', '654-987-321'),
       ('Operations Department', '789-123-456'),
       ('Legal Department', '159-753-486'),
       ('Customer Support', '852-741-963'),
       ('Engineering Department', '741-852-963'),
       ('Research and Development', '369-258-147');

INSERT INTO passports (type, number)
VALUES ('National', 'AB123456'),
       ('National', 'CD789012'),
       ('National', 'EF345678'),
       ('International', 'GH901234'),
       ('International', 'IJ567890'),
       ('National', 'KL123456'),
       ('International', 'MN789012'),
       ('National', 'OP345678'),
       ('International', 'QR901234'),
       ('National', 'ST567890');

INSERT INTO employees (name, surname, phone, company_id, department_id, passport_id)
VALUES ('John', 'Doe', '555-1234', 1, 1, 1),
       ('Jane', 'Smith', '555-5678', 1, 2, 2),
       ('Alice', 'Johnson', '555-9012', 2, 3, 3),
       ('Bob', 'Williams', '555-3456', 2, 4, 4),
       ('Charlie', 'Brown', '555-7890', 3, 5, 5),
       ('Diana', 'Prince', '555-2345', 3, 6, 6),
       ('Bruce', 'Wayne', '555-6789', 4, 7, 7),
       ('Clark', 'Kent', '555-3457', 4, 8, 8),
       ('Peter', 'Parker', '555-5679', 5, 9, 9),
       ('Tony', 'Stark', '555-7891', 5, 10, 10);
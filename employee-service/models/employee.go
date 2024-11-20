package models

type Passport struct {
	Type   string
	Number string
}

type Department struct {
	Name  string
	Phone string
}

type Employee struct {
	Id         int32
	Name       string
	Surname    string
	Phone      string
	CompanyId  int32
	Passport   Passport
	Department Department
}

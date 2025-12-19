package Company

import "fmt"

type Employee interface {
	GetDetails() string
	GetID() uint64
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint64
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %v, Name: %v, Salary: %v", f.ID, f.Name, f.Salary)
}

func (f FullTimeEmployee) GetID() uint64 {
	return f.ID
}

type PartTimeEmployee struct {
	ID        uint64
	Name      string
	HourlyPay uint64
	Hours     uint64
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %v, Name: %v, HourlyPay: %v, Hours: %v", p.ID, p.Name, p.HourlyPay, p.Hours)
}

func (p PartTimeEmployee) GetID() uint64 {
	return p.ID
}

type Company struct {
	NumForAutoIncrement uint64
	Employees           map[uint64]Employee
}

func NewCompany() *Company {
	return &Company{
		NumForAutoIncrement: 1,
		Employees:           make(map[uint64]Employee),
	}
}

func (c *Company) AddEmployee(employee Employee) {
	c.Employees[employee.GetID()] = employee
	fmt.Println("Employee successfuly added")
}

func (c *Company) ListEmployees() {
	if len(c.Employees) == 0 {
		fmt.Println("No employees at company")
		return
	}

	for _, employee := range c.Employees {
		fmt.Println(employee.GetDetails())
	}
}

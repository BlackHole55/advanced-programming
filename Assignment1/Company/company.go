package Company

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint64
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %v, Name: %v, Salary: %v", f.ID, f.Name, f.Salary)
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

type Company struct {
	Employees map[uint64]Employee
}

func NewCompany() *Company {
	return &Company{
		Employees: make(map[uint64]Employee),
	}
}

func (c *Company) AddEmployee(id uint64, employee Employee) {
	for ID := range c.Employees {
		if ID == id {
			fmt.Println("Employee with such ID is already in company")
			return
		}
	}

	c.Employees[id] = employee
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

package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}
type Contract struct {
	empId    int
	basicPay int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

// Salary of a permanent employee is the sum of the
// basic pay and and pf
func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

// Salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicPay
}

func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

/*
Total expense is calculated by iterating throught the
SalaryCalculator slice and summing the salaries of
the individual employees
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total expense per month $%d\n", expense)
}

func main() {
	pemp1 := Permanent{
		empId:    1,
		basicPay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empId:    2,
		basicPay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicPay: 3000,
	}
	femp1 := Freelancer{
		empId:       4,
		ratePerHour: 70,
		totalHours:  120,
	}
	femp2 := Freelancer{
		empId:       5,
		ratePerHour: 100,
		totalHours:  100,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1, femp1, femp2}
	totalExpense(employees)
}

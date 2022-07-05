package main

func main() {
	var employeeSalary map[string]int //  The zero value of a map is 'nil'
	employeeSalary["steve"] = 12000   // Error: Assignment to a 'nil' map
}

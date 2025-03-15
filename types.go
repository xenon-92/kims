package main

type employee struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Salary  float64 `json:"price"`
}

func getEmployees() []employee {
	employees := []employee{
		{ID: "1", Name: "John", Address: "New York", Salary: 1000},
		{ID: "2", Name: "Doe", Address: "Los Angeles", Salary: 1200},
		{ID: "3", Name: "Smith", Address: "San Francisco", Salary: 1400},
	}
	return employees
}

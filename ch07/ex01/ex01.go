package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

type Manager struct {
	Employee
	Reports []Employee
}

func doThings(i any) {
	switch i := i.(type) {
	case nil:
		fmt.Printf("i=%v (%T)\n", i, i)
	case string:
		fmt.Printf("i=%s (%T)\n", i, i)
	default:
		fmt.Printf("i=%v (%T)\n", i, i)
	}
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "my name",
			ID:   "my id",
		},
		Reports: []Employee{},
	}
	fmt.Println(m)

	doThings(123)
	doThings("abc")
	doThings(nil)
}

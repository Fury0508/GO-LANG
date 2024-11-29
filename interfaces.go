package main

import "fmt"

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

type Manager struct {
	Name string
}

func (M *Manager) GetName() string {
	return M.Name
}
func (e *Engineer) GetName() string {
	return e.Name
}
func PrintDetails(e Employee) {
	fmt.Println((e.GetName()))
}
func main() {
	engineer := &Engineer{Name: "Dhananjay"}
	manager := &Manager{Name: "Anshul"}
	PrintDetails(engineer)
	PrintDetails(manager)
}

package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project Project
}
type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func (e Engineer) Print() {
	fmt.Println("Engineer Information")
	fmt.Println(e.Name)
	fmt.Println(e.Age)
	fmt.Println(e.Project.Name)

}
func (e *Engineer) UpdateAge() {
	e.Age += 1
}

func (e *Engineer) GetProjectPriority() string {
	return e.Project.Priority
}

func main() {
	engineer := Engineer{
		Name: "Dhannajay",
		Age:  26,
		Project: Project{
			Name:         "Jobs",
			Priority:     "High",
			Technologies: []string{"go", "Python"},
		},
	}
	engineer.UpdateAge()
	engineer.Print()

	fmt.Println((engineer.GetProjectPriority()))
}

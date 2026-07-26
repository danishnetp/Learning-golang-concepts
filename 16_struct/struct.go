package main

import (
	"fmt"
	"strings"
)

// Define a struct type named Person with fields for name and age.
type Person struct {
	Name string
	Age  int
}

// Initialize a new Person instance with the provided name and age.
func NewPerson(name string, age int) Person {
	if age < 0 {
		age = 0
	}
	return Person{Name: name, Age: age}
}

// Create a method for the Person struct that returns a greeting message.
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

// Create a method for the Person struct that increments the person's age.
func (p *Person) HaveBirthday() {
	p.Age++
}

// Employee embeds Person and adds role-specific data.
type Employee struct {
	Person
	Role   string
	Skills []string
}

// Summary creates a compact line about the employee.
func (e Employee) Summary() string {
	return fmt.Sprintf("%s works as a %s (skills: %s)", e.Name, e.Role, strings.Join(e.Skills, ", "))
}

// Rectangle shows methods that do calculations from struct fields.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Address and Company demonstrate nested structs.
type Address struct {
	City    string
	Country string
}

type Company struct {
	Name    string
	Address Address
}

// main demonstrates the use of structs in Go.
func main() {
	// 1) Basic struct literal and constructor.
	p := Person{Name: "Alice", Age: 30}
	p2 := NewPerson("Bob", -2)
	fmt.Println("Person literal:", p)
	fmt.Println("Person constructor (age validation):", p2)

	// 2) Value receiver and pointer receiver.
	fmt.Println(p.Greet())
	p.HaveBirthday()
	fmt.Println("After birthday:", p)

	// 3) Embedded struct.
	e := Employee{
		Person: NewPerson("Carol", 28),
		Role:   "Backend Engineer",
		Skills: []string{"Go", "SQL", "Docker"},
	}
	fmt.Println(e.Summary())

	// 4) Struct with behavior methods.
	r := Rectangle{Width: 5.5, Height: 3}
	fmt.Printf("Rectangle area=%.2f perimeter=%.2f\n", r.Area(), r.Perimeter())

	// 5) Nested struct.
	c := Company{
		Name: "Acme Labs",
		Address: Address{
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}
	fmt.Printf("Company: %s (%s, %s)\n", c.Name, c.Address.City, c.Address.Country)

	// 6) Anonymous struct.
	config := struct {
		Port int
		Host string
	}{
		Port: 8080,
		Host: "localhost",
	}
	fmt.Printf("Anonymous config: %+v\n", config)

	// Embedded struct example with method call.
	emp := Employee{
		Person: NewPerson("Dave", 35),
		Role:   "DevOps Engineer",
		Skills: []string{"Kubernetes", "Terraform"},
	}
	fmt.Println(emp.Greet()) // Calls the embedded Person's Greet method.

	// Embedded struct with Employee and Address.
	company := Company{
		Name: "Tech Solutions",
		Address: Address{
			City:    "Bandung",
			Country: "Indonesia",
		},
	}
	fmt.Printf("Company: %s located in %s, %s\n", company.Name, company.Address.City, company.Address.Country) // Output: Company: Tech Solutions located in Bandung, Indonesia
}

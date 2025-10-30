package task2

import (
  "fmt"
  "testing"
)

func TestInterface(t *testing.T) {
  circle := Circle{}
  rectangle := Rectangle{}

  fmt.Println("circle area is ", circle.Area())
  fmt.Println("circle perimeter is ", circle.Perimeter())
  fmt.Println("rectangle area is ", rectangle.Area())
  fmt.Println("rectangle perimeter is ", rectangle.Perimeter())

}

func TestEmployee(t *testing.T) {
  p := Person{
    Name: "张三",
    Age:  13,
  }

  e := Employee{
    Person:     p,
    EmployeeId: "A101",
  }

  e.PrintInfo()
}

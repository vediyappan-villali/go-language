// Go program to understand multiple variable declaration

package main

import "fmt"

func main() {
    var name, age = "Vediyappan", 27
    salary, location := 1000.00, "Bangalore"

    fmt.Println("name = ", name)
    fmt.Println("age = ", age)
    fmt.Println("salary = ", salary)
    fmt.Println("location = ", location)
}

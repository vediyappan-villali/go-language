// Go program to declare variables with initial values

package main

import "fmt"

func main() {
    // data type is string
    var name string = "Vediyappan"
    // data type is inferred
    var age = 25
    // data type is inferred
    salary := 1000.50

    fmt.Println("name = ", name)
    fmt.Println("age = ", age)
    fmt.Println("salary = ", salary)
}

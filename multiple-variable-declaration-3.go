// Go program to understand variable declaration in a block

package main

import "fmt"

func main() {
    var (
        num1 int
        num2 int = 5
        name string = "Vediyappan"
    )

    fmt.Println("num1 = ", num1)
    fmt.Println("num2 = ", num2)
    fmt.Println("name = ", name)
}

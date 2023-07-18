// Go program to understand output function - Print

package main

import "fmt"

func main() {
    var str1, str2 string = "Vediyappan", "Villali"
    var num1, num2 = 10, 20

    // fmt.Print(str1, str2)
    // fmt.Print(str1, "\n", str2, "\n")
    fmt.Print(str1, " ", str2, "\n")

    // Print inserts a space between arguments if neither are strings
    // fmt.Print(num1, num2)
    fmt.Print(num1, num2, "\n")
}

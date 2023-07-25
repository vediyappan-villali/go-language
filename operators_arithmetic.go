// Go program to understand arithmetic operators

package main

import "fmt"

func main() {
     var num1 int = 10
     var num2 int = 5

     var sum int = num1 + num2
     var diff int = num1 - num2
     var prod int = num1 * num2
     var quot int = num1 / num2
     var rem int = num1 % num2

     fmt.Println("Sum = ", sum)
     fmt.Println("Difference = ", diff)
     fmt.Println("Product = ", prod)
     fmt.Println("Quotient = ", quot)
     fmt.Println("Remainder = ", rem)
     
     num1++;
     fmt.Println("Increment = ", num1)

     num1--
     fmt.Println("Decrement = ", num1)
}

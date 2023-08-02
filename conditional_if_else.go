// Go program to understand conditional if else

package main

import "fmt"

func main() {
    num1 := 15
    num2 := 10

    if num1 > num2 {
        fmt.Println(num1, "is greater than", num2)
    } else if num1 < num2 {
        fmt.Println(num1, "is lesser than", num2)
    } else {
        fmt.Println(num1, "is equal to", num2)
    }
}

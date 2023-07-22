// Go program to create slice using make()

package main

import "fmt"

func main() {
    var slc1= make([]int, 5, 10)

    fmt.Println(slc1)
    fmt.Println(len(slc1))
    fmt.Println(cap(slc1))
    
    var slc2= make([]int, 5, 5)

    fmt.Println(slc2)
    fmt.Println(len(slc2))
    fmt.Println(cap(slc2))
}

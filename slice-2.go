// Go program to understand length and capacity of slice.

package main

import "fmt"

func main() {
    var slc1 = []int{}

    fmt.Println(slc1)
    fmt.Println(len(slc1))
    fmt.Println(cap(slc1))

    slc2 := []string{"Hi", "hello", "how", "are", "you?"}

    fmt.Println(slc2)
    fmt.Println(len(slc2))
    fmt.Println(cap(slc2))
}

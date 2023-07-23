// Go program to append elements to slice

package main

import "fmt"

func main() {
    slc := []int{1, 2, 3, 4, 5, 6}

    fmt.Printf("slc = %v\n", slc)
    fmt.Printf("length = %v\n", len(slc))
    fmt.Printf("capacity = %v\n", cap(slc))

    slc = append(slc, 7, 8)

    fmt.Printf("slc = %v\n", slc)
    fmt.Printf("length = %v\n", len(slc))
    fmt.Printf("capacity = %v\n", cap(slc))
}

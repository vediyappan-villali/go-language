// Go program to create a slice from an array

package main

import "fmt"

func main() {
    var arr = []int{1, 2, 3, 4, 5, 6}
    var slc = arr[2:5]

    fmt.Println(len(slc))
    fmt.Println(cap(slc))
    fmt.Println(slc)
}

// Go program to understand array initialization

package main

import "fmt"

func main() {
    var arr1 = [3]int{}
    var arr2 = [3]int{1}
    var arr3 = [3]int{1, 2, 3}

    fmt.Println(arr1)
    fmt.Println(arr2)
    fmt.Println(arr3)
}

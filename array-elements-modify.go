// Go program to modify array elements

package main

import "fmt"

func main() {
    var arr = [3]int{10, 20, 30}

    fmt.Println(arr)

    arr[2] = 50

    fmt.Println(arr)
}

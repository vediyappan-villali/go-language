// Go program to modify elements in slice

package main

import "fmt"

func main() {
    prices := []int{10, 20, 30}

    fmt.Println(prices[0])
    fmt.Println(prices[2])

    prices[2] = 50

    fmt.Println(prices[2])
}

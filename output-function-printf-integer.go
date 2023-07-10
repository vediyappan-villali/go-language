// Go program to understand printf - integer formatting verbs

package main

import "fmt"

func main() {
    var num = 10

    fmt.Printf("%b\n", num)
    fmt.Printf("%d\n", num)
    fmt.Printf("%+d\n", num)
    fmt.Printf("%o\n", num)
    fmt.Printf("%O\n", num)
    fmt.Printf("%x\n", num)
    fmt.Printf("%X\n", num)
    fmt.Printf("%#X\n", num)
    fmt.Printf("%4d\n", num)
    fmt.Printf("%-4d\n", num)
    fmt.Printf("%04d\n", num)
}

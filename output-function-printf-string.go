// Go program to understand printf output function - string formatting

package main

import "fmt"

func main() {
    var message = "All is well!"

    fmt.Printf("%s\n", message)
    fmt.Printf("%q\n", message)
    fmt.Printf("%16s\n", message)
    fmt.Printf("%-16s\n", message)
    fmt.Printf("%s\n", message)
    fmt.Printf("%x\n", message)
    fmt.Printf("% x\n", message)
}

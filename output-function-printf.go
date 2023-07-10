// Go program to understand output function - printf

package main

import "fmt"

func main() {
    var humidity = 54.5
    var message = "Hello world!"

    fmt.Printf("%v\n", humidity)
    fmt.Printf("%#v\n", humidity)
    fmt.Printf("%v%%\n", humidity)
    fmt.Printf("%T\n", humidity)

    fmt.Printf("%v\n", message)
    fmt.Printf("%#v\n", message)
    fmt.Printf("%T\n", message)
}

// Go program to read from stdin - method 1

package main

import ( 
    "fmt"
    "bufio"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter name: ")
    name, _ := reader.ReadString('\n')
    fmt.Println(name)
}

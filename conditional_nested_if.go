// Go program to understand nested if

package main

import "fmt"

func main() {
    num := 100

    if(num > 50){
        fmt.Println(num, "is greater than 50")

        if(num > 75){ 
            fmt.Println(num, "is also greater than 75")
        }
    }
}

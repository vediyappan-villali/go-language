// Go program to convert integer to encoded string

package main

import "fmt"
import "strconv"

func main() {
    // num := 127
    // num := 32767
    // num := 2147483647
    num := 9223372036854775807
    encoded := strconv.FormatInt(int64(num), 16)
    fmt.Println("Encoded string:", encoded)
}


// Go program to understand switch with multiple case

package main

import "fmt"

func main() {
	day := 7

	switch day {
	case 2,3,4,5,6:
		fmt.Println("Weekday")
	case 1,7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}

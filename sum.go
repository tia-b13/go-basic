package main

import "fmt"

func main() {
	var message string = "TaiBbbTAAJDBTV"
	fmt.Printf("print the message string %s \n ", message)

	sum := 0
	for _, char := range message {
		sum = sum + int(char)
	}
	fmt.Printf("print the sum of ascii  %d ", sum)

	if (sum % 4) == 2 {
		fmt.Printf("true")
	} else {
		fmt.Printf("false")
	}
}

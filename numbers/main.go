package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Decimal \t Binary \t Hex \t UTF8 \n")
	fmt.Printf("%d \t %b \t %x \t %q \n", 42, 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \t %q \n", 42, 42, 42, 42)
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %#x \t %q \n", i, i, i, i)
	}
}

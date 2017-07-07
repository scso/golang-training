package main

import "fmt"

func main() {
	fmt.Println(42)
	fmt.Printf("%d - %b - %X \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X \t %#o \n", 42, 42, 42, 42)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d - %b - %#X \t %#o \n", i, i, i, i)
	}
}

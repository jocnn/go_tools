package main

import "fmt"

func main() {
	PrintHello()
	for i:=0; i < 5; i++ {
		PrintNumber(i)
	}
}

//revive:disable:exported

// PrintHello escribe el clásico Hola 
func PrintHello() {
	fmt.Println("Hello, Go")
}

//revive:disable:exported

// PrintNumber escribe un número usando la función fmt.Println
func PrintNumber(number int) {
	fmt.Println(number)
}
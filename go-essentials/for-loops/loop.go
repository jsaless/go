package main

import "fmt"

func main() {
	printRangeNumbers()
	printForLoopVariable()
}

func printRangeNumbers() {
	for number := range 10 {
		fmt.Println(number)
	}
}

func printForLoopVariable() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

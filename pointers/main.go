package main

import "fmt"

func main() {
	age := 30

	agePointer := &age
	fmt.Println("Age(Pointer):", *agePointer)

	getAdultYears(agePointer)
	fmt.Println("Age(Original):", age)
}

func getAdultYears(age *int) {
	*age -= 18
}

package main

import "fmt"

var (
	i int = 27
)

func main() {
	grades := []int{97, 85, 93, 90, 95}

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Amount of grades, %v\n", len(grades))
	fmt.Printf("Capacity of grades, %v\n", cap(grades))
	fmt.Printf("\n")
	// slicing operations

	a := grades[:]   // slice of all elements
	b := grades[3:]  // slice from the 4th element to the end
	c := grades[:3]  // slice of first two elements
	d := grades[1:4] // slice of 2nd, 3rd, and 4th elements

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	x := make([]int, 3)
	fmt.Println(x)
	fmt.Printf("Length: %v\n", len(x))
	fmt.Printf("Capacity: %v\n", cap(x))
}

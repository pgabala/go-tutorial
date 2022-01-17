package main

import "fmt"

func main() {
	statePopulations := map[string]int{
		"California": 123,
		"Texas":      456,
	}

	fmt.Println(statePopulations)

	statePopulations["Ohio"] = 789
	fmt.Println(statePopulations)

	delete(statePopulations, "Ohio")
	fmt.Println(statePopulations)

	// checking for presence
	_, ok := statePopulations["oho"]
	fmt.Println(statePopulations, ok)
}

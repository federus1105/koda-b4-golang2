package main

import "fmt"

func main() {
	scores := []int{50, 75, 66, 20, 32, 90}

	result := []int{}

	result = append(result, scores[:3]...)
	fmt.Println(result)
	result = append(result, 88)
	fmt.Println(result)
	result = append(result, scores[3:]...)

	for x := range len(result) {
		fmt.Println(result[x])
	}
}

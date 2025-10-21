package main

import "fmt"

func main() {
	scores := []int{50, 75, 66, 20, 32, 90}
	result := []int{}

	// TASK 1
	result = append(result, scores[:3]...)
	result = append(result, 88)
	result = append(result, scores[3:]...)

	for x := range len(result) {
		fmt.Println(result[x])
	}

	// TASK 2
	for i := 0; i < len(scores); i++ {
		if scores[i] == 66 {
			result = append(result, 88)
		}
		result = append(result, scores[i])
	}

	for _, val := range result {
		fmt.Println(val)
	}
}

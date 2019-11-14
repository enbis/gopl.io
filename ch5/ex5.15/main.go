package main

import "fmt"

func main() {
	fmt.Println("Max is ", variadicMax(1, 2, 3, 4, 5, 6))
	fmt.Println("Min is ", variadicMin(-1, -2, -3, -4, -5, -6))
	fmt.Println("Max is ", variadicMax())

}

func variadicMax(inputs ...int) int {
	max := 0
	for _, v := range inputs {
		if v > max {
			max = v
		}
	}
	return max
}

func variadicMin(inputs ...int) int {
	min := 0
	for _, v := range inputs {
		if v < min {
			min = v
		}
	}
	return min
}

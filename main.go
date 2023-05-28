package main

import "fmt"

func main() {
	var input [4]int
	var count int
	for i := 0; i < len(input); i++ {
		fmt.Scanln(&input[i])
		if input[i] >= 10 {
			count++
		}
	}
	fmt.Println(input, count)
}

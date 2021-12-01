package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./input.txt")
	input := strings.Fields(string(data))
	sum := 0
	WINDOW_SIZE := 3
	for i := range input {
		if (i + WINDOW_SIZE) >= len(input) {
			continue
		}
		currentWindowSum := 0
		nextWindowSum := 0
		for j := 0; j < WINDOW_SIZE; j++ {
			val, _ := strconv.Atoi(input[i+j])
			currentWindowSum = currentWindowSum + val
		}

		for r := 0; r < WINDOW_SIZE; r++ {
			val, _ := strconv.Atoi(input[i+1+r])
			nextWindowSum = nextWindowSum + val
		}

		if nextWindowSum > currentWindowSum {
			sum = sum + 1
		}
	}
	fmt.Print(sum)

}

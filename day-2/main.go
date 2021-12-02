package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Location struct {
	depth      int64
	horizontal int64
	aim        int64
}

func main() {
	fileContent, _ := os.ReadFile("./input.txt")
	input := strings.Split(string(fileContent), "\n")
	location := Location{depth: 0, horizontal: 0, aim: 0}
	for _, v := range input {
		current := strings.Split(v, " ")
		action := current[0]
		value, _ := strconv.Atoi(current[1])

		switch action {
		case "forward":
			location.horizontal += int64(value)
			location.depth += location.aim * int64(value)
		case "up":
			// location.depth -= int64(value)
			location.aim -= int64(value)
		case "down":
			// location.depth += int64(value)
			location.aim += int64(value)
		}
	}
	result := location.depth * location.horizontal
	fmt.Print(result)
}

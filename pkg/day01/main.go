package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadPuzzleInput() []int {
	file, err := os.Open("./puzzle-input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var numbers []int

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}

	return numbers
}

func main() {
	depths := loadPuzzleInput()
	numIncreases := 0
	for i := 1; i < len(depths); i++ {
		prev := depths[i-1]
		current := depths[i]
		if current > prev {
			numIncreases++
		}
	}
	fmt.Println(numIncreases)
}

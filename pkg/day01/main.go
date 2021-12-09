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

func Part1(depths []int) int {
	numIncreases := 0
	for i := 1; i < len(depths); i++ {
		prev := depths[i-1]
		current := depths[i]
		if current > prev {
			numIncreases++
		}
	}
	return numIncreases
}

func Part2(depths []int) int {
	numIncreases := 0
	prevSum := 0
	for i := 1; i < len(depths)-2; i++ {
		threeSum := 0
		for j := i; j < i+3; j++ {
			threeSum += depths[j]
		}
		if threeSum > prevSum {
			numIncreases++
		}
		prevSum = threeSum
	}
	return numIncreases
}

func main() {
	fmt.Println(Part1(loadPuzzleInput()))
	fmt.Println(Part2(loadPuzzleInput()))
}

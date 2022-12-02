package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputPath = "./day01/input.txt"

func Run() error {
	if err := Part1(); err != nil {
		return err
	}
	if err := Part2(); err != nil {
		return err
	}
	return nil
}

func Part1() error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	maxSum := 0
	currentSum := 0

	for scanner.Scan() {
		if line := strings.TrimSpace(scanner.Text()); len(line) > 0 {
			lineInt, err := strconv.Atoi(line)
			if err != nil {
				return err
			}
			currentSum += lineInt
			if currentSum >= maxSum {
				maxSum = currentSum
			}
		} else {
			currentSum = 0
		}
	}
	fmt.Printf("part 1: %d\n", maxSum)
	return nil
}

func Part2() error {
	// open file lazily
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	// create a buffer scan of the file split by lines
	scanner := bufio.NewScanner(file)
	// store top 3 calorie counts descending
	top3 := make([]int, 3)
	// store current elf's calories
	currentSum := 0
	// enumerate lines of file
	for scanner.Scan() {
		if line := strings.TrimSpace(scanner.Text()); len(line) > 0 {
			lineInt, err := strconv.Atoi(line)
			if err != nil {
				return err
			}
			currentSum += lineInt
		} else {
			evalTop3(top3, currentSum)
			currentSum = 0
		}
	}
	evalTop3(top3, currentSum)
	out := 0
	for _, value := range top3 {
		out += value
	}
	fmt.Printf("part 2: %d\n", out)
	return nil
}

func evalTop3(top3 []int, newValue int) {
	for ix, value := range top3 {
		if value < newValue {
			for backwards := 2; backwards > ix; backwards-- {
				top3[backwards] = top3[backwards-1]
			}
			top3[ix] = newValue
			return
		}
	}
}

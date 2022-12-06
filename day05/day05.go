package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zenyui/aoc-22/util"
)

const (
	inputPath = "day05/input.txt"
)

func Run() error {
	if err := part1(); err != nil {
		return err
	}
	if err := part2(); err != nil {
		return err
	}
	return nil
}

func part1() error {
	lines, err := util.ReadFile(inputPath)
	if err != nil {
		return err
	}

	firstSection := make([]string, 0, 5)
	// find the end of the first section first
	for line := range lines {
		if len(line) == 0 {
			break
		}
		line = strings.TrimRight(line, " ")
		firstSection = append(firstSection, line)
	}
	// create store for creates
	labels := firstSection[len(firstSection)-1]
	labels = strings.ReplaceAll(labels, " ", "")
	firstSection = firstSection[:len(firstSection)-1]
	stacks := make(map[string][]string, len(labels))
	for ix := 0; ix < len(labels); ix++ {
		label := string(labels[ix])
		stacks[label] = make([]string, 0, len(firstSection)-1)
		col := 4*ix + 1
		for row := len(firstSection) - 1; row >= 0; row-- {
			line := firstSection[row]
			if len(line) < col || string(line[col]) == " " {
				break
			}
			stacks[label] = append(stacks[label], string(line[col]))
		}
	}
	// process each instruction
	for line := range lines {
		instructions := strings.Split(line, " ")
		count, err := strconv.Atoi(instructions[1])
		if err != nil {
			return err
		}
		fromKey := string(instructions[3])
		toKey := string(instructions[5])
		fromStack := stacks[fromKey]
		toStack := stacks[toKey]

		for ix := 0; ix < count; ix++ {
			toStack = append(toStack, fromStack[len(fromStack)-1])
			fromStack = fromStack[:len(fromStack)-1]
		}

		stacks[toKey] = toStack
		stacks[fromKey] = fromStack
	}
	out := ""
	for ix := 0; ix < len(labels); ix++ {
		stack := stacks[string(labels[ix])]
		out = out + stack[len(stack)-1]
	}
	fmt.Printf("part 1: %s\n", out)
	return nil
}

func part2() error {
	lines, err := util.ReadFile(inputPath)
	if err != nil {
		return err
	}

	firstSection := make([]string, 0, 5)
	// find the end of the first section first
	for line := range lines {
		if len(line) == 0 {
			break
		}
		line = strings.TrimRight(line, " ")
		firstSection = append(firstSection, line)
	}
	// create store for creates
	labels := firstSection[len(firstSection)-1]
	labels = strings.ReplaceAll(labels, " ", "")
	firstSection = firstSection[:len(firstSection)-1]
	stacks := make(map[string][]string, len(labels))
	for ix := 0; ix < len(labels); ix++ {
		label := string(labels[ix])
		stacks[label] = make([]string, 0, len(firstSection)-1)
		col := 4*ix + 1
		for row := len(firstSection) - 1; row >= 0; row-- {
			line := firstSection[row]
			if len(line) < col || string(line[col]) == " " {
				break
			}
			stacks[label] = append(stacks[label], string(line[col]))
		}
	}
	// process each instruction
	for line := range lines {
		instructions := strings.Split(line, " ")
		count, err := strconv.Atoi(instructions[1])
		if err != nil {
			return err
		}
		fromKey := string(instructions[3])
		toKey := string(instructions[5])
		fromStack := stacks[fromKey]
		toStack := stacks[toKey]

		toStack = append(toStack, fromStack[len(fromStack)-count:]...)
		fromStack = fromStack[:len(fromStack)-count]

		stacks[toKey] = toStack
		stacks[fromKey] = fromStack
	}
	out := ""
	for ix := 0; ix < len(labels); ix++ {
		stack := stacks[string(labels[ix])]
		out = out + stack[len(stack)-1]
	}
	fmt.Printf("part 2: %s\n", out)
	return nil
}

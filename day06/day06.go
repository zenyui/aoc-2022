package day06

import (
	"fmt"

	"github.com/zenyui/aoc-22/util"
)

const (
	inputPath = "./day06/input.txt"
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
	chars, err := util.ReadCharacters(inputPath)
	if err != nil {
		return err
	}
	ix := 0
	letters := make([]string, 4)
	for c := range chars {
		// skip newlines
		if c == "\n" {
			continue
		}
		letters[ix%len(letters)] = c
		ix += 1
		if ix >= 4 && allDifferent(letters...) {
			fmt.Printf("part 1: %d, %v\n", ix, letters)
			return nil
		}
	}
	fmt.Printf("part 1: no code found\n")
	return nil
}

func part2() error {
	chars, err := util.ReadCharacters(inputPath)
	if err != nil {
		return err
	}
	markerSize := 14
	ix := 0
	letters := make([]string, markerSize)
	for c := range chars {
		// skip newlines
		if c == "\n" {
			continue
		}
		letters[ix%len(letters)] = c
		ix += 1
		if ix >= markerSize && allDifferent(letters...) {
			fmt.Printf("part 2: %d, %v\n", ix, letters)
			return nil
		}
	}
	fmt.Printf("part 2: no code found\n")
	return nil
}

func allDifferent(chars ...string) bool {
	records := make(map[string]*struct{}, len(chars))
	var exists bool
	for _, c := range chars {
		if records[c], exists = records[c]; exists {
			return false
		}
	}
	return true
}

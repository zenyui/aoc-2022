package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zenyui/aoc-22/util"
)

const (
	inputPath = "./day04/input.txt"
)

// Run part 4
func Run() error {
	lines, err := util.ReadFile(inputPath)
	if err != nil {
		return err
	}
	numSubsets := 0
	numOverlaps := 0
	for line := range lines {
		intRanges, err := lineToRanges(line)
		if err != nil {
			return err
		}
		if len(intRanges) != 2 {
			return fmt.Errorf("malformed line %s", line)
		}
		if isOverlap(intRanges[0], intRanges[1]) {
			numOverlaps += 1
			if isSubset(intRanges[0], intRanges[1]) {
				numSubsets += 1
			}
		}
	}
	fmt.Printf("part 1: %d\n", numSubsets)
	fmt.Printf("part 2: %d\n", numOverlaps)
	return nil
}

// inclusiveRange represents an inclusive range of ints
type inclusiveRange struct {
	Start int
	End   int
}

// lineToRanges parses an input line and returns the int ranges found
func lineToRanges(line string) ([]inclusiveRange, error) {
	out := make([]inclusiveRange, 0, 2)
	for _, x := range strings.Split(line, ",") {
		elements := strings.Split(x, "-")
		if len(elements) != 2 {
			return nil, fmt.Errorf("malformed line %s", line)
		}
		start, err := strconv.Atoi(elements[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(elements[1])
		if err != nil {
			return nil, err
		}
		out = append(out, inclusiveRange{Start: start, End: end})
	}
	return out, nil
}

// isSubset returns true if one range is a subset of the other
func isSubset(left, right inclusiveRange) bool {
	return (left.Start <= right.Start && left.End >= right.End) ||
		(right.Start <= left.Start) && (right.End >= left.End)
}

// isOverlap returns true if either intRange overlap at all
func isOverlap(left, right inclusiveRange) bool {
	if left.Start <= right.Start {
		return left.End >= right.Start
	} else {
		return isOverlap(right, left)
	}
}

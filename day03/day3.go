package day03

import (
	"fmt"
	"strings"

	"github.com/zenyui/aoc-22/util"
)

const (
	inputPath = "./day03/input.txt"
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
	letters := getLetters()
	totalScore := 0
	for line := range lines {
		items := strings.Split(line, "")
		// compute length of line
		compartmentSize := len(items) / 2
		// turn left side into a set
		leftMap := map[string]*struct{}{}
		for i := 0; i < compartmentSize; i++ {
			leftMap[items[i]] = nil
		}
		// traverse right side to find items also in the left
		for i := compartmentSize; i < len(items); i++ {
			item := items[i]
			if _, found := leftMap[item]; found {
				points := letters[item]
				totalScore += points
				break
			}
		}
	}
	fmt.Printf("part 1: %d\n", totalScore)
	return nil
}

func part2() error {
	lines, err := util.ReadFile(inputPath)
	if err != nil {
		return err
	}
	letters := getLetters()
	totalScore := 0
	groupCount := 0
	var group []string
	for line := range lines {
		group = append(group, line)
		if len(group) == 3 {
			// fmt.Println("****************")
			groupCount += 1
			groupBadge := findGroupBadge(group)
			groupPoints := letters[groupBadge]
			// fmt.Printf("group %d: badge=%s, points=%d\n", groupCount, groupBadge, groupPoints)
			totalScore += groupPoints
			group = nil

			// if groupCount >= 2 {
			// 	break
			// }
		}
	}
	fmt.Printf("part 2: %d\n", totalScore)
	return nil
}

func findGroupBadge(group []string) string {
	pack0 := strings.Split(group[0], "")
	pack1 := strings.Split(group[1], "")
	pack2 := strings.Split(group[2], "")
	common0 := map[string]bool{}
	for _, letter := range pack0 {
		common0[letter] = true
	}
	common1 := map[string]bool{}
	for _, letter := range pack1 {
		common1[letter] = true
	}
	for _, letter := range pack2 {
		_, found0 := common0[letter]
		_, found1 := common1[letter]
		if found0 && found1 {
			return letter
		}
	}
	return ""
}

// get all letters as a map to points
func getLetters() map[string]int {
	out := make(map[string]int, 26*2)
	points := 1
	for r := 'a'; r <= 'z'; r++ {
		letter := string(r)
		out[letter] = points
		out[strings.ToUpper(letter)] = points + 26
		points += 1
	}
	return out
}

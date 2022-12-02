package main

import (
	"log"

	"github.com/zenyui/aoc-22/day01"
)

func main() {
	if err := day01.Run(); err != nil {
		log.Fatal(err)
	}
}

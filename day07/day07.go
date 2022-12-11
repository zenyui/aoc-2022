package day07

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/zenyui/aoc-22/util"
)

const (
	inputPath = "./day07/input.txt"
)

func Run() error {
	lines, err := util.ReadFile(inputPath)
	if err != nil {
		return err
	}
	root := &Folder{Parent: nil, Files: nil, Folders: make(map[string]*Folder)}
	currentLocation := root
	currentCmd := ""

	for line := range lines {
		// if it's a command
		if line[0] == '$' {
			lineSplit := strings.SplitN(line, " ", 3)
			if len(lineSplit) < 2 {
				return fmt.Errorf("malformed line %s", line)
			}
			currentCmd = lineSplit[1]
			ending := ""
			if len(lineSplit) > 2 {
				ending = lineSplit[2]
			}
			// fmt.Printf("%s | %s\n", currentCmd, ending)
			switch currentCmd {
			case "cd":
				switch ending {
				case "/":
					currentLocation = root
				case ".":
					// go nowhere
				case "..":
					if currentLocation.Parent != nil {
						currentLocation = currentLocation.Parent
					}
				default:
					folder, exists := currentLocation.Folders[ending]
					if !exists {
						return fmt.Errorf("not found, %s", folder.Name)
					}
					currentLocation = folder
				}
				// fmt.Printf("loc --> %s\n", currentLocation.AbsPath())

			case "ls":
				// we are listing now
				currentLocation.Size = 0
				currentLocation.Folders = make(map[string]*Folder)
			}
		} else {
			// process lines after cmd
			switch currentCmd {
			case "ls":
				if line[:3] == "dir" {
					dirName := strings.SplitN(line, " ", 2)[1]
					folder := &Folder{Parent: currentLocation, Name: dirName, Folders: make(map[string]*Folder)}
					currentLocation.Folders[dirName] = folder
				} else {
					lineSplit := strings.SplitN(line, " ", 2)
					fileSize, err := strconv.Atoi(lineSplit[0])
					if err != nil {
						return err
					}
					currentLocation.Size += fileSize
					currentLocation.Files = append(currentLocation.Files, lineSplit[1])
				}
			}
		}
	}

	part1 := 0
	totalSize := dfsFolderSizes(root, 100000, &part1)
	fmt.Printf("part 1: %d\n", part1)
	freeSpace := 70000000 - totalSize
	neededSpace := 30000000 - freeSpace
	part2 := 0
	findDelete(root, neededSpace, &part2)
	fmt.Printf("part 2: %d\n", part2)

	return nil
}

type Folder struct {
	Parent  *Folder
	Name    string
	Files   []string
	Folders map[string]*Folder // name -> folder
	Size    int
}

func (f *Folder) AbsPath() string {
	if f == nil {
		return "/"
	}
	return path.Join(f.Parent.AbsPath(), f.Name)
}

func dfsFolderSizes(folder *Folder, maxSize int, total *int) int {
	out := folder.Size
	for _, child := range folder.Folders {
		out += dfsFolderSizes(child, maxSize, total)
	}
	if out <= maxSize {
		var newTotal int = *total + out
		*total = newTotal
	}
	return out
}

func findDelete(folder *Folder, spaceNeeded int, target *int) int {
	out := folder.Size
	for _, child := range folder.Folders {
		out += findDelete(child, spaceNeeded, target)
	}
	if out >= spaceNeeded {
		if *target == 0 || *target > out {
			*target = out
		}
	}
	return out
}

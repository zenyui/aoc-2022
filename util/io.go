package util

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"unicode"
)

// ReadFile reads file per line into a channel
func ReadFile(path string) (chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// create output channel
	out := make(chan string, 10)
	// process file async
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			out <- scanner.Text()
		}
		file.Close()
		close(out)
	}()
	// return the channel
	return out, nil
}

// ReadCharacters reads a single character at a time
func ReadCharacters(path string) (chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// create output channel
	out := make(chan string, 10)
	// process file async
	go func() {
		reader := bufio.NewReader(file)
		for {
			if c, _, err := reader.ReadRune(); err != nil {
				if err == io.EOF {
					break
				} else {
					log.Fatal(err)
				}
			} else {
				if c == unicode.ReplacementChar {
					log.Fatal(errors.New("invalid byte read"))
				}
				out <- string(c)
			}
		}
		file.Close()
		close(out)
	}()
	// return the channel
	return out, nil
}

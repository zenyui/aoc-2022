package util

import (
	"bufio"
	"os"
	"strings"
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
			out <- strings.TrimSpace(scanner.Text())
		}
		file.Close()
		close(out)
	}()
	// return the channel
	return out, nil
}

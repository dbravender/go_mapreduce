// Package fileiter provides utilities for iterating over file contents.
package fileiter

import (
	"bufio"
	"iter"
	"os"
)

// Lines returns an iterator over lines in a file.
// The iterator yields each line without the trailing newline character.
// If the file cannot be opened, the iterator yields no values.
func Lines(filename string) iter.Seq[string] {
	return func(yield func(string) bool) {
		file, err := os.Open(filename)
		if err != nil {
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}
}

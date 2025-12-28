// Command wordcount demonstrates using the mapreduce package to count words across files.
package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"

	"github.com/dbravender/go_mapreduce/fileiter"
	"github.com/dbravender/go_mapreduce/mapreduce"
)

// WordCounts maps words to their occurrence counts.
type WordCounts map[string]int

// findFiles returns a channel that yields all file paths under the given directory.
func findFiles(root string) <-chan string {
	output := make(chan string)
	go func() {
		defer close(output)
		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if !d.IsDir() {
				output <- path
			}
			return nil
		})
	}()
	return output
}

var wordsRE = regexp.MustCompile(`[A-Za-z0-9_]+`)

// countWords processes a single file and sends word counts to the output channel.
func countWords(filename string, output chan<- WordCounts) {
	counts := make(WordCounts)
	for line := range fileiter.Lines(filename) {
		for _, word := range wordsRE.FindAllString(line, -1) {
			counts[word]++
		}
	}
	output <- counts
}

// reduce aggregates word counts from all mappers into a final result.
func reduce(input <-chan WordCounts, output chan<- WordCounts) {
	totals := make(WordCounts)
	for counts := range input {
		for word, count := range counts {
			totals[word] += count
		}
	}
	output <- totals
}

func main() {
	result := mapreduce.MapReduce(countWords, reduce, findFiles("."), 20)
	fmt.Println(result)
}

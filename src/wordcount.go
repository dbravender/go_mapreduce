package main

import (
    "os"
    "fmt"
    "regexp"
    "file_iter"
    "mapreduce"
)

func find_files(dirname string) chan interface{} {
    output := make(chan interface{})
    go func() {
        _find_files(dirname, output)
        close(output)
    }()
    return output
}

func _find_files(dirname string, output chan interface{}) {
    dir, _ := os.Open(dirname, os.O_RDONLY, 0)
    dirnames, _ := dir.Readdirnames(-1)
    for i := 0; i < len(dirnames); i++ {
        fullpath := dirname + "/" + dirnames[i]
        file, _ := os.Stat(fullpath)
        if file.IsDirectory() {
            _find_files(fullpath, output)
        } else {
            output <- fullpath
        }
    }
}

func wordcount(filename interface{}, output chan interface{}) {
    results := map[string]int{}
    wordsRE := regexp.MustCompile(`[A-Za-z0-9_]*`)
    for line := range file_iter.EachLine(filename.(string)) {
        for _, match := range wordsRE.FindAllString(line, -1) {
            results[match]++
        }
    }
    output <- results
}

func reducer(input chan interface{}, output chan interface{}) {
    results := map[string]int{}
    for new_matches := range input {
        for key, value := range new_matches.(map[string]int) {
            previous_count, exists := results[key]
            if !exists {
                results[key] = value
            } else {
                results[key] = previous_count + value
            }
        }
    }
    output <- results
}

func main() {
    fmt.Print(mapreduce.MapReduce(wordcount, reducer, find_files("."), 10))
}

package main

import (
    "mapreduce"
    "fmt"
    "time"
)

func main() {
    input_channel := make(chan interface{})
    go func() {
        for i := 0; i < 100; i++ {
            // This helps reveal time-dependant bugs
            time.Sleep(1000000)
            input_channel <- i
        }
        close(input_channel)
    }()
    result := mapreduce.MapReduce(
        func(x interface{}, output chan interface{}) {
            fmt.Println("Mapping ", x)
            time.Sleep(100000000)
            output <- x.(int) * x.(int)
        },
        func(input chan interface{}, output chan interface{}) {
            total := 0
            for item := range input {
                fmt.Println("Reducing: ", item.(int))
                total += item.(int)
            }
            output <- total
        },
        input_channel, 10)

    if result.(int) != 328350 {
        fmt.Println(result.(int), "Unexpected MapReduce result")
    } else {
        fmt.Println("OK")
    }
}

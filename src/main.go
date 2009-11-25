package main

import (
    "mapreduce";
    "fmt";
    "time";
)

func main() {
    input_channel := make(chan interface{});
    go func() {
        for _, value := range [...]int {1, 2, 3, 4} {
            time.Sleep(10000);
            input_channel <- value;
        }
        close(input_channel);
    }();
    result := mapreduce.MapReduce(
        func(x interface{}, output chan interface{}) {
            output <- x.(int) * x.(int);
        },
        func(input chan interface{}, output chan interface{}) {
            total := 0;
            for item := range input {
                total += item.(int);
            }
            output <- total;
        }, 
        input_channel
    );

    if result.(int) != 30 {
        fmt.Println("Unexpected MapReduce result");
    } else {
        fmt.Println("OK");
    }
}

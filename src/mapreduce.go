package mapreduce

func MapReduce(mapper func(interface{}, chan interface{}),
               reducer func(chan interface{}, chan interface{}),
               input chan interface{}) interface{} 
{
    reduce_input     := make(chan interface{});
    reduce_output    := make(chan interface{});
    worker_output    := make(chan interface{});
    worker_processes := 0;
    go reducer(reduce_input, reduce_output);
    go func() {
        for item := range input {
            go mapper(item, worker_output);
            worker_processes += 1;
        }
    }();
    for item := range worker_output {
        worker_processes -= 1;
        reduce_input <- item;
        if worker_processes <= 0 {
            close(reduce_input);
            break;
        }
    }
    return <- reduce_output;
}

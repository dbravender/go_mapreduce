// Package mapreduce provides a simple, generic MapReduce implementation using Go channels.
package mapreduce

// MapReduce executes a map-reduce operation with configurable parallelism.
//
// Type parameters:
//   - In: the type of input items
//   - Mid: the intermediate type produced by mappers and consumed by the reducer
//   - Out: the final output type produced by the reducer
//
// Parameters:
//   - mapper: processes each input item and sends results to the provided channel
//   - reducer: aggregates results from all mappers and sends final result to output channel
//   - input: channel providing input items to process
//   - poolSize: maximum number of concurrent mapper goroutines
func MapReduce[In, Mid, Out any](
	mapper func(In, chan<- Mid),
	reducer func(<-chan Mid, chan<- Out),
	input <-chan In,
	poolSize int,
) Out {
	reduceInput := make(chan Mid)
	reduceOutput := make(chan Out)
	workerOutput := make(chan chan Mid, poolSize)

	go reducer(reduceInput, reduceOutput)

	// Collector: receives mapper outputs and forwards to reducer
	go func() {
		for workerChan := range workerOutput {
			reduceInput <- <-workerChan
		}
		close(reduceInput)
	}()

	// Dispatcher: spawns mapper goroutines for each input
	go func() {
		for item := range input {
			ch := make(chan Mid)
			go mapper(item, ch)
			workerOutput <- ch
		}
		close(workerOutput)
	}()

	return <-reduceOutput
}

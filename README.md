# go_mapreduce

A simple MapReduce implementation in Go using channels and generics.

This is a learning project for understanding Go channels and MapReduce basics. Not intended for production use.

This code was written within a year of Go's release so the initial commits have semicolons but I have kept it up-to-date in terms of compatibility and style.

## Installation

```bash
go get github.com/dbravender/go_mapreduce
```

## Usage

```go
import "github.com/dbravender/go_mapreduce/mapreduce"

result := mapreduce.MapReduce(mapper, reducer, inputChan, poolSize)
```

- `mapper`: `func(In, chan<- Mid)` - processes each input, sends result to channel
- `reducer`: `func(<-chan Mid, chan<- Out)` - aggregates mapper outputs
- `inputChan`: `<-chan In` - input items to process
- `poolSize`: `int` - max concurrent mappers

## Example

See `cmd/wordcount/main.go` for a word counting example.

```bash
go run ./cmd/wordcount
```

## License

MIT

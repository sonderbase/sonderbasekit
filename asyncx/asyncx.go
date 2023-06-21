package asyncx

import (
	"sync"
)

// Result is the main struct that holds future result
type Result[T any] struct {
	Res T
	Err error
}

// Future is a future wrapper to Result
type Future[T any] struct {
	result Result[T]
	ch     chan Result[T]
	// make sure we only await channel once
	awaitOnce sync.Once
}

// Async wraps a function to async fashion
func Async[T any](
	fn func() (T, error),
) *Future[T] {
	f := &Future[T]{
		ch: make(chan Result[T]),
	}

	go func() {
		output, err := fn()
		f.ch <- Result[T]{Res: output, Err: err}
	}()

	return f
}

// Await returns result
func (f *Future[T]) Await() (T, error) {
	f.awaitOnce.Do(func() {
		f.result = <-f.ch
	})
	return f.result.Res, f.result.Err
}

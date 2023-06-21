package asyncx_test

import (
	"testing"

	"github.com/sonderbase/sonderbasekit/asyncx"
)

func asyncJob() (any, error) {
	return nil, nil
}

func BenchmarkAsync(b *testing.B) {
	benchCases := []struct {
		name  string
		count int
	}{
		{
			name:  "1 async",
			count: 1,
		},
		{
			name:  "1,000 async",
			count: 1_000,
		},
		{
			name:  "1,000,000 async",
			count: 1_000_000,
		},
	}

	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			// construct asyncs
			jobs := make([]*asyncx.Future[any], bc.count)
			for i := range jobs {
				jobs[i] = asyncx.Async(asyncJob)
			}
		})
	}
}

func BenchmarkAwait(b *testing.B) {
	benchCases := []struct {
		name  string
		count int
	}{
		{
			name:  "1 await",
			count: 1,
		},
		{
			name:  "1,000 await",
			count: 1_000,
		},
		{
			name:  "1,000,000 await",
			count: 1_000_000,
		},
	}

	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			// construct asyncs
			jobs := make([]*asyncx.Future[any], bc.count)
			for i := range jobs {
				jobs[i] = asyncx.Async(asyncJob)
			}

			b.ResetTimer()
			for _, job := range jobs {
				job.Await()
			}
		})
	}
}

func BenchmarkAsyncAwait(b *testing.B) {
	benchCases := []struct {
		name  string
		count int
	}{
		{
			name:  "1 await",
			count: 1,
		},
		{
			name:  "1,000 await",
			count: 1_000,
		},
		{
			name:  "1,000,000 await",
			count: 1_000_000,
		},
	}

	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			// construct asyncs
			jobs := make([]*asyncx.Future[any], bc.count)
			for i := range jobs {
				jobs[i] = asyncx.Async(asyncJob)
			}

			for _, job := range jobs {
				job.Await()
			}
		})
	}
}

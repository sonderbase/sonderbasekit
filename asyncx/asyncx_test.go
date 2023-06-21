package asyncx_test

import (
	"errors"
	"testing"

	"github.com/samber/lo"
	"github.com/sonderbase/sonderbasekit/asyncx"
	"github.com/stretchr/testify/require"
)

func TestAsync(t *testing.T) {
	tests := []struct {
		name string
		fn   func() (interface{}, error)
	}{
		{
			name: "resolve",
			fn: func() (interface{}, error) {
				return true, nil
			},
		},
		{
			name: "reject",
			fn: func() (interface{}, error) {
				return nil, errors.New("invalid")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := asyncx.Async(tc.fn)
			require.NotNil(t, f)
		})
	}
}

func TestAwait(t *testing.T) {
	tests := []struct {
		name   string
		fn     func() (interface{}, error)
		expect asyncx.Result[any]
	}{
		{
			name: "resolve with true",
			fn: func() (interface{}, error) {
				return true, nil
			},
			expect: asyncx.Result[any]{Res: true},
		},
		{
			name: "resolve with nil",
			fn: func() (interface{}, error) {
				return nil, nil
			},
			expect: asyncx.Result[any]{Res: nil},
		},
		{
			name: "resolve with pointer",
			fn: func() (interface{}, error) {
				return lo.ToPtr("something"), nil
			},
			expect: asyncx.Result[any]{Res: lo.ToPtr("something")},
		},
		{
			name: "reject with error",
			fn: func() (interface{}, error) {
				return nil, errors.New("err")
			},
			expect: asyncx.Result[any]{Err: errors.New("err")},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := asyncx.Async(tc.fn)
			result, err := f.Await()
			require.Equal(t, tc.expect.Res, result)
			require.Equal(t, tc.expect.Err, err)
		})
	}
}

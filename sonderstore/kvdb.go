package sonderstore

import (
	"context"
	"io"
)

type DocGenerator struct {
	Path string
	RC   io.ReadCloser
	Err  error
}

type KVDB interface {
	GetOne(ctx context.Context, path string) (io.ReadCloser, error)
	Set(ctx context.Context, path string, r io.Reader) error
	Delete(ctx context.Context, path string) error

	// ForEach(ctx context.Context, prefix string, afterID *string, beforeID *string) (<-chan DocGenerator, error)
}

package maxbytesreader

import (
	"errors"
	"io"
)

// Throws an error if the reader is bigger than limit.

var SIZEEXCEEDED = errors.New("http: response too big")

type MaxBytesReader struct {
	io.ReadCloser       // reader object
	N             int64 // max bytes remaining.
}

func NewMaxBytesReader(r io.ReadCloser, limit int64) *MaxBytesReader {
	return &MaxBytesReader{r, limit}
}

func (b *MaxBytesReader) Read(p []byte) (n int, err error) {
	if b.N <= 0 {
		return 0, SIZEEXCEEDED
	}

	if int64(len(p)) > b.N {
		p = p[0:b.N]
	}

	n, err = b.ReadCloser.Read(p)
	b.N -= int64(n)
	return
}

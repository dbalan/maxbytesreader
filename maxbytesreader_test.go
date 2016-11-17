package maxbytesreader

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

type DummyReadCloser struct {
	io.Reader
}

func (d *DummyReadCloser) Close() error {
	return nil
}

func NewDummyReader(val string) *DummyReadCloser {
	rdr := strings.NewReader(val)
	return &DummyReadCloser{rdr}
}

func TestReader(t *testing.T) {
	val := "hello world"
	rdr := NewMaxBytesReader(NewDummyReader(val), 1000)
	_, err := ioutil.ReadAll(rdr)

	if err != nil {
		t.Error("an error occured", err)
	}
}

func TestFailingReader(t *testing.T) {
	val := "hello world"
	l := len(val) - 2
	rdr := NewMaxBytesReader(NewDummyReader(val), int64(l))

	res, err := ioutil.ReadAll(rdr)

	if err == nil {
		t.Error("Should've raised an error: response toobig")
	}

	if string(res) != val[:l] {
		t.Error("got unexpected response: ", string(res))
	}
}

package bertlv

import (
	"io"
)

type countReader struct {
	io.Reader
	Length *int64
}

func (c *countReader) Read(p []byte) (n int, err error) {
	n, err = c.Reader.Read(p)
	*c.Length += int64(n)
	return
}

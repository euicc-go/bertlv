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

func (c *countReader) Offset() (offset int64) {
	for r, ok := c, true; ok; r, ok = r.Reader.(*countReader) {
		offset += *r.Length
	}
	return
}

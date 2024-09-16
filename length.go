package bertlv

import (
	"errors"
	"fmt"
	"io"
)

func marshalLength(n uint16) []byte {
	switch {
	case n < 128:
		return []byte{byte(n)}
	case n < 256:
		return []byte{0x81, byte(n)}
	}
	return []byte{0x82, byte(n >> 8), byte(n)}
}

func readLength(r io.Reader) (value uint16, err error) {
	var length [3]byte
	if _, err = r.Read(length[0:1]); err != nil {
		return 0, fmt.Errorf("failed to read length: %w", err)
	}
	if length[0] < 128 {
		return uint16(length[0]), nil
	}
	switch length[0] {
	case 0x81:
		if _, err = r.Read(length[1:2]); err == nil {
			value = uint16(length[1])
		}
	case 0x82:
		if _, err = r.Read(length[1:3]); err == nil {
			value = uint16(length[1])<<8 | uint16(length[2])
		}
	default:
		err = errors.New("if length is greater than 127, first byte must indicate encoding of length")
	}
	if err != nil {
		err = fmt.Errorf("failed to read length: %w", err)
	}
	return
}

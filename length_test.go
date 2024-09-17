package bertlv

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLength(t *testing.T) {
	fixtures := map[uint16][]byte{
		0x00:   {0x00},
		0x01:   {0x01},
		0x7f:   {0x7f},
		0x80:   {0x81, 0x80},
		0xff:   {0x81, 0xff},
		0x100:  {0x82, 0x01, 0x00},
		0xffff: {0x82, 0xff, 0xff},
	}
	for length, expected := range fixtures {
		assert.Equal(t, expected, marshalLength(length))
		value, err := readLength(bytes.NewReader(expected))
		assert.Equal(t, length, value)
		assert.NoError(t, err)
	}
}

func TestLength_Error(t *testing.T) {
	type Fixture struct {
		Length []byte
		Error  string
	}
	fixtures := []*Fixture{
		{[]byte{}, "failed to read length: EOF"},
		{[]byte{0x80}, "failed to read length: if length is greater than 127, first byte must indicate encoding of length"},
		{[]byte{0x81}, "failed to read length: EOF"},
		{[]byte{0x82}, "failed to read length: EOF"},
	}
	var err error
	for _, fixture := range fixtures {
		_, err = readLength(bytes.NewReader(fixture.Length))
		assert.EqualError(t, err, fixture.Error)
	}
}

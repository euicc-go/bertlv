package primitive

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitString(t *testing.T) {
	var bits BitString
	assert.NoError(t, (&bits).UnmarshalBinary([]byte{0x06, 0x6E, 0x5D, 0xC0}))
	assert.Equal(t, "011011100101110111", bits.String())
	data, err := bits.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, []byte{0x06, 0x6E, 0x5D, 0xC0}, data)

	flags := make([]bool, 4)
	bits.Read(&flags[0], &flags[1], &flags[2], nil, &flags[3])
	assert.Equal(t, bits[0], flags[0])
	assert.Equal(t, bits[1], flags[1])
	assert.Equal(t, bits[2], flags[2])
	assert.Equal(t, bits[4], flags[3])
}

func TestBitStringError(t *testing.T) {
	var bits BitString
	assert.Error(t, (&bits).UnmarshalBinary([]byte{0x08, 0x6E, 0x5D, 0xC0}))
}

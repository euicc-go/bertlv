package primitive

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBigInt(t *testing.T) {
	var bigInt BigInt
	assert.NoError(t, bigInt.UnmarshalBinary([]byte{0x7f}))
	assert.Equal(t, int64(127), bigInt.Int64())
	data, err := bigInt.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, []byte{0x7f}, data)
}

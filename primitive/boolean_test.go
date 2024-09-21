package primitive

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolean(t *testing.T) {
	type Fixture struct {
		Expected Boolean
		Variants [][]byte
	}
	fixtures := []*Fixture{
		{false, [][]byte{{0x00}, nil, {0x01}}},
		{true, [][]byte{{0xff}}},
	}
	var err error
	var output []byte
	for _, fixture := range fixtures {
		for _, variant := range fixture.Variants {
			var parsed Boolean
			_ = (&parsed).UnmarshalBinary(variant)
			assert.Equal(t, fixture.Expected, parsed)
		}
		output, err = fixture.Expected.MarshalBinary()
		assert.NoError(t, err)
		assert.Equal(t, fixture.Variants[0], output)
	}
}

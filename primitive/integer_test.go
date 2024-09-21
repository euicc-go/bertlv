package primitive

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestInteger(t *testing.T) {
	testInt(t, 8, map[Int64][][]byte{
		0:             {{0x00}},
		127:           {{0x7f}, {0x00, 0x7f}},
		128:           {{0x00, 0x80}},
		256:           {{0x01, 0x00}},
		-1:            {{0xff}, {0xff, 0xff}, {0xff, 0xff, 0xff, 0xff}},
		-128:          {{0x80}, {0xff, 0x80}},
		-129:          {{0xff, 0x7f}, {0xff, 0xff, 0xff, 0x7f}},
		-1000:         {{0xfc, 0x18}, {0xff, 0xff, 0xfc, 0x18}},
		-8388607:      {{0x80, 0x00, 0x01}},
		math.MaxInt64: {{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		math.MinInt64: {{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	})
	testInt(t, 4, map[Int32][][]byte{
		math.MinInt32: {{0x80, 0x00, 0x00, 0x00}},
		math.MaxInt32: {{0x7f, 0xff, 0xff, 0xff}},
	})
	testInt(t, 2, map[Int16][][]byte{
		math.MinInt16: {{0x80, 0x00}},
		math.MaxInt16: {{0x7f, 0xff}},
	})
	testInt(t, 1, map[Int8][][]byte{
		0:            {{0x00}},
		1:            {{0x01}},
		math.MaxInt8: {{0x7f}},
		math.MinInt8: {{0x80}},
		-1:           {{0xff}},
	})
}

func TestIntegerError(t *testing.T) {
	var value Int8
	assert.NoError(t, unmarshalInt(nil, 1, &value))
	assert.Error(t, unmarshalInt([]byte{0xff, 0xff}, 1, &value))
}

func TestInt64(t *testing.T) {
	var value Int64
	expected := []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	expectedValue := Int64(math.MaxInt64)
	assert.NoError(t, value.UnmarshalBinary(expected))
	assert.Equal(t, expectedValue, value)
	data, err := value.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func TestInt32(t *testing.T) {
	var value Int32
	expected := []byte{0x7f, 0xff, 0xff, 0xff}
	expectedValue := Int32(math.MaxInt32)
	assert.NoError(t, value.UnmarshalBinary(expected))
	assert.Equal(t, expectedValue, value)
	data, err := value.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func TestInt16(t *testing.T) {
	var value Int16
	expected := []byte{0x7f, 0xff}
	expectedValue := Int16(math.MaxInt16)
	assert.NoError(t, value.UnmarshalBinary(expected))
	assert.Equal(t, expectedValue, value)
	data, err := value.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func TestInt8(t *testing.T) {
	var value Int8
	expected := []byte{0x7f}
	expectedValue := Int8(math.MaxInt8)
	assert.NoError(t, value.UnmarshalBinary(expected))
	assert.Equal(t, expectedValue, value)
	data, err := value.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}

func testInt[T signedInt](t *testing.T, maxBytes int, fixtures map[T][][]byte) {
	for expected, variants := range fixtures {
		var value T
		for _, variant := range variants {
			assert.NoError(t, unmarshalInt(variant, maxBytes, &value))
			assert.Equal(t, expected, value, "UnmarshalBinary")
		}
		actual, err := marshalInt(maxBytes, expected)
		assert.NoError(t, err)
		assert.Equal(t, variants[0], actual, "MarshalBinary")
	}
}

package primitive

import (
	"fmt"
)

type Int8 int8

func (n *Int8) UnmarshalBinary(data []byte) error { return unmarshalInt(data, 1, n) }
func (n *Int8) MarshalBinary() ([]byte, error)    { return marshalInt(1, *n) }

type Int16 int16

func (n *Int16) UnmarshalBinary(data []byte) error { return unmarshalInt(data, 2, n) }
func (n *Int16) MarshalBinary() ([]byte, error)    { return marshalInt(2, *n) }

type Int32 int32

func (n *Int32) UnmarshalBinary(data []byte) error { return unmarshalInt(data, 4, n) }
func (n *Int32) MarshalBinary() ([]byte, error)    { return marshalInt(4, *n) }

type Int64 int64

func (n *Int64) UnmarshalBinary(data []byte) error { return unmarshalInt(data, 8, n) }
func (n *Int64) MarshalBinary() ([]byte, error)    { return marshalInt(8, *n) }

type signedInt interface {
	~int8 | ~int16 | ~int32 | ~int64
}

func unmarshalInt[Int signedInt](data []byte, maxBytes int, value *Int) error {
	if len(data) == 0 {
		return nil
	} else if len(data) > maxBytes {
		return fmt.Errorf("the value is too large, expected at most %d bytes, got %d", maxBytes, len(data))
	}
	var n Int
	var index int
	if data[0] == 0x00 || data[0] == 0xff {
		for index = 0; index < len(data)-1; index++ {
			if data[index] != data[0] || data[index+1]>>7 != data[0]>>7 {
				break
			}
		}
		data = data[index:]
	}
	if data[0] > 0x7f {
		n = -1
	}
	for index = 0; index < len(data); index++ {
		n <<= 8
		n ^= Int(data[index])
	}
	*value = n
	return nil
}

func marshalInt[Int signedInt](maxBytes int, value Int) (data []byte, err error) {
	data = make([]byte, maxBytes)
	var index int
	for index = len(data) - 1; index >= 0; index-- {
		data[index] = byte(value)
		value >>= 8
	}
	for index = 0; index < len(data)-1; index++ {
		if data[index] != data[0] || data[index+1]>>7 != data[0]>>7 {
			break
		}
	}
	data = data[index:]
	return
}

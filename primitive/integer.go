package primitive

import (
	"fmt"
	"unsafe"
)

type signedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func UnmarshalInt[Int signedInt](value *Int) Unmarshaler {
	size := int(unsafe.Sizeof(*value))
	return func(data []byte) error {
		if len(data) == 0 {
			return nil
		} else if len(data) > size {
			return fmt.Errorf("the value is too large, expected at most %d bytes, got %d", size, len(data))
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
}

func MarshalInt[Int signedInt](value Int) Marshaller {
	size := int(unsafe.Sizeof(value))
	return func() (data []byte, err error) {
		data = make([]byte, size)
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
}

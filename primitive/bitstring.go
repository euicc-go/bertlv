package primitive

import (
	"errors"
)

type BitString []bool

func (bits *BitString) Read(flags ...*bool) {
	n := min(len(flags), len(*bits))
	for index := 0; index < n; index++ {
		if flags[index] == nil {
			continue
		}
		*flags[index] = (*bits)[index]
	}
}

func (bits *BitString) UnmarshalBinary(data []byte) error {
	paddingBits := int(data[0])
	if paddingBits > 7 || len(data) == 1 && paddingBits > 0 || data[len(data)-1]&((1<<data[0])-1) != 0 {
		return errors.New("invalid padding bits")
	}
	bitLength := (len(data)-1)*8 - paddingBits

	data = data[1:]
	flags := make([]bool, bitLength)
	var x byte
	for i := 0; i < bitLength; i++ {
		x = 7 - byte(i%8)
		flags[i] = data[i/8]>>x&0b1 == 0b1
	}
	*bits = flags
	return nil
}

func (bits *BitString) MarshalBinary() (data []byte, err error) {
	data = make([]byte, len(*bits)/8+2)
	paddingBits := 8 - len(*bits)%8
	data[0] = byte(paddingBits)

	values := *bits
	var x byte
	var offset int
	for index := 0; index < len(values); index++ {
		offset = index / 8
		if x = 7 - byte(index%8); values[index] {
			data[1+offset] |= 0b1 << x
		}
	}
	return data, nil
}

func (bits *BitString) String() string {
	values := make([]byte, len(*bits))
	for index, bit := range *bits {
		if values[index] = '0'; bit {
			values[index] = '1'
		}
	}
	return string(values)
}

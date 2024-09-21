package primitive

type Boolean bool

func (b *Boolean) UnmarshalBinary(data []byte) error {
	*b = len(data) == 1 && data[0] == 0xff
	return nil
}

func (b *Boolean) MarshalBinary() ([]byte, error) {
	if b == nil || *b == false {
		return []byte{0x00}, nil
	}
	return []byte{0xff}, nil
}

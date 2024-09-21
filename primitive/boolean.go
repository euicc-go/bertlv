package primitive

func UnmarshalBool(value *bool) Unmarshaler {
	return func(data []byte) error {
		*value = len(data) == 1 && data[0] == 0xff
		return nil
	}
}

func MarshalBool(value bool) Marshaller {
	return func() ([]byte, error) {
		if value {
			return []byte{0xff}, nil
		}
		return []byte{0x00}, nil
	}
}

package primitive

import "math/big"

func UnmarshalBigInt(value *big.Int) Unmarshaler {
	return func(data []byte) error {
		value.SetBytes(data)
		return nil
	}
}

func MarshalBigInt(value *big.Int) Marshaller {
	return func() ([]byte, error) {
		return value.Bytes(), nil
	}
}

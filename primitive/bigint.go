package primitive

import "math/big"

type BigInt struct{ *big.Int }

func (v *BigInt) UnmarshalBinary(data []byte) error {
	v.Int = new(big.Int).SetBytes(data)
	return nil
}

func (v *BigInt) MarshalBinary() ([]byte, error) {
	return v.Int.Bytes(), nil
}

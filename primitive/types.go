package primitive

type Unmarshaler func([]byte) error

func (u Unmarshaler) UnmarshalBinary(data []byte) error {
	return u(data)
}

type Marshaller func() ([]byte, error)

func (m Marshaller) MarshalBinary() ([]byte, error) {
	return m()
}

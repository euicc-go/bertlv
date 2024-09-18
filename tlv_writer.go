package bertlv

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
)

func (tlv *TLV) WriteTo(w io.Writer) (n int64, err error) {
	switch {
	case len(tlv.Value) > 0 && tlv.Tag.Constructed():
		return 0, errors.New("tlv: constructed tag cannot have value")
	case len(tlv.Children) > 0 && tlv.Tag.Primitive():
		return 0, errors.New("tlv: primitive tag cannot have children")
	}
	value := tlv.Value
	if tlv.Tag.Constructed() {
		var buf bytes.Buffer
		for _, child := range tlv.Children {
			if child == nil {
				continue
			}
			if _, err = child.WriteTo(&buf); err != nil {
				return
			}
		}
		value = buf.Bytes()
	}
	if len(value) > 0xffff {
		err = errors.New("tlv: invalid value or children")
		return
	}
	var buf bytes.Buffer
	buf.Write(tlv.Tag)
	buf.Write(marshalLength(uint16(len(value))))
	buf.Write(value)
	return buf.WriteTo(w)
}

func (tlv *TLV) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	_, err := tlv.WriteTo(encoder)
	_ = encoder.Close()
	return buf.Bytes(), err
}

func (tlv *TLV) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	_, err := tlv.WriteTo(&buf)
	return buf.Bytes(), err
}

func (tlv *TLV) Bytes() []byte {
	encoded, err := tlv.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return encoded
}

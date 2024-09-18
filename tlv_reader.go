package bertlv

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

func (tlv *TLV) ReadFrom(r io.Reader) (n int64, err error) {
	r = &countReader{Reader: r, Length: &n}
	var t TLV
	var length uint16
	if _, err = t.Tag.ReadFrom(r); err != nil {
		return
	}
	if length, err = readLength(r); err != nil {
		return n, fmt.Errorf("tag %02X: invalid length encoding, offset: %d\n%w", t.Tag, r.(*countReader).Offset(), err)
	}
	if t.Tag.Constructed() {
		var _n int64
		var child TLV
		for index := uint16(0); index < length; index += uint16(_n) {
			if _n, err = child.ReadFrom(r); err != nil {
				return n, fmt.Errorf("tag %02X: invalid child object\n%w", t.Tag, err)
			}
			t.Children = append(t.Children, &child)
		}
	} else if length > 0 {
		t.Value = make([]byte, length)
		if _, err = io.ReadAtLeast(r, t.Value, len(t.Value)); err != nil {
			return n, fmt.Errorf("tag %02X: invalid length encoding\n%w", t.Tag, err)
		}
	}
	*tlv = t
	return
}

func (tlv *TLV) UnmarshalText(text []byte) error {
	_, err := tlv.ReadFrom(base64.NewDecoder(base64.StdEncoding, bytes.NewReader(text)))
	return err
}

func (tlv *TLV) UnmarshalBinary(data []byte) error {
	_, err := tlv.ReadFrom(bytes.NewReader(data))
	return err
}

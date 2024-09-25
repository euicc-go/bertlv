package bertlv

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

func (tlv *TLV) Len() int {
	n := contentLength(tlv)
	if n < 128 {
		n += 1
	} else if n < 256 {
		n += 2
	} else {
		n += 3
	}
	n += len(tlv.Tag)
	return n
}

func (tlv *TLV) WriteTo(w io.Writer) (n int64, err error) {
	w = &countWriter{Writer: w, Written: &n}
	length := contentLength(tlv)
	switch {
	case len(tlv.Value) > 0 && tlv.Tag.Constructed():
		return 0, errors.New("tlv: constructed tag cannot have value")
	case len(tlv.Children) > 0 && tlv.Tag.Primitive():
		return 0, errors.New("tlv: primitive tag cannot have children")
	case length > 0xffff:
		return 0, fmt.Errorf("tlv: length exceeds maximum (%d), got %d", 0xffff, length)
	}
	if _, err = w.Write(tlv.Tag); err != nil {
		return
	}
	if _, err = w.Write(marshalLength(uint16(length))); err != nil {
		return
	}
	if tlv.Tag.Primitive() {
		_, err = w.Write(tlv.Value)
		return
	}
	for _, child := range tlv.Children {
		if child == nil {
			continue
		}
		if _, err = child.WriteTo(w); err != nil {
			return
		}
	}
	return
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
	buf.Grow(tlv.Len())
	_, err := tlv.WriteTo(&buf)
	return buf.Bytes(), err
}

func (tlv *TLV) MarshalBERTLV() (*TLV, error) {
	return tlv.Clone(), nil
}

func (tlv *TLV) Clone() *TLV {
	var cloned TLV
	cloned.Tag = make(Tag, len(tlv.Tag))
	copy(cloned.Tag, tlv.Tag)
	if tlv.Tag.Primitive() {
		cloned.Value = make([]byte, len(tlv.Value))
		copy(cloned.Value, tlv.Value)
		return &cloned
	}
	cloned.Children = make([]*TLV, 0, len(tlv.Children))
	for _, child := range tlv.Children {
		if child == nil {
			continue
		}
		cloned.Children = append(cloned.Children, child.Clone())
	}
	return &cloned
}

func (tlv *TLV) Bytes() []byte {
	encoded, err := tlv.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return encoded
}

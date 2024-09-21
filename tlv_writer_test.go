package bertlv

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTLV_Len(t *testing.T) {
	var tlv TLV
	tlv.Tag = Primitive.ContextSpecific(0)
	tlv.Value = make([]byte, 127)
	assert.Equal(t, 129, tlv.Len())
	tlv.Value = make([]byte, 255)
	assert.Equal(t, 258, tlv.Len())
	tlv.Value = make([]byte, 256)
	assert.Equal(t, 260, tlv.Len())
}

func TestTLV_MarshalText(t *testing.T) {
	tlv := NewValue(Primitive.ContextSpecific(0), []byte{0xff})
	text, err := tlv.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, []byte("gAH/"), text)
	err = tlv.UnmarshalText(text)
	assert.NoError(t, err)
	assert.Equal(t, []byte{0xff}, tlv.Value)
}

func TestTLV_MarshalBinary(t *testing.T) {
	tlv := NewChildren(
		Constructed.ContextSpecific(0),
		NewValue(Primitive.Universal(0), []byte{0xff}),
	)
	binary, err := tlv.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, []byte{0xa0, 0x03, 0x00, 0x01, 0xff}, binary)
	err = tlv.UnmarshalBinary(binary)
	assert.NoError(t, err)
	assert.Equal(t, []byte{0xff}, tlv.Children[0].Value)
}

func TestTLV_MarshalJSON(t *testing.T) {
	tlv := NewValue(Primitive.ContextSpecific(0), []byte{0xff})
	encoded, err := json.Marshal(tlv)
	assert.NoError(t, err)
	assert.Equal(t, `"gAH/"`, string(encoded))
	assert.NoError(t, json.Unmarshal(encoded, &tlv))
}

func TestTLV_LargeValue(t *testing.T) {
	tlv := NewChildren(
		Constructed.ContextSpecific(0),
		NewValue(Primitive.ContextSpecific(0), make([]byte, 0x10000)),
	)
	_, err := tlv.MarshalBinary()
	assert.EqualError(t, err, "tlv: invalid value or children")
}

func TestTLV_InvalidConstructedTag(t *testing.T) {
	tlv := NewChildren(
		Constructed.ContextSpecific(0),
		NewValue(Primitive.ContextSpecific(1), []byte{0x01}),
	)
	tlv.Tag = Primitive.ContextSpecific(0)
	assert.Panics(t, func() { tlv.Bytes() })
}

func TestTLV_InvalidValueTag(t *testing.T) {
	tlv := NewValue(
		Primitive.ContextSpecific(0),
		[]byte{0xff},
	)
	tlv.Tag = Constructed.ContextSpecific(0)
	assert.Panics(t, func() { tlv.Bytes() })
}

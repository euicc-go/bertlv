package bertlv

import (
	"github.com/euicc-go/bertlv/primitive"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewValue(t *testing.T) {
	tlv := NewValue(Primitive.ContextSpecific(0), []byte{0xff})
	assert.Equal(t, Tag{0x80}, tlv.Tag)
	assert.Len(t, tlv.Value, 1)
	assert.Len(t, tlv.Children, 0)
	assert.Equal(t, []byte{0x80, 0x01, 0xff}, tlv.Bytes())
	assert.Panics(t, func() { NewValue(Constructed.ContextSpecific(0), nil) })
}

func TestNewChildren(t *testing.T) {
	tlv := NewChildren(Constructed.ContextSpecific(0))
	assert.Equal(t, Tag{0xa0}, tlv.Tag)
	assert.Len(t, tlv.Value, 0)
	assert.Equal(t, []byte{0xa0, 0x00}, tlv.Bytes())
	assert.Panics(t, func() { NewChildren(Primitive.ContextSpecific(0)) })
}

func TestNewChildrenIter(t *testing.T) {
	tlv := NewChildrenIter(Constructed.ContextSpecific(0), func(yield func(*TLV) bool) {
		yield(NewValue(Primitive.ContextSpecific(0), []byte{0xff}))
		yield(NewValue(Primitive.ContextSpecific(1), []byte{0xff}))
	})
	assert.Equal(t, Tag{0xa0}, tlv.Tag)
	assert.Len(t, tlv.Value, 0)
	assert.Equal(t, []byte{0xa0, 0x6, 0x80, 0x1, 0xff, 0x81, 0x1, 0xff}, tlv.Bytes())
}

func TestMarshalValue(t *testing.T) {
	value := primitive.Int8(-1)
	tlv, err := MarshalValue(Primitive.ContextSpecific(0), &value)
	assert.NoError(t, err)
	assert.Equal(t, Tag{0x80}, tlv.Tag)
	assert.Len(t, tlv.Value, 1)
	assert.Len(t, tlv.Children, 0)
	assert.Equal(t, []byte{0x80, 0x01, 0xff}, tlv.Bytes())
}

func TestTLV_MarshalValue(t *testing.T) {
	tlv := NewValue(Primitive.ContextSpecific(0), nil)
	value := primitive.Int8(-1)
	assert.NoError(t, tlv.MarshalValue(&value))
	assert.Equal(t, []byte{0x80, 0x01, 0xff}, tlv.Bytes())
	value = 0
	assert.NoError(t, tlv.UnmarshalValue(&value))
	assert.Equal(t, primitive.Int8(-1), value)
}

func TestTLV_MarshalValueError(t *testing.T) {
	tlv := NewChildren(Constructed.ContextSpecific(0))
	value := primitive.Int8(-1)
	assert.Error(t, tlv.MarshalValue(&value))
	assert.Error(t, tlv.UnmarshalValue(&value))
}

func TestTLV_String(t *testing.T) {
	tlv := NewValue(Primitive.ContextSpecific(0), []byte{0xff})
	assert.Equal(t, "[0] (1 byte)", tlv.String())
	tlv = NewChildren(Constructed.ContextSpecific(0))
	assert.Equal(t, "[0] (0 elem)", tlv.String())
}

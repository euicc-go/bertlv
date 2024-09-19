package bertlv

import (
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

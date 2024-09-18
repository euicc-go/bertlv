package bertlv

import (
	"fmt"
)

type TLV struct {
	Tag      Tag
	Value    []byte
	Children []*TLV
}

func (tlv *TLV) String() string {
	if tlv.Tag.Primitive() {
		return fmt.Sprintf("%s (%d byte)", tlv.Tag.String(), len(tlv.Value))
	}
	return fmt.Sprintf("%s (%d elem)", tlv.Tag.String(), len(tlv.Children))
}

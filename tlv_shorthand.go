package bertlv

import "iter"

func NewValue(tag Tag, value []byte) *TLV {
	if tag.Constructed() {
		panic("tlv: constructed tag cannot have value")
	}
	return &TLV{Tag: tag, Value: value}
}

func NewChildren(tag Tag, children ...*TLV) *TLV {
	if tag.Primitive() {
		panic("tlv: primitive tag cannot have children")
	}
	return &TLV{Tag: tag, Children: children}
}

func NewChildrenIter(tag Tag, children iter.Seq[*TLV]) *TLV {
	var elements []*TLV
	for child := range children {
		elements = append(elements, child)
	}
	return NewChildren(tag, elements...)
}

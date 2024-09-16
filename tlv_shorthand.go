package bertlv

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

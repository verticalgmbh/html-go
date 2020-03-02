package html

// Text text data
type Text struct {
	Data string
}

// Clone clones this text section
func (text *Text) Clone() Node {
	return &Text{Data: text.Data}
}

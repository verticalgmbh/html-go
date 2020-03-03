package html

// Attribute attribute in a html tag
type Attribute struct {
	Name  string
	Value *Text
}

// Clone clones this attribute
func (attr *Attribute) Clone() *Attribute {
	return &Attribute{
		Name:  attr.Name,
		Value: attr.Value}
}

package html

// Tag tag in a html document
type Tag struct {
	Name       string
	Attributes []*Attribute
	Children   []Node
}

// Clone clones this tag
func (tag *Tag) Clone() Node {
	var attributes []*Attribute
	var children []Node

	for _, attribute := range tag.Attributes {
		attributes = append(attributes, attribute.Clone())
	}

	for _, child := range tag.Children {
		children = append(children, child.Clone())
	}

	return &Tag{
		Name:       tag.Name,
		Attributes: attributes,
		Children:   children}
}

// GetAttribute get an attribute by name
func (tag *Tag) GetAttribute(name string) *Attribute {
	for _, attr := range tag.Attributes {
		if attr.Name == name {
			return attr
		}
	}
	return nil
}

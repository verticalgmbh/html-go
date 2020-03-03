package html

import "strings"

// Text text data
type Text struct {
	Tokens []*TextToken
}

// Clone clones this text section
func (text *Text) Clone() Node {
	return &Text{Tokens: text.Tokens}
}

// String get string representation of this token
func (text *Text) String() string {
	var builder strings.Builder
	for _, token := range text.Tokens {
		builder.WriteString(token.String())
	}
	return builder.String()
}

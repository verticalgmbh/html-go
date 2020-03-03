package html

var tokentorune = map[string]rune{
	"quot": '"',
	"apos": '\'',
	"amp":  '&',
	"lt":   '<',
	"gt":   '>'}

var runetotoken = map[rune]string{
	'"':  "quot",
	'\'': "apos",
	'&':  "amp",
	'<':  "lt",
	'>':  "gt"}

// TextToken tokens making up a html text
type TextToken struct {
	IsSpecial bool
	Data      string
}

// Token token without any text translation
func (token *TextToken) Token() string {
	if token.IsSpecial {
		return "&" + token.Data + ";"
	}
	return token.Data
}

// String get string representation of this token
func (token *TextToken) String() string {
	if token.IsSpecial {
		char, ok := tokentorune[token.Data]
		if ok {
			return string(char)
		}

		return "&" + token.Data + ";"
	}
	return token.Data
}

// IsSpecial determines whether the token is a known special
func IsSpecial(token string) bool {
	_, ok := tokentorune[token]
	return ok
}

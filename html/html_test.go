package html

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	file, err := os.Open("../testdata/template1.html")
	assert.NoError(t, err)

	document, err := Parse(file)
	assert.NoError(t, err)

	assert.NotNil(t, document.HTML)
	assert.Equal(t, "html", document.HTML.Name)
}

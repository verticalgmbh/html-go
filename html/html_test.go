package html

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	file, err := os.Open("../testdata/template1.html")
	require.NoError(t, err)

	document, err := Parse(file)
	require.NoError(t, err)

	require.NotNil(t, document.HTML)
	require.Equal(t, "html", document.HTML.Name)
}

func TestSpecialCharacters(t *testing.T) {
	file, err := os.Open("../testdata/ifcondition.html")
	require.NoError(t, err)

	document, err := Parse(file)
	require.NoError(t, err)

	iftags := document.HTML.GetTagChild("body").GetTagChildren("if")
	require.Equal(t, 2, len(iftags))

	require.Equal(t, "##TRIGGER.COUNT## < 5", iftags[0].GetAttribute("condition").Value.String())
	require.Equal(t, "##TRIGGER.COUNT## >= 5", iftags[1].GetAttribute("condition").Value.String())

	require.NotNil(t, document.HTML)
	require.Equal(t, "html", document.HTML.Name)
}

func TestReencode(t *testing.T) {
	file, err := os.Open("../testdata/ifcondition.html")
	require.NoError(t, err)

	document, err := Parse(file)
	require.NoError(t, err)

	file, err = os.Open("../testdata/ifcondition.html")
	require.NoError(t, err)

	var input bytes.Buffer
	io.Copy(&input, file)

	var encoded bytes.Buffer
	Write(document, &encoded)

	require.Equal(t, strings.ReplaceAll(input.String(), "\r", ""), strings.ReplaceAll(encoded.String(), "\r", ""))
}

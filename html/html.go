package html

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

func parseDocumentParent(reader *bufio.Reader, document *Document) error {
	var err error

	reader.ReadBytes(byte('<'))
	data, err := reader.ReadBytes(byte('>'))

	if err != nil {
		return err
	}
	if len(data) == 0 {
		return errors.New("Invalid tag")
	}

	document.HTML, err = parseTag(data)
	if err != nil {
		return err
	}

	err = parseTagParent(reader, document.HTML)

	if err != nil {
		return err
	}

	return nil
}

func parseTagParent(reader *bufio.Reader, parent *Tag) error {
	var err error
	var tag *Tag
	for err != io.EOF {
		data, err := reader.ReadBytes(byte('<'))
		if err != nil {
			return err
		}
		if len(data) > 0 {
			parent.Children = append(parent.Children, &Text{
				Data: string(data[:len(data)-1])})
		}

		data, err = reader.ReadBytes(byte('>'))
		if err != nil {
			return err
		}
		if len(data) == 0 {
			return errors.New("Invalid tag")
		}

		if data[0] == byte('/') {
			if string(data[1:len(data)-1]) == parent.Name {
				// parent tag was closed
				return nil
			} else {
				return errors.New("Invalid tag close")
			}
		}

		// comment
		if data[0] == byte('!') {
			continue
		}

		tag, err = parseTag(data)
		if err != nil {
			return err
		}

		if data[len(data)-2] != '/' {
			switch strings.ToLower(tag.Name) {
			case "meta", "br":
			default:
				err = parseTagParent(reader, tag)
				if err != nil {
					return err
				}
			}
		}
		parent.Children = append(parent.Children, tag)
	}

	return nil
}

func parseTag(data []byte) (*Tag, error) {
	var builder strings.Builder

	tag := Tag{}
	attr := &Attribute{}
	state := 0
	for offset := 0; offset < len(data); offset++ {
		switch state {
		case 0:
			switch data[offset] {
			case ' ', '>':
				if builder.Len() == 0 {
					return nil, errors.New("Invalid tag name")
				}
				tag.Name = builder.String()
				builder.Reset()
				state = 1
			default:
				builder.WriteByte(data[offset])
			}
		case 1:
			switch data[offset] {
			case '=':
				attr = &Attribute{
					Name: strings.TrimSpace(builder.String())}
				builder.Reset()
				state = 2
			default:
				builder.WriteByte(data[offset])
			}
		case 2:
			switch data[offset] {
			case '"':
				state = 3
			}
		case 3:
			switch data[offset] {
			case '"':
				attr.Value = strings.TrimSpace(builder.String())
				builder.Reset()
				tag.Attributes = append(tag.Attributes, attr)
				attr = nil
				state = 1
			default:
				builder.WriteByte(data[offset])
			}
		}
	}
	return &tag, nil
}

// Parse parses a html document
func Parse(reader io.Reader) (*Document, error) {
	buffer := bufio.NewReader(reader)
	buffer.ReadBytes(byte('<'))
	data, err := buffer.ReadBytes(byte('>'))

	if len(data) == 0 {
		return nil, errors.New("Invalid tag")
	}

	document := &Document{}
	if data[0] == byte('!') {
		// TODO: analyse DOCTYPE
		document.Type = "html"
		err = parseDocumentParent(buffer, document)
	} else {
		document.HTML, err = parseTag(data)
		err = parseTagParent(buffer, document.HTML)
	}

	if err != nil && err != io.EOF {
		return nil, err
	}

	return document, nil
}

func writeNode(node Node, writer io.Writer) {
	switch node.(type) {
	case *Text:
		text := node.(*Text)
		io.WriteString(writer, text.Data)
	case *Tag:
		tag := node.(*Tag)
		io.WriteString(writer, fmt.Sprintf("<%s", tag.Name))
		for _, attr := range tag.Attributes {
			io.WriteString(writer, fmt.Sprintf(" %s=\"%s\"", attr.Name, attr.Value))
		}

		switch tag.Name {
		case "meta", "br":
			io.WriteString(writer, ">")
			return
		}

		if len(tag.Children) == 0 {
			io.WriteString(writer, "/>")
			return
		}
		io.WriteString(writer, ">")

		for _, child := range tag.Children {
			writeNode(child, writer)
		}

		io.WriteString(writer, fmt.Sprintf("</%s>", tag.Name))
	}
}

// Write writes a html document to a writer
func Write(document *Document, writer io.Writer) {
	if len(document.Type) > 0 {
		io.WriteString(writer, fmt.Sprintf("<!DOCTYPE %s>\n", document.Type))
	}

	writeNode(document.HTML, writer)
}

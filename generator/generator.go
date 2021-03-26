package generator

import (
	"encoding/xml"
	"io"
	"log"
)

type Generator struct {
	Tree  Element
	stack Stack
}

func NewGenerator(reader io.Reader) Generator {
	g := Generator{}
	g.build(reader)
	return g
}

func (g *Generator) build(reader io.Reader) {
	decoder := xml.NewDecoder(reader)
	for {
		token, err := decoder.Token()
		if token == nil || err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error decoding token: %s", err)
		}
		g.handle(token)
	}
}

func (g *Generator) handle(token xml.Token) {
	switch el := token.(type) {
	case xml.StartElement:
		g.stack.Push(&Element{
			Tag:  el.Name.Local,
			Attr: attrToMap(el.Attr),
		})
	case xml.EndElement:
		element := g.stack.Pop()
		if g.stack.Empty() {
			g.Tree = *element
		} else {
			peek := g.stack.Peek()
			peek.Children = append(peek.Children, *element)
		}
	}
}

func attrToMap(attr []xml.Attr) map[string]string {
	res := make(map[string]string, len(attr))
	for _, item := range attr {
		res[item.Name.Local] = item.Value
	}
	return res
}

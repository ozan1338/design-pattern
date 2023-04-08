package main

import (
	"fmt"
	"strings"
)

const (
	IndentSize = 2
)

type HTMLElement struct {
	name, text string
	elements   []HTMLElement
}

func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", IndentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))

	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", IndentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}

	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HTMLElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root:     HTMLElement{rootName, "", []HTMLElement{}},
	}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(chilName, childText string) {
	e := HTMLElement{chilName, childText, []HTMLElement{}}

	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(chilName, childText string) *HtmlBuilder {
	e := HTMLElement{chilName, childText, []HTMLElement{}}

	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {
	// hello := "hello"

	// sb := strings.Builder{}
	// sb.WriteString("<p>")
	// sb.WriteString(hello)
	// sb.WriteString("</p>")
	// log.Println(sb.String())

	// words := []string{"hello", "world"}
	// sb.Reset()
	// //<ul><li>....</li><li>....</li><li>....</li></ul>
	// sb.WriteString("<ul>")
	// for _, word := range words {
	// 	sb.WriteString("<li>")
	// 	sb.WriteString(word)
	// 	sb.WriteString("</li>")
	// }
	// sb.WriteString("</ul>")
	// log.Println(sb.String())

	// b := NewHtmlBuilder("p")
	// b.root.text = "hello"
	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li", "siang").AddChildFluent("li", "hari")
	fmt.Println(b.String())
}

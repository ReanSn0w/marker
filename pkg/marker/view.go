package marker

import (
	"context"
	"fmt"
	"io"

	"github.com/ReanSn0w/gew/pkg/view"
)

func NewNode(name string, attr AttributeStorage, content interface{}) view.View {
	return view.External(&Node{
		Name:    name,
		Attr:    attr,
		Content: content,
	})
}

// Базовая структура для построения HTML элементов
type Node struct {
	Name    string
	Attr    AttributeStorage
	Content interface{}
}

func (n *Node) Build(ctx context.Context, wr io.Writer) {
	builder := GetBuilder(ctx)

	switch n.Name {
	case "img", "br", "hr", "input", "meta", "link", "base", "area", "col", "param", "source", "track", "wbr":
		wr.Write([]byte(fmt.Sprintf("\n<%s%s>", n.Name, n.Attr.String())))
		return
	default:
		wr.Write([]byte(fmt.Sprintf("\n<%s%s>", n.Name, n.Attr.String())))
	}

	if item, ok := n.Content.(view.View); ok {
		builder.Build(ctx, wr, item)
	} else {
		if n.Content != nil {
			wr.Write([]byte("\n"))
			wr.Write([]byte(fmt.Sprintf("%v", n.Content)))
		}
	}

	wr.Write([]byte(fmt.Sprintf("\n</%s>", n.Name)))
}

// MARK: - Инициализация тегов

func A(attr AttributeStorage, content interface{}) view.View {
	return NewNode("a", attr, content)
}

func Abbr(attr AttributeStorage, content interface{}) view.View {
	return NewNode("abbr", attr, content)
}

func Address(attr AttributeStorage, content interface{}) view.View {
	return NewNode("address", attr, content)
}

func Area(attr AttributeStorage, content interface{}) view.View {
	return NewNode("area", attr, content)
}

func Article(attr AttributeStorage, content interface{}) view.View {
	return NewNode("article", attr, content)
}

func Aside(attr AttributeStorage, content interface{}) view.View {
	return NewNode("aside", attr, content)
}

func Audio(attr AttributeStorage, content interface{}) view.View {
	return NewNode("audio", attr, content)
}

func B(attr AttributeStorage, content interface{}) view.View {
	return NewNode("b", attr, content)
}

func Base(attr AttributeStorage, content interface{}) view.View {
	return NewNode("base", attr, content)
}

func Bdi(attr AttributeStorage, content interface{}) view.View {
	return NewNode("bdi", attr, content)
}

func Bdo(attr AttributeStorage, content interface{}) view.View {
	return NewNode("bdo", attr, content)
}

func Blockquote(attr AttributeStorage, content interface{}) view.View {
	return NewNode("blockquote", attr, content)
}

func Body(attr AttributeStorage, content interface{}) view.View {
	return NewNode("body", attr, content)
}

func Br(attr AttributeStorage, content interface{}) view.View {
	return NewNode("br", attr, content)
}

func Button(attr AttributeStorage, content interface{}) view.View {
	return NewNode("button", attr, content)
}

func Canvas(attr AttributeStorage, content interface{}) view.View {
	return NewNode("canvas", attr, content)
}

func Caption(attr AttributeStorage, content interface{}) view.View {
	return NewNode("caption", attr, content)
}

func Cite(attr AttributeStorage, content interface{}) view.View {
	return NewNode("cite", attr, content)
}

func Code(attr AttributeStorage, content interface{}) view.View {
	return NewNode("code", attr, content)
}

func Col(attr AttributeStorage, content interface{}) view.View {
	return NewNode("col", attr, content)
}

func Colgroup(attr AttributeStorage, content interface{}) view.View {
	return NewNode("colgroup", attr, content)
}

func Commend(attr AttributeStorage, content interface{}) view.View {
	return NewNode("commend", attr, content)
}

func Comment(attr AttributeStorage, content interface{}) view.View {
	return NewNode("comment", attr, content)
}

func Details(attr AttributeStorage, content interface{}) view.View {
	return NewNode("details", attr, content)
}

func Dfn(attr AttributeStorage, content interface{}) view.View {
	return NewNode("dfn", attr, content)
}

func Dl(attr AttributeStorage, content interface{}) view.View {
	return NewNode("dl", attr, content)
}

func Dt(attr AttributeStorage, content interface{}) view.View {
	return NewNode("dt", attr, content)
}

func Em(attr AttributeStorage, content interface{}) view.View {
	return NewNode("em", attr, content)
}

func Embed(attr AttributeStorage, content interface{}) view.View {
	return NewNode("embed", attr, content)
}

func Fieldset(attr AttributeStorage, content interface{}) view.View {
	return NewNode("fieldset", attr, content)
}

func Figcaption(attr AttributeStorage, content interface{}) view.View {
	return NewNode("figcaption", attr, content)
}

func Figure(attr AttributeStorage, content interface{}) view.View {
	return NewNode("figure", attr, content)
}

func Footer(attr AttributeStorage, content interface{}) view.View {
	return NewNode("footer", attr, content)
}

func Form(attr AttributeStorage, content interface{}) view.View {
	return NewNode("form", attr, content)
}

func H1(attr AttributeStorage, content interface{}) view.View {
	return NewNode("h1", attr, content)
}

func H2(attr AttributeStorage, content interface{}) view.View {
	return NewNode("h2", attr, content)
}

func H3(attr AttributeStorage, content interface{}) view.View {
	return NewNode("h3", attr, content)
}

func H4(attr AttributeStorage, content interface{}) view.View {
	return NewNode("h4", attr, content)
}

func H5(attr AttributeStorage, content interface{}) view.View {
	return NewNode("h5", attr, content)
}

func H6(attr AttributeStorage, content interface{}) view.View {
	return NewNode("h6", attr, content)
}

func Head(attr AttributeStorage, content interface{}) view.View {
	return NewNode("head", attr, content)
}

func Header(attr AttributeStorage, content interface{}) view.View {
	return NewNode("header", attr, content)
}

func Hr(attr AttributeStorage, content interface{}) view.View {
	return NewNode("hr", attr, content)
}

func Html(attr AttributeStorage, content interface{}) view.View {
	return NewNode("html", attr, content)
}

func I(attr AttributeStorage, content interface{}) view.View {
	return NewNode("i", attr, content)
}

func Iframe(attr AttributeStorage, content interface{}) view.View {
	return NewNode("iframe", attr, content)
}

func Img(attr AttributeStorage, content interface{}) view.View {
	return NewNode("img", attr, content)
}

func Input(attr AttributeStorage, content interface{}) view.View {
	return NewNode("input", attr, content)
}

func Ins(attr AttributeStorage, content interface{}) view.View {
	return NewNode("ins", attr, content)
}

func Kbd(attr AttributeStorage, content interface{}) view.View {
	return NewNode("kbd", attr, content)
}

func Keygen(attr AttributeStorage, content interface{}) view.View {
	return NewNode("keygen", attr, content)
}

func Label(attr AttributeStorage, content interface{}) view.View {
	return NewNode("label", attr, content)
}

func Legend(attr AttributeStorage, content interface{}) view.View {
	return NewNode("legend", attr, content)
}

func Li(attr AttributeStorage, content interface{}) view.View {
	return NewNode("li", attr, content)
}

func Link(attr AttributeStorage, content interface{}) view.View {
	return NewNode("link", attr, content)
}

func Main(attr AttributeStorage, content interface{}) view.View {
	return NewNode("main", attr, content)
}

func Map(attr AttributeStorage, content interface{}) view.View {
	return NewNode("map", attr, content)
}

func Mark(attr AttributeStorage, content interface{}) view.View {
	return NewNode("mark", attr, content)
}

func Menu(attr AttributeStorage, content interface{}) view.View {
	return NewNode("menu", attr, content)
}

func Meta(attr AttributeStorage, content interface{}) view.View {
	return NewNode("meta", attr, content)
}

func Meter(attr AttributeStorage, content interface{}) view.View {
	return NewNode("meter", attr, content)
}

func Nav(attr AttributeStorage, content interface{}) view.View {
	return NewNode("nav", attr, content)
}

func Noscript(attr AttributeStorage, content interface{}) view.View {
	return NewNode("noscript", attr, content)
}

func Ol(attr AttributeStorage, content interface{}) view.View {
	return NewNode("ol", attr, content)
}

func Optgroup(attr AttributeStorage, content interface{}) view.View {
	return NewNode("optgroup", attr, content)
}

func Option(attr AttributeStorage, content interface{}) view.View {
	return NewNode("option", attr, content)
}

func Output(attr AttributeStorage, content interface{}) view.View {
	return NewNode("output", attr, content)
}

func P(attr AttributeStorage, content interface{}) view.View {
	return NewNode("p", attr, content)
}

func Param(attr AttributeStorage, content interface{}) view.View {
	return NewNode("param", attr, content)
}

func Pre(attr AttributeStorage, content interface{}) view.View {
	return NewNode("pre", attr, content)
}

func Progress(attr AttributeStorage, content interface{}) view.View {
	return NewNode("progress", attr, content)
}

func Q(attr AttributeStorage, content interface{}) view.View {
	return NewNode("q", attr, content)
}

func Rp(attr AttributeStorage, content interface{}) view.View {
	return NewNode("rp", attr, content)
}

func Rt(attr AttributeStorage, content interface{}) view.View {
	return NewNode("rt", attr, content)
}

func Ruby(attr AttributeStorage, content interface{}) view.View {
	return NewNode("ruby", attr, content)
}

func S(attr AttributeStorage, content interface{}) view.View {
	return NewNode("s", attr, content)
}

func Samp(attr AttributeStorage, content interface{}) view.View {
	return NewNode("samp", attr, content)
}

func Script(attr AttributeStorage, content interface{}) view.View {
	return NewNode("script", attr, content)
}

func Section(attr AttributeStorage, content interface{}) view.View {
	return NewNode("section", attr, content)
}

func Select(attr AttributeStorage, content interface{}) view.View {
	return NewNode("select", attr, content)
}

func Small(attr AttributeStorage, content interface{}) view.View {
	return NewNode("small", attr, content)
}

func Source(attr AttributeStorage, content interface{}) view.View {
	return NewNode("source", attr, content)
}

func Span(attr AttributeStorage, content interface{}) view.View {
	return NewNode("span", attr, content)
}

func Strong(attr AttributeStorage, content interface{}) view.View {
	return NewNode("strong", attr, content)
}

func Style(attr AttributeStorage, content interface{}) view.View {
	return NewNode("style", attr, content)
}

func Sub(attr AttributeStorage, content interface{}) view.View {
	return NewNode("sub", attr, content)
}

func Summary(attr AttributeStorage, content interface{}) view.View {
	return NewNode("summary", attr, content)
}

func Sup(attr AttributeStorage, content interface{}) view.View {
	return NewNode("sup", attr, content)
}

func Table(attr AttributeStorage, content interface{}) view.View {
	return NewNode("table", attr, content)
}

func Tbody(attr AttributeStorage, content interface{}) view.View {
	return NewNode("tbody", attr, content)
}

func Td(attr AttributeStorage, content interface{}) view.View {
	return NewNode("td", attr, content)
}

func Textarea(attr AttributeStorage, content interface{}) view.View {
	return NewNode("textarea", attr, content)
}

func Tfoot(attr AttributeStorage, content interface{}) view.View {
	return NewNode("tfoot", attr, content)
}

func Th(attr AttributeStorage, content interface{}) view.View {
	return NewNode("th", attr, content)
}

func Thead(attr AttributeStorage, content interface{}) view.View {
	return NewNode("thead", attr, content)
}

func Time(attr AttributeStorage, content interface{}) view.View {
	return NewNode("time", attr, content)
}

func Title(attr AttributeStorage, content interface{}) view.View {
	return NewNode("title", attr, content)
}

func Tr(attr AttributeStorage, content interface{}) view.View {
	return NewNode("tr", attr, content)
}

func Track(attr AttributeStorage, content interface{}) view.View {
	return NewNode("track", attr, content)
}

func Ul(attr AttributeStorage, content interface{}) view.View {
	return NewNode("ul", attr, content)
}

func Var(attr AttributeStorage, content interface{}) view.View {
	return NewNode("var", attr, content)
}

func Video(attr AttributeStorage, content interface{}) view.View {
	return NewNode("video", attr, content)
}

func Wbr(attr AttributeStorage, content interface{}) view.View {
	return NewNode("wbr", attr, content)
}

func Div(attr AttributeStorage, content interface{}) view.View {
	return NewNode("div", attr, content)
}

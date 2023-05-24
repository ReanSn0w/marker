package marker

import (
	"fmt"

	"github.com/ReanSn0w/gew/v2/pkg/view"
)

func Doctype() view.View {
	return view.External("<!DOCTYPE html>")
}

func Comment(val string) view.View {
	return view.External("<!-- " + val + " -->")
}

func Text(val string) view.View {
	return view.External(val)()
}

func Snippet(snippet string, values ...interface{}) view.View {
	return view.External(fmt.Sprintf(snippet, values...))
}

func Error(err error) view.View {
	return view.External(err)
}

// MARK: - HTML Tags

func A(content ...view.View) view.ModificationApplyer {
	return Container("a", content...)
}

func Abbr(content ...view.View) view.ModificationApplyer {
	return Container("abbr", content...)
}

func Address(content ...view.View) view.ModificationApplyer {
	return Container("address", content...)
}

func Area(content ...view.View) view.ModificationApplyer {
	return Container("area", content...)
}

func Article(content ...view.View) view.ModificationApplyer {
	return Container("article", content...)
}

func Aside(content ...view.View) view.ModificationApplyer {
	return Container("aside", content...)
}

func Audio(content ...view.View) view.ModificationApplyer {
	return Container("audio", content...)
}

func B(content ...view.View) view.ModificationApplyer {
	return Container("b", content...)
}

func Base(content ...view.View) view.ModificationApplyer {
	return Container("base", content...)
}

func Bdi(content ...view.View) view.ModificationApplyer {
	return Container("bdi", content...)
}

func Bdo(content ...view.View) view.ModificationApplyer {
	return Container("bdo", content...)
}

func Blockquote(content ...view.View) view.ModificationApplyer {
	return Container("blockquote", content...)
}

func Body(content ...view.View) view.ModificationApplyer {
	return Container("body", content...)
}

func Br(content ...view.View) view.ModificationApplyer {
	return Container("br", content...)
}

func Button(content ...view.View) view.ModificationApplyer {
	return Container("button", content...)
}

func Canvas(content ...view.View) view.ModificationApplyer {
	return Container("canvas", content...)
}

func Caption(content ...view.View) view.ModificationApplyer {
	return Container("caption", content...)
}

func Cite(content ...view.View) view.ModificationApplyer {
	return Container("cite", content...)
}

func Code(content ...view.View) view.ModificationApplyer {
	return Container("code", content...)
}

func Col(content ...view.View) view.ModificationApplyer {
	return Container("col", content...)
}

func Colgroup(content ...view.View) view.ModificationApplyer {
	return Container("colgroup", content...)
}

func Data(content ...view.View) view.ModificationApplyer {
	return Container("data", content...)
}

func Datalist(content ...view.View) view.ModificationApplyer {
	return Container("datalist", content...)
}

func Dd(content ...view.View) view.ModificationApplyer {
	return Container("dd", content...)
}

func Del(content ...view.View) view.ModificationApplyer {
	return Container("del", content...)
}

func Details(content ...view.View) view.ModificationApplyer {
	return Container("details", content...)
}

func Dfn(content ...view.View) view.ModificationApplyer {
	return Container("dfn", content...)
}

func Dialog(content ...view.View) view.ModificationApplyer {
	return Container("dialog", content...)
}

func Div(content ...view.View) view.ModificationApplyer {
	return Container("div", content...)
}

func Dl(content ...view.View) view.ModificationApplyer {
	return Container("dl", content...)
}

func Dt(content ...view.View) view.ModificationApplyer {
	return Container("dt", content...)
}

func Em(content ...view.View) view.ModificationApplyer {
	return Container("em", content...)
}

func Embed(content ...view.View) view.ModificationApplyer {
	return Container("embed", content...)
}

func Fieldset(content ...view.View) view.ModificationApplyer {
	return Container("fieldset", content...)
}

func Figcaption(content ...view.View) view.ModificationApplyer {
	return Container("figcaption", content...)
}

func Figure(content ...view.View) view.ModificationApplyer {
	return Container("figure", content...)
}

func Footer(content ...view.View) view.ModificationApplyer {
	return Container("footer", content...)
}

func Form(content ...view.View) view.ModificationApplyer {
	return Container("form", content...)
}

func H1(content ...view.View) view.ModificationApplyer {
	return Container("h1", content...)
}

func H2(content ...view.View) view.ModificationApplyer {
	return Container("h2", content...)
}

func H3(content ...view.View) view.ModificationApplyer {
	return Container("h3", content...)
}

func H4(content ...view.View) view.ModificationApplyer {
	return Container("h4", content...)
}

func H5(content ...view.View) view.ModificationApplyer {
	return Container("h5", content...)
}

func H6(content ...view.View) view.ModificationApplyer {
	return Container("h6", content...)
}

func Head(content ...view.View) view.ModificationApplyer {
	return Container("head", content...)
}

func Header(content ...view.View) view.ModificationApplyer {
	return Container("header", content...)
}

func Hgroup(content ...view.View) view.ModificationApplyer {
	return Container("hgroup", content...)
}

func Hr(content ...view.View) view.ModificationApplyer {
	return Container("hr", content...)
}

func Html(content ...view.View) view.ModificationApplyer {
	return Container("html", content...)
}

func I(content ...view.View) view.ModificationApplyer {
	return Container("i", content...)
}

func Iframe(content ...view.View) view.ModificationApplyer {
	return Container("iframe", content...)
}

func Img() view.ModificationApplyer {
	return Inline("img")
}

func Input() view.ModificationApplyer {
	return Inline("input")
}

func Ins(content ...view.View) view.ModificationApplyer {
	return Container("ins", content...)
}

func Kbd(content ...view.View) view.ModificationApplyer {
	return Container("kbd", content...)
}

func Label(content ...view.View) view.ModificationApplyer {
	return Container("label", content...)
}

func Legend(content ...view.View) view.ModificationApplyer {
	return Container("legend", content...)
}

func Li(content ...view.View) view.ModificationApplyer {
	return Container("li", content...)
}

func Link() view.ModificationApplyer {
	return Inline("link")
}

func Main(content ...view.View) view.ModificationApplyer {
	return Container("main", content...)
}

func Map(content ...view.View) view.ModificationApplyer {
	return Container("map", content...)
}

func Mark(content ...view.View) view.ModificationApplyer {
	return Container("mark", content...)
}

func Menu(content ...view.View) view.ModificationApplyer {
	return Container("menu", content...)
}

func Menuitem(content ...view.View) view.ModificationApplyer {
	return Container("menuitem", content...)
}

func Meta() view.ModificationApplyer {
	return Inline("meta")
}

func Meter(content ...view.View) view.ModificationApplyer {
	return Container("meter", content...)
}

func Nav(content ...view.View) view.ModificationApplyer {
	return Container("nav", content...)
}

func Noscript(content ...view.View) view.ModificationApplyer {
	return Container("noscript", content...)
}

func Object(content ...view.View) view.ModificationApplyer {
	return Container("object", content...)
}

func Ol(content ...view.View) view.ModificationApplyer {
	return Container("ol", content...)
}

func Optgroup(content ...view.View) view.ModificationApplyer {
	return Container("optgroup", content...)
}

func Option(content ...view.View) view.ModificationApplyer {
	return Container("option", content...)
}

func Output(content ...view.View) view.ModificationApplyer {
	return Container("output", content...)
}

func P(content ...view.View) view.ModificationApplyer {
	return Container("p", content...)
}

func Param() view.ModificationApplyer {
	return Inline("param")
}

func Pre(content ...view.View) view.ModificationApplyer {
	return Container("pre", content...)
}

func Progress(content ...view.View) view.ModificationApplyer {
	return Container("progress", content...)
}

func Q(content ...view.View) view.ModificationApplyer {
	return Container("q", content...)
}

func Rp(content ...view.View) view.ModificationApplyer {
	return Container("rp", content...)
}

func Rt(content ...view.View) view.ModificationApplyer {
	return Container("rt", content...)
}

func Ruby(content ...view.View) view.ModificationApplyer {
	return Container("ruby", content...)
}

func S(content ...view.View) view.ModificationApplyer {
	return Container("s", content...)
}

func Samp(content ...view.View) view.ModificationApplyer {
	return Container("samp", content...)
}

func Script(content ...view.View) view.ModificationApplyer {
	return Container("script", content...)
}

func Section(content ...view.View) view.ModificationApplyer {
	return Container("section", content...)
}

func Select(content ...view.View) view.ModificationApplyer {
	return Container("select", content...)
}

func Small(content ...view.View) view.ModificationApplyer {
	return Container("small", content...)
}

func Source() view.ModificationApplyer {
	return Inline("source")
}

func Span(content ...view.View) view.ModificationApplyer {
	return Container("span", content...)
}

func Strong(content ...view.View) view.ModificationApplyer {
	return Container("strong", content...)
}

func Style(content ...view.View) view.ModificationApplyer {
	return Container("style", content...)
}

func Sub(content ...view.View) view.ModificationApplyer {
	return Container("sub", content...)
}

func Summary(content ...view.View) view.ModificationApplyer {
	return Container("summary", content...)
}

func Sup(content ...view.View) view.ModificationApplyer {
	return Container("sup", content...)
}

func Table(content ...view.View) view.ModificationApplyer {
	return Container("table", content...)
}

func Tbody(content ...view.View) view.ModificationApplyer {
	return Container("tbody", content...)
}

func Td(content ...view.View) view.ModificationApplyer {
	return Container("td", content...)
}

func Textarea(content ...view.View) view.ModificationApplyer {
	return Container("textarea", content...)
}

func Tfoot(content ...view.View) view.ModificationApplyer {
	return Container("tfoot", content...)
}

func Th(content ...view.View) view.ModificationApplyer {
	return Container("th", content...)
}

func Thead(content ...view.View) view.ModificationApplyer {
	return Container("thead", content...)
}

func Time(content ...view.View) view.ModificationApplyer {
	return Container("time", content...)
}

func Title(content ...view.View) view.ModificationApplyer {
	return Container("title", content...)
}

func Tr(content ...view.View) view.ModificationApplyer {
	return Container("tr", content...)
}

func Track() view.ModificationApplyer {
	return Inline("track")
}

func U(content ...view.View) view.ModificationApplyer {
	return Container("u", content...)
}

func Ul(content ...view.View) view.ModificationApplyer {
	return Container("ul", content...)
}

func Var(content ...view.View) view.ModificationApplyer {
	return Container("var", content...)
}

func Video(content ...view.View) view.ModificationApplyer {
	return Container("video", content...)
}

func Wbr() view.ModificationApplyer {
	return Inline("wbr")
}

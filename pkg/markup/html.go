package markup

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/marker/pkg/base"
)

func A(content ...view.View) base.ContainerModificator {
	return basicContainerTag("a", content...)
}

func Abbr(content ...view.View) base.ContainerModificator {
	return basicContainerTag("abbr", content...)
}

func Address(content ...view.View) base.ContainerModificator {
	return basicContainerTag("address", content...)
}

func Area(content ...view.View) base.ContainerModificator {
	return basicContainerTag("area", content...)
}

func Article(content ...view.View) base.ContainerModificator {
	return basicContainerTag("article", content...)
}

func Aside(content ...view.View) base.ContainerModificator {
	return basicContainerTag("aside", content...)
}

func Audio(content ...view.View) base.ContainerModificator {
	return basicContainerTag("audio", content...)
}

func B(content ...view.View) base.ContainerModificator {
	return basicContainerTag("b", content...)
}

func Base(content ...view.View) base.ContainerModificator {
	return basicContainerTag("base", content...)
}

func Bdi(content ...view.View) base.ContainerModificator {
	return basicContainerTag("bdi", content...)
}

func Bdo(content ...view.View) base.ContainerModificator {
	return basicContainerTag("bdo", content...)
}

func Blockquote(content ...view.View) base.ContainerModificator {
	return basicContainerTag("blockquote", content...)
}

func Body(content ...view.View) base.ContainerModificator {
	return basicContainerTag("body", content...)
}

func Br(content ...view.View) base.ContainerModificator {
	return basicContainerTag("br", content...)
}

func Button(content ...view.View) base.ContainerModificator {
	return basicContainerTag("button", content...)
}

func Canvas(content ...view.View) base.ContainerModificator {
	return basicContainerTag("canvas", content...)
}

func Caption(content ...view.View) base.ContainerModificator {
	return basicContainerTag("caption", content...)
}

func Cite(content ...view.View) base.ContainerModificator {
	return basicContainerTag("cite", content...)
}

func Code(content ...view.View) base.ContainerModificator {
	return basicContainerTag("code", content...)
}

func Col(content ...view.View) base.ContainerModificator {
	return basicContainerTag("col", content...)
}

func Colgroup(content ...view.View) base.ContainerModificator {
	return basicContainerTag("colgroup", content...)
}

func Data(content ...view.View) base.ContainerModificator {
	return basicContainerTag("data", content...)
}

func Datalist(content ...view.View) base.ContainerModificator {
	return basicContainerTag("datalist", content...)
}

func Dd(content ...view.View) base.ContainerModificator {
	return basicContainerTag("dd", content...)
}

func Del(content ...view.View) base.ContainerModificator {
	return basicContainerTag("del", content...)
}

func Details(content ...view.View) base.ContainerModificator {
	return basicContainerTag("details", content...)
}

func Dfn(content ...view.View) base.ContainerModificator {
	return basicContainerTag("dfn", content...)
}

func Dialog(content ...view.View) base.ContainerModificator {
	return basicContainerTag("dialog", content...)
}

func Div(content ...view.View) base.ContainerModificator {
	return basicContainerTag("div", content...)
}

func Dl(content ...view.View) base.ContainerModificator {
	return basicContainerTag("dl", content...)
}

func Dt(content ...view.View) base.ContainerModificator {
	return basicContainerTag("dt", content...)
}

func Em(content ...view.View) base.ContainerModificator {
	return basicContainerTag("em", content...)
}

func Embed(content ...view.View) base.ContainerModificator {
	return basicContainerTag("embed", content...)
}

func Fieldset(content ...view.View) base.ContainerModificator {
	return basicContainerTag("fieldset", content...)
}

func Figcaption(content ...view.View) base.ContainerModificator {
	return basicContainerTag("figcaption", content...)
}

func Figure(content ...view.View) base.ContainerModificator {
	return basicContainerTag("figure", content...)
}

func Footer(content ...view.View) base.ContainerModificator {
	return basicContainerTag("footer", content...)
}

func Form(content ...view.View) base.ContainerModificator {
	return basicContainerTag("form", content...)
}

func H1(content ...view.View) base.ContainerModificator {
	return basicContainerTag("h1", content...)
}

func H2(content ...view.View) base.ContainerModificator {
	return basicContainerTag("h2", content...)
}

func H3(content ...view.View) base.ContainerModificator {
	return basicContainerTag("h3", content...)
}

func H4(content ...view.View) base.ContainerModificator {
	return basicContainerTag("h4", content...)
}

func H5(content ...view.View) base.ContainerModificator {
	return basicContainerTag("h5", content...)
}

func H6(content ...view.View) base.ContainerModificator {
	return basicContainerTag("h6", content...)
}

func Head(content ...view.View) base.ContainerModificator {
	return basicContainerTag("head", content...)
}

func Header(content ...view.View) base.ContainerModificator {
	return basicContainerTag("header", content...)
}

func Hgroup(content ...view.View) base.ContainerModificator {
	return basicContainerTag("hgroup", content...)
}

func Hr(content ...view.View) base.ContainerModificator {
	return basicContainerTag("hr", content...)
}

func Html(content ...view.View) base.ContainerModificator {
	return basicContainerTag("html", content...)
}

func I(content ...view.View) base.ContainerModificator {
	return basicContainerTag("i", content...)
}

func Iframe(content ...view.View) base.ContainerModificator {
	return basicContainerTag("iframe", content...)
}

func Img(mod ...base.TagModificator) view.View {
	return base.Inline("img", mod...)
}

func Input(mod ...base.TagModificator) view.View {
	return base.Inline("input", mod...)
}

func Ins(content ...view.View) base.ContainerModificator {
	return basicContainerTag("ins", content...)
}

func Kbd(content ...view.View) base.ContainerModificator {
	return basicContainerTag("kbd", content...)
}

func Label(content ...view.View) base.ContainerModificator {
	return basicContainerTag("label", content...)
}

func Legend(content ...view.View) base.ContainerModificator {
	return basicContainerTag("legend", content...)
}

func Li(content ...view.View) base.ContainerModificator {
	return basicContainerTag("li", content...)
}

func Link(mod ...base.TagModificator) view.View {
	return base.Inline("link", mod...)
}

func Main(content ...view.View) base.ContainerModificator {
	return basicContainerTag("main", content...)
}

func Map(content ...view.View) base.ContainerModificator {
	return basicContainerTag("map", content...)
}

func Mark(content ...view.View) base.ContainerModificator {
	return basicContainerTag("mark", content...)
}

func Menu(content ...view.View) base.ContainerModificator {
	return basicContainerTag("menu", content...)
}

func Menuitem(content ...view.View) base.ContainerModificator {
	return basicContainerTag("menuitem", content...)
}

func Meta(mod ...base.TagModificator) view.View {
	return base.Inline("meta", mod...)
}

func Meter(content ...view.View) base.ContainerModificator {
	return basicContainerTag("meter", content...)
}

func Nav(content ...view.View) base.ContainerModificator {
	return basicContainerTag("nav", content...)
}

func Noscript(content ...view.View) base.ContainerModificator {
	return basicContainerTag("noscript", content...)
}

func Object(content ...view.View) base.ContainerModificator {
	return basicContainerTag("object", content...)
}

func Ol(content ...view.View) base.ContainerModificator {
	return basicContainerTag("ol", content...)
}

func Optgroup(content ...view.View) base.ContainerModificator {
	return basicContainerTag("optgroup", content...)
}

func Option(content ...view.View) base.ContainerModificator {
	return basicContainerTag("option", content...)
}

func Output(content ...view.View) base.ContainerModificator {
	return basicContainerTag("output", content...)
}

func P(content ...view.View) base.ContainerModificator {
	return basicContainerTag("p", content...)
}

func Param(mod ...base.TagModificator) view.View {
	return base.Inline("param", mod...)
}

func Pre(content ...view.View) base.ContainerModificator {
	return basicContainerTag("pre", content...)
}

func Progress(content ...view.View) base.ContainerModificator {
	return basicContainerTag("progress", content...)
}

func Q(content ...view.View) base.ContainerModificator {
	return basicContainerTag("q", content...)
}

func Rp(content ...view.View) base.ContainerModificator {
	return basicContainerTag("rp", content...)
}

func Rt(content ...view.View) base.ContainerModificator {
	return basicContainerTag("rt", content...)
}

func Ruby(content ...view.View) base.ContainerModificator {
	return basicContainerTag("ruby", content...)
}

func S(content ...view.View) base.ContainerModificator {
	return basicContainerTag("s", content...)
}

func Samp(content ...view.View) base.ContainerModificator {
	return basicContainerTag("samp", content...)
}

func Script(content ...view.View) base.ContainerModificator {
	return basicContainerTag("script", content...)
}

func Section(content ...view.View) base.ContainerModificator {
	return basicContainerTag("section", content...)
}

func Select(content ...view.View) base.ContainerModificator {
	return basicContainerTag("select", content...)
}

func Small(content ...view.View) base.ContainerModificator {
	return basicContainerTag("small", content...)
}

func Source(mod ...base.TagModificator) view.View {
	return base.Inline("source", mod...)
}

func Span(content ...view.View) base.ContainerModificator {
	return basicContainerTag("span", content...)
}

func Strong(content ...view.View) base.ContainerModificator {
	return basicContainerTag("strong", content...)
}

func Style(content ...view.View) base.ContainerModificator {
	return basicContainerTag("style", content...)
}

func Sub(content ...view.View) base.ContainerModificator {
	return basicContainerTag("sub", content...)
}

func Summary(content ...view.View) base.ContainerModificator {
	return basicContainerTag("summary", content...)
}

func Sup(content ...view.View) base.ContainerModificator {
	return basicContainerTag("sup", content...)
}

func Table(content ...view.View) base.ContainerModificator {
	return basicContainerTag("table", content...)
}

func Tbody(content ...view.View) base.ContainerModificator {
	return basicContainerTag("tbody", content...)
}

func Td(content ...view.View) base.ContainerModificator {
	return basicContainerTag("td", content...)
}

func Textarea(content ...view.View) base.ContainerModificator {
	return basicContainerTag("textarea", content...)
}

func Tfoot(content ...view.View) base.ContainerModificator {
	return basicContainerTag("tfoot", content...)
}

func Th(content ...view.View) base.ContainerModificator {
	return basicContainerTag("th", content...)
}

func Thead(content ...view.View) base.ContainerModificator {
	return basicContainerTag("thead", content...)
}

func Time(content ...view.View) base.ContainerModificator {
	return basicContainerTag("time", content...)
}

func Title(content ...view.View) base.ContainerModificator {
	return basicContainerTag("title", content...)
}

func Tr(content ...view.View) base.ContainerModificator {
	return basicContainerTag("tr", content...)
}

func Track(mod ...base.TagModificator) view.View {
	return base.Inline("track", mod...)
}

func U(content ...view.View) base.ContainerModificator {
	return basicContainerTag("u", content...)
}

func Ul(content ...view.View) base.ContainerModificator {
	return basicContainerTag("ul", content...)
}

func Var(content ...view.View) base.ContainerModificator {
	return basicContainerTag("var", content...)
}

func Video(content ...view.View) base.ContainerModificator {
	return basicContainerTag("video", content...)
}

func Wbr(mod ...base.TagModificator) view.View {
	return base.Inline("wbr", mod...)
}

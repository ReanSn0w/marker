package markup

import (
	"github.com/ReanSn0w/gew/v2/pkg/view"
	"github.com/ReanSn0w/marker/pkg/base"
)

func A(content ...view.View) view.ModificationApplyer {
	return base.Container("a")(content...)
}

func Abbr(content ...view.View) view.ModificationApplyer {
	return base.Container("abbr")(content...)
}

func Address(content ...view.View) view.ModificationApplyer {
	return base.Container("address")(content...)
}

func Area(content ...view.View) view.ModificationApplyer {
	return base.Container("area")(content...)
}

func Article(content ...view.View) view.ModificationApplyer {
	return base.Container("article")(content...)
}

func Aside(content ...view.View) view.ModificationApplyer {
	return base.Container("aside")(content...)
}

func Audio(content ...view.View) view.ModificationApplyer {
	return base.Container("audio")(content...)
}

func B(content ...view.View) view.ModificationApplyer {
	return base.Container("b")(content...)
}

func Base(content ...view.View) view.ModificationApplyer {
	return base.Container("base")(content...)
}

func Bdi(content ...view.View) view.ModificationApplyer {
	return base.Container("bdi")(content...)
}

func Bdo(content ...view.View) view.ModificationApplyer {
	return base.Container("bdo")(content...)
}

func Blockquote(content ...view.View) view.ModificationApplyer {
	return base.Container("blockquote")(content...)
}

func Body(content ...view.View) view.ModificationApplyer {
	return base.Container("body")(content...)
}

func Br(content ...view.View) view.ModificationApplyer {
	return base.Container("br")(content...)
}

func Button(content ...view.View) view.ModificationApplyer {
	return base.Container("button")(content...)
}

func Canvas(content ...view.View) view.ModificationApplyer {
	return base.Container("canvas")(content...)
}

func Caption(content ...view.View) view.ModificationApplyer {
	return base.Container("caption")(content...)
}

func Cite(content ...view.View) view.ModificationApplyer {
	return base.Container("cite")(content...)
}

func Code(content ...view.View) view.ModificationApplyer {
	return base.Container("code")(content...)
}

func Col(content ...view.View) view.ModificationApplyer {
	return base.Container("col")(content...)
}

func Colgroup(content ...view.View) view.ModificationApplyer {
	return base.Container("colgroup")(content...)
}

func Data(content ...view.View) view.ModificationApplyer {
	return base.Container("data")(content...)
}

func Datalist(content ...view.View) view.ModificationApplyer {
	return base.Container("datalist")(content...)
}

func Dd(content ...view.View) view.ModificationApplyer {
	return base.Container("dd")(content...)
}

func Del(content ...view.View) view.ModificationApplyer {
	return base.Container("del")(content...)
}

func Details(content ...view.View) view.ModificationApplyer {
	return base.Container("details")(content...)
}

func Dfn(content ...view.View) view.ModificationApplyer {
	return base.Container("dfn")(content...)
}

func Dialog(content ...view.View) view.ModificationApplyer {
	return base.Container("dialog")(content...)
}

func Div(content ...view.View) view.ModificationApplyer {
	return base.Container("div")(content...)
}

func Dl(content ...view.View) view.ModificationApplyer {
	return base.Container("dl")(content...)
}

func Dt(content ...view.View) view.ModificationApplyer {
	return base.Container("dt")(content...)
}

func Em(content ...view.View) view.ModificationApplyer {
	return base.Container("em")(content...)
}

func Embed(content ...view.View) view.ModificationApplyer {
	return base.Container("embed")(content...)
}

func Fieldset(content ...view.View) view.ModificationApplyer {
	return base.Container("fieldset")(content...)
}

func Figcaption(content ...view.View) view.ModificationApplyer {
	return base.Container("figcaption")(content...)
}

func Figure(content ...view.View) view.ModificationApplyer {
	return base.Container("figure")(content...)
}

func Footer(content ...view.View) view.ModificationApplyer {
	return base.Container("footer")(content...)
}

func Form(content ...view.View) view.ModificationApplyer {
	return base.Container("form")(content...)
}

func H1(content ...view.View) view.ModificationApplyer {
	return base.Container("h1")(content...)
}

func H2(content ...view.View) view.ModificationApplyer {
	return base.Container("h2")(content...)
}

func H3(content ...view.View) view.ModificationApplyer {
	return base.Container("h3")(content...)
}

func H4(content ...view.View) view.ModificationApplyer {
	return base.Container("h4")(content...)
}

func H5(content ...view.View) view.ModificationApplyer {
	return base.Container("h5")(content...)
}

func H6(content ...view.View) view.ModificationApplyer {
	return base.Container("h6")(content...)
}

func Head(content ...view.View) view.ModificationApplyer {
	return base.Container("head")(content...)
}

func Header(content ...view.View) view.ModificationApplyer {
	return base.Container("header")(content...)
}

func Hgroup(content ...view.View) view.ModificationApplyer {
	return base.Container("hgroup")(content...)
}

func Hr(content ...view.View) view.ModificationApplyer {
	return base.Container("hr")(content...)
}

func Html(content ...view.View) view.ModificationApplyer {
	return base.Container("html")(content...)
}

func I(content ...view.View) view.ModificationApplyer {
	return base.Container("i")(content...)
}

func Iframe(content ...view.View) view.ModificationApplyer {
	return base.Container("iframe")(content...)
}

func Img() view.ModificationApplyer {
	return base.Inline("img")
}

func Input() view.ModificationApplyer {
	return base.Inline("input")
}

func Ins(content ...view.View) view.ModificationApplyer {
	return base.Container("ins")(content...)
}

func Kbd(content ...view.View) view.ModificationApplyer {
	return base.Container("kbd")(content...)
}

func Label(content ...view.View) view.ModificationApplyer {
	return base.Container("label")(content...)
}

func Legend(content ...view.View) view.ModificationApplyer {
	return base.Container("legend")(content...)
}

func Li(content ...view.View) view.ModificationApplyer {
	return base.Container("li")(content...)
}

func Link() view.ModificationApplyer {
	return base.Inline("link")
}

func Main(content ...view.View) view.ModificationApplyer {
	return base.Container("main")(content...)
}

func Map(content ...view.View) view.ModificationApplyer {
	return base.Container("map")(content...)
}

func Mark(content ...view.View) view.ModificationApplyer {
	return base.Container("mark")(content...)
}

func Menu(content ...view.View) view.ModificationApplyer {
	return base.Container("menu")(content...)
}

func Menuitem(content ...view.View) view.ModificationApplyer {
	return base.Container("menuitem")(content...)
}

func Meta() view.ModificationApplyer {
	return base.Inline("meta")
}

func Meter(content ...view.View) view.ModificationApplyer {
	return base.Container("meter")(content...)
}

func Nav(content ...view.View) view.ModificationApplyer {
	return base.Container("nav")(content...)
}

func Noscript(content ...view.View) view.ModificationApplyer {
	return base.Container("noscript")(content...)
}

func Object(content ...view.View) view.ModificationApplyer {
	return base.Container("object")(content...)
}

func Ol(content ...view.View) view.ModificationApplyer {
	return base.Container("ol")(content...)
}

func Optgroup(content ...view.View) view.ModificationApplyer {
	return base.Container("optgroup")(content...)
}

func Option(content ...view.View) view.ModificationApplyer {
	return base.Container("option")(content...)
}

func Output(content ...view.View) view.ModificationApplyer {
	return base.Container("output")(content...)
}

func P(content ...view.View) view.ModificationApplyer {
	return base.Container("p")(content...)
}

func Param() view.ModificationApplyer {
	return base.Inline("param")
}

func Pre(content ...view.View) view.ModificationApplyer {
	return base.Container("pre")(content...)
}

func Progress(content ...view.View) view.ModificationApplyer {
	return base.Container("progress")(content...)
}

func Q(content ...view.View) view.ModificationApplyer {
	return base.Container("q")(content...)
}

func Rp(content ...view.View) view.ModificationApplyer {
	return base.Container("rp")(content...)
}

func Rt(content ...view.View) view.ModificationApplyer {
	return base.Container("rt")(content...)
}

func Ruby(content ...view.View) view.ModificationApplyer {
	return base.Container("ruby")(content...)
}

func S(content ...view.View) view.ModificationApplyer {
	return base.Container("s")(content...)
}

func Samp(content ...view.View) view.ModificationApplyer {
	return base.Container("samp")(content...)
}

func Script(content ...view.View) view.ModificationApplyer {
	return base.Container("script")(content...)
}

func Section(content ...view.View) view.ModificationApplyer {
	return base.Container("section")(content...)
}

func Select(content ...view.View) view.ModificationApplyer {
	return base.Container("select")(content...)
}

func Small(content ...view.View) view.ModificationApplyer {
	return base.Container("small")(content...)
}

func Source() view.ModificationApplyer {
	return base.Inline("source")
}

func Span(content ...view.View) view.ModificationApplyer {
	return base.Container("span")(content...)
}

func Strong(content ...view.View) view.ModificationApplyer {
	return base.Container("strong")(content...)
}

func Style(content ...view.View) view.ModificationApplyer {
	return base.Container("style")(content...)
}

func Sub(content ...view.View) view.ModificationApplyer {
	return base.Container("sub")(content...)
}

func Summary(content ...view.View) view.ModificationApplyer {
	return base.Container("summary")(content...)
}

func Sup(content ...view.View) view.ModificationApplyer {
	return base.Container("sup")(content...)
}

func Table(content ...view.View) view.ModificationApplyer {
	return base.Container("table")(content...)
}

func Tbody(content ...view.View) view.ModificationApplyer {
	return base.Container("tbody")(content...)
}

func Td(content ...view.View) view.ModificationApplyer {
	return base.Container("td")(content...)
}

func Textarea(content ...view.View) view.ModificationApplyer {
	return base.Container("textarea")(content...)
}

func Tfoot(content ...view.View) view.ModificationApplyer {
	return base.Container("tfoot")(content...)
}

func Th(content ...view.View) view.ModificationApplyer {
	return base.Container("th")(content...)
}

func Thead(content ...view.View) view.ModificationApplyer {
	return base.Container("thead")(content...)
}

func Time(content ...view.View) view.ModificationApplyer {
	return base.Container("time")(content...)
}

func Title(content ...view.View) view.ModificationApplyer {
	return base.Container("title")(content...)
}

func Tr(content ...view.View) view.ModificationApplyer {
	return base.Container("tr")(content...)
}

func Track() view.ModificationApplyer {
	return base.Inline("track")
}

func U(content ...view.View) view.ModificationApplyer {
	return base.Container("u")(content...)
}

func Ul(content ...view.View) view.ModificationApplyer {
	return base.Container("ul")(content...)
}

func Var(content ...view.View) view.ModificationApplyer {
	return base.Container("var")(content...)
}

func Video(content ...view.View) view.ModificationApplyer {
	return base.Container("video")(content...)
}

func Wbr() view.ModificationApplyer {
	return base.Inline("wbr")
}

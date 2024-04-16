package tag

import (
	"github.com/ReanSn0w/gew/v2/pkg/view"
	"github.com/ReanSn0w/marker/pkg/mk"
)

// Doctype generates a <!DOCTYPE html> declaration.
func Doctype() view.View {
	return mk.Text("<!DOCTYPE html>")
}

// Abbr generates an <abbr> tag, which represents
// an abbreviation or acronym.
func Abbr(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "abbr")(content...)
}

// Address generates an <address> tag, which provides
// contact information for a document or a major part of it.
func Address(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "address")(content...)
}

// Html generates an <html> tag, which is the root element
// of an HTML document.
func Html(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "html")(content...)
}

// Head generates a <head> tag, which contains machine-readable
// information (metadata) about the document, like its title,
// scripts, and style sheets.
func Head(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "head")(content...)
}

// Title generates a <title> tag, which defines the title
// of the document, shown in a browser's title bar or a page's tab.
func Title(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "title")(content...)
}

// Base specifies the base URL to use for all relative URLs in a document.
func Base() view.ModificationApplyer {
	return mk.NewTag(true, "base")()
}

// Link generates a <link> tag, which specifies relationships
// between the current document and an external resource.
func Link() view.ModificationApplyer {
	return mk.NewTag(true, "link")()
}

// Meta generates a <meta> tag, which represents metadata
// that cannot be represented by other HTML meta-related elements,
// like base, link, script, style or title.
func Meta() view.ModificationApplyer {
	return mk.NewTag(true, "meta")()
}

// Style generates a <style> tag, which contains style information
// for a document, or part of a document.
func Style(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "style")(content...)
}

// Script generates a <script> tag, which is used to embed
// or reference executable code.
func Script(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "script")(content...)
}

// Noscript generates a <noscript> tag, which defines an alternate
// content for users that have disabled scripts in their browser
// or have a browser that doesn't support script.
func Noscript(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "noscript")(content...)
}

// Body generates a <body> tag, which represents the content
// of an HTML document. There can be only one <body> element in a document.
func Body(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "body")(content...)
}

// Section generates a <section> tag, which represents a standalone
// section — which doesn't have a more specific semantic element
// to represent it — contained within an HTML document.
func Section(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "section")(content...)
}

// Nav generates a <nav> tag, which represents a section
// of a page whose purpose is to provide navigation links,
// either within the current document or to other documents.
func Nav(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "nav")(content...)
}

// Article generates an <article> tag, which represents a self-contained
// composition in a document, page, application, or site, which is intended
// to be independently distributable or reusable.
func Article(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "article")(content...)
}

// Aside generates an <aside> tag, which represents a portion of a document
// whose content is only indirectly related to the document's main content.
func Aside(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "aside")(content...)
}

// H1 generates a <h1> tag, which represents a first level heading.
func H1(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "h1")(content...)
}

// H2 generates a <h2> tag, which represents a second level heading.
func H2(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "h2")(content...)
}

// H3 generates a <h3> tag, which represents a third level heading.
func H3(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "h3")(content...)
}

// H4 generates a <h4> tag, which represents a fourth level heading.
func H4(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "h4")(content...)
}

// H5 generates a <h5> tag, which represents a fifth level heading.
func H5(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "h5")(content...)
}

// H6 generates a <h6> tag, which represents a sixth level heading.
func H6(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "h6")(content...)
}

// Header generates a <header> tag, which represents a container
// for introductory content or a set of navigational links.
func Header(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "header")(content...)
}

// Footer generates a <footer> tag, which represents a container
// for the footer of a document or a section.
func Footer(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "footer")(content...)
}

// Main generates a <main> tag, which represents the dominant
// content of the <body> of a document. The main content area consists
// of content that is directly related to or expands upon the central
// topic of a document, or the central functionality of an application.
func Main(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "main")(content...)
}

// P generates a <p> tag, which represents a paragraph.
func P(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "p")(content...)
}

// Hr generates a <hr> tag, which represents a thematic break
// between paragraph-level elements: for example, a change of scene
// in a story, or a shift of topic within a section.
func Hr() view.ModificationApplyer {
	return mk.NewTag(true, "hr")()
}

// Pre generates a <pre> tag, which represents preformatted text
// which is to be presented exactly as written in the HTML file.
func Pre(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "pre")(content...)
}

// Blockquote generates a <blockquote> tag, which indicates
// that the enclosed text is an extended quotation.
func Blockquote(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "blockquote")(content...)
}

// Ol generates an <ol> tag, which represents an ordered list
// of items — typically rendered as a numbered list.
func Ol(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "ol")(content...)
}

// Ul generates a <ul> tag, which represents an unordered list
// of items, typically rendered as a bulleted list.
func Ul(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "ul")(content...)
}

// Li generates a <li> tag, which is used to represent an item in a list.
func Li(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "li")(content...)
}

// Dl generates a <dl> tag, which represents a description list.
// The element encloses a list of groups of terms and descriptions.
func Dl(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "dl")(content...)
}

// Dt generates a <dt> tag, which specifies a term in a description
// or definition list, and as such must be used inside a <dl> element.
func Dt(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "dt")(content...)
}

// Dd generates a <dd> tag, which indicates the description
// of a term in a description list (<dl>).
func Dd(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "dd")(content...)
}

// Figure generates a <figure> tag, which represents self-contained
// content, potentially with an optional caption, which is specified
// using the (<figcaption>) element.
func Figure(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "figure")(content...)
}

// Figcaption generates a <figcaption> tag, which represents a caption
// or a legend associated with a figure or an illustration described by
// the rest of the data of the <figure> element the <figcaption> is included in.
func Figcaption(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "figcaption")(content...)
}

// Div generates a <div> tag, which is the generic container for flow
// content. It has no effect on the content or layout until styled using CSS.
func Div(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "div")(content...)
}

// Strong generates a <strong> tag, which indicates that its contents
// have strong importance, seriousness, or urgency.
func Strong(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "strong")(content...)
}

// Small generates a <small> tag, which represents side comments
// such as small print.
func Small(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "small")(content...)
}

// S generates a <s> tag, which renders text with a strikethrough,
// or a line through it.
func S(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "s")(content...)
}

// Cite generates a <cite> tag, which represents the title
// of a work (e.g. a book, a paper, an essay, a poem, a score, a song,
// a script, a film, a TV show, a game, a sculpture, a painting,
// a theatre production, a play, an opera, a musical, an exhibition,
// a legal case report, etc). This can be a work that is being quoted
// or referenced in detail (i.e. a citation), or it can just be a work
// that is mentioned in passing.
func Cite(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "cite")(content...)
}

// A generates a <a> tag, which creates a hyperlink to web pages,
// files, email addresses, locations in the same page, or anything
// else a URL can address.
func A(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "a")(content...)
}

// Em generates an <em> tag, which marks text that has stress emphasis.
func Em(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "em")(content...)
}

// Q generates a <q> tag, which indicates that the enclosed text
// is a short inline quotation.
func Q(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "q")(content...)
}

// Br generates a <br> tag, which produces a line break
// in text (carriage-return).
func Br() view.ModificationApplyer {
	return mk.NewTag(true, "br")()
}

// Wbr generates a <wbr> tag, which represents a word break
// opportunity—a position within text where the browser
// may optionally break a line, though its line-breaking rules
// would not otherwise create a break at that location.
func Wbr() view.ModificationApplyer {
	return mk.NewTag(true, "wbr")()
}

// Ins generates an <ins> tag, which represents a range of text
// that has been added to a document.
func Ins(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "ins")(content...)
}

// Del generates a <del> tag, which represents a range of text
// that has been deleted from a document.
func Del(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "del")(content...)
}

// Img generates an <img> tag, which embeds an image into the document.
func Img() view.ModificationApplyer {
	return mk.NewTag(true, "img")()
}

// Iframe generates an <iframe> tag, which represents a nested
// browsing context, embedding another HTML page into the current one.
func Iframe(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "iframe")(content...)
}

// Embed generates an <embed> tag, which provides an integration
// point for an external (typically non-HTML) application
// or interactive content.
func Embed(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "embed")(content...)
}

// Object generates an <object> tag, which represents an external
// resource, which can be treated as an image, a nested browsing
// context, or a resource to be handled by a plugin.
func Object(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "object")(content...)
}

// Param generates a <param> tag, which defines parameters
// for an <object> element.
func Param() view.ModificationApplyer {
	return mk.NewTag(true, "param")()
}

// Video generates a <video> tag, which embeds a media player
// which supports video playback into the document.
func Video(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "video")(content...)
}

// Audio generates an <audio> tag, which embeds a media player
// which supports audio playback into the document.
func Audio(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "audio")(content...)
}

// Picture generates a <picture> tag, which contains
// zero or more <source> elements and one <img> element
// to offer alternative versions of an image for different display/device scenarios.
func Picture(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "picture")(content...)
}

// Source generates a <source> tag, which is used to specify multiple
// media resources for media elements, such as <video> and <audio>.
func Source(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "source")(content...)
}

// Track generates a <track> tag, which is used as a child of the
// media elements—<audio> and <video>. It lets you specify timed text
// tracks (or time-based data), for example to automatically handle subtitles.
func Track() view.ModificationApplyer {
	return mk.NewTag(true, "track")()
}

// Canvas generates a <canvas> tag, which is used to draw graphics
// via JavaScript. It can, for instance, be used to draw graphs,
// make photo compositions, create animations, or even do real-time
// video processing or rendering.
func Canvas(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "canvas")(content...)
}

// Map generates a <map> tag, which contains a class of elements,
// known as <area>, that have varying geometric shapes,
// and that are associated with hyperlinks.
func Map(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "map")(content...)
}

// Area generates an <area> tag, which defines a hot-spot region
// on an image, and optionally associates it with a hypertext link.
// This element is used only within a <map> element.
func Area() view.ModificationApplyer {
	return mk.NewTag(true, "area")()
}

// Table generates a <table> tag, which represents tabular data — that is,
// information presented in a two-dimensional table comprised of rows
// and columns of cells containing data.
func Table(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "table")(content...)
}

// Caption generates a <caption> tag, which represents the title of a table.
// Though it is always the immediate first child of a <table>, its styling,
// using CSS, may place it elsewhere, relative to the table.
func Caption(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "caption")(content...)
}

// Colgroup generates a <colgroup> tag, which specifies a group
// of one or more columns in a table for formatting.
func Colgroup(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "colgroup")(content...)
}

// Col generates a <col> tag, which is used with the <colgroup> element
// to specify column properties for each column within a <colgroup> element.
func Col(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "col")(content...)
}

// Tbody generates a <tbody> tag, which is used to group the body content
// in an HTML table.
func Tbody(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "tbody")(content...)
}

// Thead generates a <thead> tag, which is used to group the header
// content in an HTML table.
func Thead(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "thead")(content...)
}

// Tfoot generates a <tfoot> tag, which is used to group the footer
// content in an HTML table.
func Tfoot(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "tfoot")(content...)
}

// Tr generates a <tr> tag, which is used to group the horizontal rows in an HTML table.
func Tr(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "tr")(content...)
}

// Td generates a <td> tag, which is used to define a cell of an HTML
// table that contains data.
func Td(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "td")(content...)
}

// Th generates a <th> tag, which is used to define a cell that contains
// header information in an HTML table.
func Th(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "th")(content...)
}

// Form generates a <form> tag, which represents a document section
// that contains interactive controls for submitting
// information to a web server.
func Form(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "form")(content...)
}

// Fieldset generates a <fieldset> tag, which is used to group several
// controls as well as labels (<label>) within a web form.
func Fieldset(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "fieldset")(content...)
}

// Legend generates a <legend> tag, which represents a caption
// for the content of its parent <fieldset>.
func Legend(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "legend")(content...)
}

// Label generates a <label> tag, which represents a caption
// for an item in a user interface.
func Label(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "label")(content...)
}

// Input generates an <input> tag, which is used to create interactive
// controls for web-based forms in order to accept data from the user.
func Input() view.ModificationApplyer {
	return mk.NewTag(true, "input")()
}

// Button generates a <button> tag, which represents a clickable button.
func Button(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "button")(content...)
}

// Select generates a <select> tag, which is used to create
// a drop-down list.
func Select(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "select")(content...)
}

// Datalist generates a <datalist> tag, which contains
// a set of <option> elements that represent the permissible
// or recommended options available to choose from within other controls.
func Datalist(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "datalist")(content...)
}

// Optgroup generates an <optgroup> tag, which creates a grouping
// of options within a <select> dropdown list.
func Optgroup(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "optgroup")(content...)
}

// Option generates an <option> tag, which is used to define
// the items in the dropdown list.
func Option(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "option")(content...)
}

// Textarea generates a <textarea> tag, which represents a multi-line
// plain-text editing control for a user to enter text.
func Textarea(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "textarea")(content...)
}

// Keygen generates a <keygen> tag, which represents a control
// for generating a public-private key pair and for submitting
// the public key from that key pair.
func Keygen() view.ModificationApplyer {
	return mk.NewTag(true, "keygen")()
}

// Output generates an <output> tag, which represents the result
// of a calculation or user action.
func Output(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "output")(content...)
}

// Progress generates a <progress> tag, which represents the progress
// of a task, such as a download or a setup.
func Progress(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "progress")(content...)
}

// Meter generates a <meter> tag, which represents a scalar measurement
// within a known range, or a fractional value.
func Meter(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "meter")(content...)
}

// Details generates a <details> tag, which is used as a disclosure
// widget from which the user can retrieve additional information.
func Details(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "details")(content...)
}

// Summary generates a <summary> tag, which is used as a heading
// for the <details> element. The <summary> element is used
// as a toggle to show or hide the additional information.
func Summary(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "summary")(content...)
}

// Menu generates a <menu> tag, which represents a group of commands
// that a user can perform or activate.
func Menu(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "menu")(content...)
}

// MenuItem generates a <menuitem> tag, which represents a command
// that a user can invoke from a popup menu.
func MenuItem(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "menuitem")(content...)
}

// Dialog generates a <dialog> tag, which represents a part
// of an application that a user interacts with to perform
// a single task, such as a dialog box, inspector, or window.
func Dialog(content ...view.View) view.ModificationApplyer {
	return mk.NewTag(false, "dialog")(content...)
}

package base

import (
	"context"
	"io"

	"github.com/ReanSn0w/gew/pkg/view"
)

const (
	TagImg    InlineTagName = "img"
	TagBr     InlineTagName = "br"
	TagHr     InlineTagName = "hr"
	TagInput  InlineTagName = "input"
	TagMeta   InlineTagName = "meta"
	TagLink   InlineTagName = "link"
	TagBase   InlineTagName = "base"
	TagArea   InlineTagName = "area"
	TagCol    InlineTagName = "col"
	TagParam  InlineTagName = "param"
	TagSource InlineTagName = "source"
	TagTrack  InlineTagName = "track"
	TagWbr    InlineTagName = "wbr"
)

type InlineTagName string

func Inline(tag InlineTagName, modificators ...TagModificator) view.View {
	var item view.View = &inlineTag{
		tag:        tag,
		attributes: Attributes{},
	}

	for _, m := range modificators {
		item = m(item)
	}

	return item
}

type inlineTag struct {
	tag        InlineTagName
	attributes Attributes
}

// MARK: - View

// Метод для удовлетворения интерфейсу view.View
func (i *inlineTag) Body(ctx context.Context) view.View {
	return view.External(i)
}

// MARK: - AttributeEditor

func (i *inlineTag) Get(key string) string {
	return i.attributes[key]
}

func (i *inlineTag) Set(key, value string) {
	i.attributes[key] = value
}

func (i *inlineTag) Append(key, value string) {
	i.attributes[key] += " " + value
}

func (i *inlineTag) Delete(key string) {
	delete(i.attributes, key)
}

// MARK: - Builder

func (i *inlineTag) Build(ctx context.Context, wr io.Writer) {
	wr.Write([]byte("<"))
	wr.Write([]byte(string(i.tag)))
	wr.Write([]byte(" "))
	wr.Write([]byte(i.attributes.build()))
	wr.Write([]byte(">"))
}

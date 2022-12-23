package marker

import (
	"context"
	"io"

	"github.com/ReanSn0w/gew/v2/pkg/view"
)

func Inline(tag string) view.ModificationApplyer {
	return view.NewModificationApplyer(&inlineTag{
		tag:        tag,
		attributes: Attributes{},
	})
}

type inlineTag struct {
	tag        string
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

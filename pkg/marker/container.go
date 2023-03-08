package marker

import (
	"context"
	"io"

	"github.com/ReanSn0w/gew/v2/pkg/view"
)

func Container(tag string, content ...view.View) view.ModificationApplyer {
	return view.NewModificationApplyer(&containerTag{
		tag:        tag,
		attributes: Attributes{},
		content:    view.Group(content...),
	})
}

type containerTag struct {
	tag        string
	attributes Attributes
	content    view.View
}

// MARK: - View

// Метод для удовлетворения интерфейсу view.View
func (i *containerTag) Body(ctx context.Context) view.View {
	return view.External(i)
}

// MARK: - AttributeEditor

func (i *containerTag) Get(key string) string {
	return i.attributes[key]
}

func (i *containerTag) Set(key, value string) {
	i.attributes[key] = value
}

func (i *containerTag) Append(key, value string) {
	i.attributes[key] += " " + value
}

func (i *containerTag) Delete(key string) {
	delete(i.attributes, key)
}

// MARK: - Builder

func (i *containerTag) Build(ctx context.Context, wr io.Writer) {
	wr.Write([]byte("<"))
	wr.Write([]byte(string(i.tag)))
	if val := i.attributes.build(); val != "" {
		wr.Write([]byte(" "))
		wr.Write([]byte(val))
	}
	wr.Write([]byte(">"))

	if i.content != nil {
		wr.Write([]byte("\n"))
		NewBuilder().Build(ctx, wr, i.content)
		wr.Write([]byte("\n"))
	}

	wr.Write([]byte("</"))
	wr.Write([]byte(string(i.tag)))
	wr.Write([]byte(">"))
}

// MARK: - Attribute Modificator

func PlaceInContainer(tag string) view.Modificator {
	return func(v view.View) view.View {
		return Container(tag, v)()
	}
}

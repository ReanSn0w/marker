package base

import (
	"context"
	"io"

	"github.com/ReanSn0w/gew/v2/pkg/view"
)

type ContainerTagName string

func Container(tag ContainerTagName) view.Builder {
	return func(content ...view.View) view.ModificationApplyer {
		return view.NewModificationApplyer(&containerTag{
			tag:        tag,
			attributes: Attributes{},
			content: func() view.View {
				group := []view.View{}

				for _, c := range content {
					v, ok := c.(view.View)

					if !ok {
						return view.External(content)
					}

					group = append(group, v)
				}

				return view.Group(group...)
			}(),
		})
	}
}

type containerTag struct {
	tag        ContainerTagName
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
	wr.Write([]byte(">\n"))

	GetBuilderFromContext(ctx).Build(ctx, wr, i.content)

	wr.Write([]byte("\n</"))
	wr.Write([]byte(string(i.tag)))
	wr.Write([]byte(">"))
}

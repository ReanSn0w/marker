package mk

import (
	"context"
	"fmt"
	"io"

	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/primitive"
)

func Build(ctx context.Context, wr io.Writer, element view.View) error {
	ctx = primitive.Set(ctx)
	build(ctx, wr, element)
	return nil
}

func build(ctx context.Context, wr io.Writer, element view.View) {
	view.Builder(ctx, element, func(ctx context.Context, i interface{}) {
		switch el := i.(type) {
		case *Page:
			el.build(ctx, wr)
		case *Tag:
			el.build(ctx, wr)
		case *style:
			el.write(ctx, wr)
		case Text:
			wr.Write([]byte(el))
		}
	})
}

// MARK: Tag

func NewTag(inline bool, name string) func(content ...view.View) view.Use {
	return func(content ...view.View) view.Use {
		return view.NewView(&Tag{
			inline:  inline,
			name:    name,
			content: view.Group(content...),
		})
	}
}

type Tag struct {
	inline  bool
	name    string
	content view.View
}

func (t *Tag) Body(ctx context.Context) view.View {
	return view.External(t)
}

func (t *Tag) build(ctx context.Context, wr io.Writer) {
	attr := primitive.Get(ctx).ExtractData()

	wr.Write([]byte(fmt.Sprintf("<%s", t.name)))

	for k, v := range attr {
		if v == "" {
			wr.Write([]byte(" " + string(k)))
			continue
		}

		wr.Write([]byte(" " + string(k) + "=" + "\"" + string(v) + "\""))
	}

	wr.Write([]byte(">"))

	if !t.inline {
		build(ctx, wr, t.content)
		wr.Write([]byte("</" + t.name + ">"))
	}
}

// MARK: Text

type Text string

// метод для удовлетворения интерфейсу View
func (t Text) Body(ctx context.Context) view.View {
	return view.External(t)
}

// MARK: Style

func Style() view.View {
	return &style{}
}

type style struct{}

func (s *style) Body(ctx context.Context) view.View {
	return nil
}

func (s *style) write(ctx context.Context, wr io.Writer) {
	globalData := primitive.Get(ctx)
	wr.Write([]byte("<style>"))
	globalData.WriteStyle(wr)
	wr.Write([]byte("</style>"))
}

package page

import (
	"context"
	"io"

	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/marker/pkg/base"
	"github.com/ReanSn0w/marker/pkg/markup"
)

type PageInterface interface {
	Write(ctx context.Context, w io.Writer)
}

func Page(headers ...view.View) func(content ...view.View) PageInterface {
	return func(content ...view.View) PageInterface {
		return &page{
			headers: headers,
			content: content,
			builder: base.NewBuilder(),
		}
	}
}

type page struct {
	builder base.Builder
	headers []view.View
	content []view.View
}

func (p *page) Body(ctx context.Context) view.View {
	return view.Group(
		markup.Doctype(),
		markup.Html(
			markup.Head(
				p.headers...,
			),
			markup.Body(
				p.content...,
			),
		),
	)
}

func (p *page) Write(ctx context.Context, w io.Writer) {
	p.builder.Build(ctx, w, p)
}

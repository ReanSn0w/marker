package marker

import (
	"context"
	"io"

	"github.com/ReanSn0w/gew/v2/pkg/view"
)

type PageInterface interface {
	Write(ctx context.Context, w io.Writer)
}

func Page(headers ...view.View) func(content ...view.View) PageInterface {
	return func(content ...view.View) PageInterface {
		return &page{
			headers: headers,
			content: content,
			builder: NewBuilder(nil),
		}
	}
}

type page struct {
	builder Builder
	headers []view.View
	content []view.View
}

func (p *page) Body(ctx context.Context) view.View {
	return view.Group(
		Doctype(),
		Html(
			Head(
				p.headers...,
			),
			Body(
				p.content...,
			),
		),
	)
}

func (p *page) Write(ctx context.Context, w io.Writer) {
	p.builder.Build(ctx, w, p)
}

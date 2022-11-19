package marker

import (
	"context"
	"io"

	"github.com/ReanSn0w/gew/pkg/view"
)

func NewPage(head, body view.View) *Marker {
	return &Marker{
		Content: view.Group(
			view.External("<!DOCTYPE html>"),
			Html(nil, view.Group(
				Head(nil, head),
				Body(nil, body),
			)),
		),
	}
}

func NewMarker(content view.View) *Marker {
	return &Marker{
		Content: content,
	}
}

type Marker struct {
	Content view.View
}

func (p *Marker) WritePage(ctx context.Context, wr io.Writer) {
	builder := &builder{}
	ctx = withBuilder(ctx, builder)
	builder.Build(ctx, wr, p.Content)
}

package marker

import (
	"context"
	"io"

	"github.com/ReanSn0w/gew/pkg/view"
)

func NewMarker(content view.View) *Marker {
	return &Marker{
		Content: content,
	}
}

type Marker struct {
	Content view.View
}

func (p *Marker) Write(ctx context.Context, wr io.Writer) {
	builder := &builder{}
	ctx = withBuilder(ctx, builder)
	builder.Build(ctx, wr, p.Content)
}

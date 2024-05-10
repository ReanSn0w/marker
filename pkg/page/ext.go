package page

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/common"
	"github.com/ReanSn0w/marker/pkg/mk"
)

func (p *Page) Write(ctx context.Context, code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "text/html")
	mk.Build(ctx, w, p)
}

func (p *Page) Body(ctx context.Context) view.View {
	return view.External(p)
}

func (p *Page) Build(ctx context.Context, wr io.Writer) {
	body := p.buildPart(ctx, view.NewView(view.Group(p.body...)(p.mod...))(
		view.ContextModificator(func(ctx context.Context) context.Context {
			common.Get(ctx).EnableStyle()
			return ctx
		}),
	))

	head := p.buildPart(ctx, view.Group(p.head...)(
		view.ContextModificator(func(ctx context.Context) context.Context {
			common.Get(ctx).DisableStyle()
			return ctx
		}),
	))

	style := p.buildPart(ctx, mk.Style())

	wr.Write([]byte("<!doctype html>"))
	wr.Write([]byte("<html lang=\"" + p.lang + "\">"))
	wr.Write([]byte("<head>"))
	wr.Write(head)
	wr.Write(style)
	wr.Write([]byte("</head>"))
	wr.Write([]byte("<body>"))
	wr.Write(body)
	wr.Write([]byte("</body>"))
	wr.Write([]byte("</html>"))
}

func (p *Page) buildPart(ctx context.Context, item view.View) []byte {
	buf := new(bytes.Buffer)
	mk.Build(ctx, buf, item)
	return buf.Bytes()
}

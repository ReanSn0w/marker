package mk

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/primitive"
)

func NewPage(head ...view.View) func(body ...view.View) *Page {
	return func(body ...view.View) *Page {
		return &Page{
			Lang: "ru",
			body: view.Group(body...)(
				view.ContextModificator(func(ctx context.Context) context.Context {
					primitive.Get(ctx).DisableStyle()
					return ctx
				}),
			),
			head: view.Group(head...)(
				view.ContextModificator(func(ctx context.Context) context.Context {
					primitive.Get(ctx).EnableStyle()
					return ctx
				}),
			),
		}
	}
}

type Page struct {
	Lang string

	head view.View
	body view.View
}

func (p *Page) Write(ctx context.Context, code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "text/html")
	Build(ctx, w, p)
}

func (p *Page) Body(ctx context.Context) view.View {
	return view.External(p)
}

func (p *Page) build(ctx context.Context, wr io.Writer) {
	body := p.buildPart(ctx, p.body)
	head := p.buildPart(ctx, p.head)
	style := p.buildPart(ctx, Style())

	wr.Write([]byte("<!doctype html>"))
	wr.Write([]byte("<html lang=\"" + p.Lang + "\">"))
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
	build(ctx, buf, item)
	return buf.Bytes()
}

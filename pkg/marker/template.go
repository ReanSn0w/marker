package marker

import (
	"bytes"
	"context"
	"html/template"

	"github.com/ReanSn0w/gew/v2/pkg/view"
)

func NewTemplate() *Template {
	return &Template{
		tmpl: template.New("render"),
	}
}

type Template struct {
	tmpl *template.Template
}

func (r *Template) Prepare(prepare func(*template.Template) *template.Template) {
	r.tmpl = prepare(r.tmpl)
}

func (r *Template) Make(name string, data any) view.View {
	return view.Closure(func(ctx context.Context) view.View {
		buffer := new(bytes.Buffer)

		err := r.tmpl.ExecuteTemplate(buffer, name, data)
		if err != nil {
			return Error(err)
		}

		return Text(buffer.String())
	})
}

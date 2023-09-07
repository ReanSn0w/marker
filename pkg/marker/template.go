package marker

import (
	"bytes"
	"context"
	"html/template"

	"github.com/ReanSn0w/gew/v2/pkg/view"
)

func NewTemplate(prepare func(*template.Template) *template.Template) *Template {
	return &Template{
		tmpl: prepare(template.New("template")),
	}
}

type Template struct {
	tmpl *template.Template
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

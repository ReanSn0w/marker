package marker

import (
	"context"
	"fmt"
	"io"

	"github.com/ReanSn0w/gew/pkg/view"
)

func Text(value string) view.View {
	return view.External(raw(value))
}

func Snippet(value string, objects ...interface{}) view.View {
	return view.External(raw(fmt.Sprintf(value, objects...)))
}

type raw string

func (r raw) Build(ctx context.Context, wr io.Writer) {
	wr.Write([]byte(fmt.Sprintf("\n%s", r)))
}

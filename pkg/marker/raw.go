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

type raw string

func (r raw) Build(ctx context.Context, wr io.Writer) {
	wr.Write([]byte(fmt.Sprintf("\n%s", r)))
}

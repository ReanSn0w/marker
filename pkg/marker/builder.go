package marker

import (
	"context"
	"fmt"
	"io"

	"github.com/ReanSn0w/gew/v2/pkg/view"
)

// var builderCtxKey = &builderContextKey{}

type Builder interface {
	Build(ctx context.Context, wr io.Writer, item view.View)
}

type Logger interface {
	Logf(format string, args ...interface{})
}

func NewBuilder(logger Logger) Builder {
	return &builder{logger: logger}
}

type builder struct {
	logger Logger
}

type builderItem interface {
	Build(ctx context.Context, wr io.Writer)
}

func (b *builder) Build(ctx context.Context, wr io.Writer, item view.View) {
	view.Build(
		//view.Context(builderCtxKey, b)(item),
		item,
		ctx,
		func(i interface{}, ctx context.Context) {
			if t, ok := i.(builderItem); ok && t != nil {
				t.Build(ctx, wr)
				return
			}

			switch t := i.(type) {
			case int:
				wr.Write([]byte(fmt.Sprint(t)))
			case string:
				wr.Write([]byte(t))
			case error:
				b.log("WARN: builder: error detected. Err: %s", t.Error())
			default:
				b.log("WARN: unknown type %T, value will be rendered", t)
				wr.Write([]byte(fmt.Sprintf("%v", t)))
			}
		},
	)
}

func (b *builder) log(format string, args ...interface{}) {
	if b.logger != nil {
		b.logger.Logf(format, args...)
	}
}

// type builderContextKey struct{}

// func GetBuilderFromContext(ctx context.Context) Builder {
// 	return ctx.Value(builderCtxKey).(Builder)
// }

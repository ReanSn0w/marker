package marker

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/ReanSn0w/gew/v2/pkg/view"
)

// var builderCtxKey = &builderContextKey{}

type Builder interface {
	Build(ctx context.Context, wr io.Writer, item view.View)
}

func NewBuilder() Builder {
	return &builder{}
}

type builder struct{}

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
			default:
				log.Println(t)
				wr.Write([]byte(fmt.Sprintf("%v", t)))
			}
		},
	)
}

// type builderContextKey struct{}

// func GetBuilderFromContext(ctx context.Context) Builder {
// 	return ctx.Value(builderCtxKey).(Builder)
// }

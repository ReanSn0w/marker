package marker

import (
	"context"
	"fmt"
	"io"

	"github.com/ReanSn0w/gew/pkg/view"
)

var (
	builderCtxKey = &builderCtx{}
)

type builderCtx struct{}

type Builder interface {
	Build(ctx context.Context, wr io.Writer, v view.View)
}

type builder struct{}

func (b *builder) Build(ctx context.Context, wr io.Writer, v view.View) {
	view.Build(v, ctx, func(i interface{}, ctx context.Context) {
		switch item := i.(type) {
		case *Node:
			item.Build(ctx, wr)
		default:
			wr.Write([]byte(fmt.Sprintf("%v", item)))
		}
	})
}

func GetBuilder(ctx context.Context) Builder {
	b, ok := ctx.Value(builderCtxKey).(Builder)
	if !ok {
		panic("builder not found")
	}

	return b
}

func withBuilder(ctx context.Context, b Builder) context.Context {
	return context.WithValue(ctx, builderCtxKey, b)
}

package ui

import (
	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/mk"
	"github.com/ReanSn0w/marker/pkg/tag"
)

func Text(text string) view.Use {
	return view.NewView(tag.P(mk.Text(text))(
		mk.WithStyle().
			Value("all", "unset").
			Value("font-size", "16px").
			Extract(),
	))
}

func Button(content ...view.View) view.Use {
	return tag.Button(content...)
}

func Block(content ...view.View) view.Use {
	return view.NewView(
		tag.Div(content...)(
			mk.Class("ui-block"),
			mk.WithStyle().
				Selector(":not(:last-child)").
				Value("margin-bottom", "1.5rem").
				Extract(),
		),
	)
}

func Box(content ...view.View) view.Use {
	return view.NewView(
		tag.Div(content...)(
			mk.WithStyle().
				Value("display", "block").
				Value("box-shadow", "0 4px 8px 0 rgba(0, 0, 0, 0.2)").
				Value("border-top-left-radius", "8px").
				Value("border-top-right-radius", "8px").
				Value("border-bottom-left-radius", "8px").
				Value("border-bottom-right-radius", "8px").
				Value("padding-top", "16px").
				Value("padding-bottom", "16px").
				Value("padding-left", "16px").
				Value("padding-right", "16px").
				Extract(),
		),
	)
}

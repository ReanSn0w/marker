package main

import (
	"net/http"

	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/mk"
	"github.com/ReanSn0w/marker/pkg/page"
	"github.com/ReanSn0w/marker/pkg/tag"
	"github.com/ReanSn0w/marker/pkg/ui"
)

func main() {
	previewElement(
		tag.Div(
			view.For(20, func(i int) view.View {
				return ui.Box(
					tag.Img()(
						mk.Set("src", "https://picsum.photos/300"),
						mk.Set("alt", "image"),
						ui.PlaceInBlock(),
						mk.WithStyle().
							Value("max-width", "100%").
							Value("height", "auto").
							Value("object-fit", "cover").
							Extract(),
					),
					ui.Text("Hello, world!"),
				)
			}),
		)(
			ui.Grid().
				Columns(ui.Fr(1), ui.Fr(1), ui.Fr(1)).
				VGap(ui.Px(20)).
				HGap(ui.Px(20)).
				Extract(),
		),
	)
}

func previewElement(item ...view.View) {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page.New("Preview page").
			Content(item...).
			Mod(
				mk.Class("test"),
				mk.WithStyle().
					Value("background-color", "white").
					Value("color", "black").
					Extract(),
				mk.WithStyle().
					Rule("media (prefers-color-scheme: dark)").
					Value("background-color", "black").
					Value("color", "white").
					Extract()
			),
			Write(r.Context(), http.StatusOK, w)
	}))
}

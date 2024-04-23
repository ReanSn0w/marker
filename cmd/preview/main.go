package main

import (
	"net/http"

	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/mk"
	"github.com/ReanSn0w/marker/pkg/tag"
	"github.com/ReanSn0w/marker/pkg/ui"
)

func main() {
	previewElement(
		ui.Grid(ui.NewGridPreference("20px", "20px", "1fr", "2fr", "3fr"))(
			tag.Div(mk.Text("Blue"))(
				ui.BackgroundColor("blue"),
				ui.Height("30px"), mk.Class("blue"),
			),
			tag.Div(mk.Text("Green"))(
				ui.BackgroundColor("green"),
				ui.Height("60px"), mk.Class("green"),
			),
			tag.Div(mk.Text("Yellow"))(
				ui.BackgroundColor("yellow"),
				ui.Height("90px"),
				mk.Class("yellow"),
			),
			mk.Text("red"),
		)(
			ui.BackgroundColor("red"),
			mk.Class("grid"),
		),
		tag.Hr(),
		ui.Grid(ui.NewGridPreference("20px", "20px", "1fr", "2fr", "3fr"))(
			tag.Div(mk.Text("Blue"))(
				ui.BackgroundColor("blue"),
				ui.Height("30px"), mk.Class("blue"),
			),
			tag.Div(mk.Text("Green"))(
				ui.BackgroundColor("green"),
				ui.Height("60px"), mk.Class("green"),
			),
			tag.Div(mk.Text("Yellow"))(
				ui.BackgroundColor("yellow"),
				ui.Height("90px"), mk.Class("yellow"),
			),
			mk.Text("red"),
		)(
			ui.BackgroundColor("red"),
			mk.Class("grid"),
		),
		tag.Hr(),
		ui.Grid(ui.NewGridPreference("20px", "20px", "1fr", "2fr", "3fr"))(
			tag.Div(mk.Text("Blue"))(
				ui.BackgroundColor("blue"),
				ui.Height("30px"), mk.Class("blue"),
			),
			tag.Div(mk.Text("Green"))(
				ui.BackgroundColor("green"),
				ui.Height("60px"), mk.Class("green"),
			),
			tag.Div(mk.Text("Yellow"))(
				ui.BackgroundColor("yellow"),
				ui.Height("90px"), mk.Class("yellow"),
			),
			mk.Text("red"),
		)(
			ui.BackgroundColor("red"),
			mk.Class("grid"),
		),
		tag.Hr(),
		ui.Grid(ui.NewGridPreference("20px", "20px", "1fr", "2fr", "3fr"))(
			tag.Div(mk.Text("Blue"))(
				ui.BackgroundColor("blue"),
				ui.Height("30px"), mk.Class("blue"),
			),
			tag.Div(mk.Text("Green"))(
				ui.BackgroundColor("green"),
				ui.Height("60px"), mk.Class("green"),
			),
			tag.Div(mk.Text("Yellow"))(
				ui.BackgroundColor("yellow"),
				ui.Height("90px"), mk.Class("yellow"),
			),
			mk.Text("red"),
		)(
			ui.BackgroundColor("red"),
			mk.Class("grid"),
		),
		tag.Hr(),
		ui.Grid(ui.NewGridPreference("20px", "20px", "1fr", "2fr", "3fr"))(
			tag.Div(mk.Text("Blue"))(
				ui.BackgroundColor("blue"),
				ui.Height("30px"), mk.Class("blue"),
			),
			tag.Div(mk.Text("Green"))(
				ui.BackgroundColor("green"),
				ui.Height("60px"), mk.Class("green"),
			),
			tag.Div(mk.Text("Yellow"))(
				ui.BackgroundColor("yellow"),
				ui.Height("90px"), mk.Class("yellow"),
			),
			mk.Text("red"),
		)(
			ui.BackgroundColor("red"),
			mk.Class("grid"),
		),
		tag.Hr(),
		ui.Grid(ui.NewGridPreference("20px", "20px", "1fr", "2fr", "3fr"))(
			tag.Div(mk.Text("Blue"))(
				ui.BackgroundColor("blue"),
				ui.Height("30px"), mk.Class("blue"),
			),
			tag.Div(mk.Text("Green"))(
				ui.BackgroundColor("green"),
				ui.Height("60px"), mk.Class("green"),
			),
			tag.Div(mk.Text("Yellow"))(
				ui.BackgroundColor("yellow"),
				ui.Height("90px"), mk.Class("yellow"),
			),
			mk.Text("red"),
		)(
			ui.BackgroundColor("red"),
			mk.Class("grid"),
		),
		tag.Hr(),
		ui.Grid(ui.NewGridPreference("20px", "20px", "1fr", "2fr", "3fr"))(
			tag.Div(mk.Text("Blue"))(
				ui.BackgroundColor("blue"),
				ui.Height("30px"), mk.Class("blue"),
			),
			tag.Div(mk.Text("Green"))(
				ui.BackgroundColor("green"),
				ui.Height("60px"), mk.Class("green"),
			),
			tag.Div(mk.Text("Yellow"))(
				ui.BackgroundColor("yellow"),
				ui.Height("90px"), mk.Class("yellow"),
			),
			mk.Text("red"),
		)(
			ui.BackgroundColor("red"),
			mk.Class("grid"),
		),
		tag.Hr(),
		ui.Grid(ui.NewGridPreference("20px", "20px", "1fr", "2fr", "3fr"))(
			tag.Div(mk.Text("Blue"))(
				ui.BackgroundColor("blue"),
				ui.Height("30px"), mk.Class("blue"),
			),
			tag.Div(mk.Text("Green"))(
				ui.BackgroundColor("green"),
				ui.Height("60px"), mk.Class("green"),
			),
			tag.Div(mk.Text("Yellow"))(
				ui.BackgroundColor("yellow"),
				ui.Height("90px"), mk.Class("yellow"),
			),
			mk.Text("red"),
		)(
			ui.BackgroundColor("red"),
			mk.Class("grid"),
		),
		tag.Hr(),
		ui.Grid(ui.NewGridPreference("20px", "20px", "1fr", "2fr", "3fr"))(
			tag.Div(mk.Text("Blue"))(
				ui.BackgroundColor("blue"),
				ui.Height("30px"), mk.Class("blue"),
			),
			tag.Div(mk.Text("Green"))(
				ui.BackgroundColor("green"),
				ui.Height("60px"), mk.Class("green"),
			),
			tag.Div(mk.Text("Yellow"))(
				ui.BackgroundColor("yellow"),
				ui.Height("90px"), mk.Class("yellow"),
			),
			mk.Text("red"),
		)(
			ui.BackgroundColor("red"),
			mk.Class("grid"),
		),
		tag.Hr(),
		ui.Grid(ui.NewGridPreference("20px", "20px", "1fr", "2fr", "3fr"))(
			tag.Div(mk.Text("Blue"))(
				ui.BackgroundColor("blue"),
				ui.Height("30px"), mk.Class("blue"),
			),
			tag.Div(mk.Text("Green"))(
				ui.BackgroundColor("green"),
				ui.Height("60px"), mk.Class("green"),
			),
			tag.Div(mk.Text("Yellow"))(
				ui.BackgroundColor("yellow"),
				ui.Height("90px"), mk.Class("yellow"),
			),
			mk.Text("red"),
		)(
			ui.BackgroundColor("red"),
			mk.Class("grid"),
		),
		tag.Hr(),
		ui.Grid(ui.NewGridPreference("20px", "20px", "1fr", "2fr", "3fr"))(
			tag.Div(mk.Text("Blue"))(
				ui.BackgroundColor("blue"),
				ui.Height("30px"), mk.Class("blue"),
			),
			tag.Div(mk.Text("Green"))(
				ui.BackgroundColor("green"),
				ui.Height("60px"), mk.Class("green"),
			),
			tag.Div(mk.Text("Yellow"))(
				ui.BackgroundColor("yellow"),
				ui.Height("90px"), mk.Class("yellow"),
			),
			mk.Text("red"),
		)(
			ui.BackgroundColor("red"),
			mk.Class("grid"),
		),
	)
}

func previewElement(item ...view.View) {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mk.NewPage(mk.NewTag(false, "title")(mk.Text("Preview Page")))(item...).
			Write(r.Context(), http.StatusOK, w)
	}))
}

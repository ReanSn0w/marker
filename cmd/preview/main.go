package main

import (
	"net/http"

	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/mk"
	"github.com/ReanSn0w/marker/pkg/tag"
)

func main() {
	previewElement(
		mk.Text("Hello World!"),
		tag.Div(
			tag.Div(
				mk.Text("Hello"),
			)(
				mk.WithStyle("", mk.Declarations{
					"background-color": "blue",
					"height":           "30px",
				}),
			),
			tag.Div(
				mk.Text("Hello"),
			)(
				mk.WithStyle("", mk.Declarations{
					"background-color": "green",
					"height":           "30px",
				}),
			),
		)(
			mk.WithStyle("", mk.Declarations{
				"background-color": "red",
				"height":           "1000px",
			}),
		),
	)
}

func previewElement(item ...view.View) {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mk.NewPage(mk.NewTag(false, "title")(mk.Text("Preview Page")))(item...).
			Write(r.Context(), http.StatusOK, w)
	}))
}

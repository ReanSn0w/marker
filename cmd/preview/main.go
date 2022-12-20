package main

import (
	"net/http"

	"github.com/ReanSn0w/gew/v2/pkg/view"
	"github.com/ReanSn0w/marker/pkg/markup"
	"github.com/ReanSn0w/marker/pkg/page"
	"github.com/go-chi/chi"
)

func main() {
	previewElement(
		markup.Text("Hello World!"),
		markup.Div(
			markup.Text("Hi, how are you?"),
		),
	)
}

func previewElement(item ...view.View) {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		page.Page(
			markup.Title(markup.Text("Element Preview")),
		)(
			item...,
		).Write(ctx, w)
	})

	http.ListenAndServe(":8080", r)
}

package main

import (
	"net/http"

	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/marker/pkg/base"
	"github.com/ReanSn0w/marker/pkg/markup"
	"github.com/ReanSn0w/marker/pkg/page"
	"github.com/go-chi/chi"
)

func main() {
	previewElement(
		markup.Text("Hello World!"),
		markup.Text("Hi, how are you?")(
			base.PlaceInContainer("div"),
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

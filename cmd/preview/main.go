package main

import (
	"net/http"

	"github.com/ReanSn0w/gew/v2/pkg/view"
	"github.com/ReanSn0w/marker/pkg/marker"
	"github.com/go-chi/chi"
)

func main() {
	previewElement(
		marker.Text("Hello World!"),
		marker.Div(
			marker.Text("Hi, how are you?"),
		),
	)
}

func previewElement(item ...view.View) {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		marker.Page(
			marker.Title(marker.Text("Element Preview")),
		)(
			item...,
		).Write(ctx, w)
	})

	http.ListenAndServe(":8080", r)
}

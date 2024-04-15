package main

import (
	"net/http"

	"github.com/ReanSn0w/gew/v2/pkg/view"
	"github.com/ReanSn0w/marker/pkg/mk"
)

func main() {
	previewElement(
		mk.Text("Hello World!"),
		mk.NewTag(false, "div")(
			mk.Text("Hi, how are you?"),
		)(
			mk.WithRuleStyle(
				"media (min-width: 768px) and (max-width: 1024px)",
				"",
				mk.Declarations{
					"animation":          "ease",
					"animation-duration": ".5s",
					"background-color":   "blue",
				},
			),
			mk.WithRuleStyle(
				"media (min-width: 768px) and (max-width: 1024px)",
				":hover",
				mk.Declarations{
					"background-color": "red",
				},
			),
			mk.Class("my_element"),
		),
	)
}

func previewElement(item ...view.View) {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mk.NewPage(mk.NewTag(false, "title")(mk.Text("Preview Page")))(item...).
			Write(r.Context(), http.StatusOK, w)
	}))
}

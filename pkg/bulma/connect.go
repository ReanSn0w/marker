package bulma

import (
	"github.com/ReanSn0w/gew/v2/pkg/view"
	"github.com/ReanSn0w/marker/pkg/marker"
)

func Connect() {
	ConnectLocal("https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css")
}

func ConnectLocal(link string) view.View {
	return marker.Link()(
		marker.SetAttribute("rel", "stylesheet"),
		marker.SetAttribute("href", link),
	)
}

func Viewport() view.View {
	return marker.Meta()(
		marker.SetAttribute("name", "viewport"),
		marker.SetAttribute("content", "width=device-width, initial-scale=1"),
	)
}

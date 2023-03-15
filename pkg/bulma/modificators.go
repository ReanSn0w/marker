package bulma

import (
	"github.com/ReanSn0w/gew/v2/pkg/view"
	"github.com/ReanSn0w/marker/pkg/marker"
)

func Columns() view.Modificator {
	return func(v view.View) view.View {
		return marker.PlaceInContainer("div")(
			marker.SetClass("columns")(v),
		)
	}
}

func Column() view.Modificator {
	return func(v view.View) view.View {
		return marker.PlaceInContainer("div")(
			marker.SetClass("column")(v),
		)
	}
}

func Block() view.Modificator {
	return func(v view.View) view.View {
		return marker.PlaceInContainer("div")(
			marker.SetClass("block")(v),
		)
	}
}

func Box() view.Modificator {
	return func(v view.View) view.View {
		return marker.PlaceInContainer("div")(
			marker.SetClass("box")(v),
		)
	}
}

func Tag() view.Modificator {
	return func(v view.View) view.View {
		return marker.PlaceInContainer("span")(
			marker.SetClass("tag")(v),
		)
	}
}

func TagsList() view.Modificator {
	return func(v view.View) view.View {
		return marker.PlaceInContainer("div")(
			marker.SetClass("tags")(v),
		)
	}
}

func HasAddons() view.Modificator {
	return marker.AppendClass("has-addons")
}

func Title() view.Modificator {
	return marker.AppendClass("title")
}

func Subtitle() view.Modificator {
	return marker.AppendClass("subtitle")
}

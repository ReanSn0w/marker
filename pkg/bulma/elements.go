package bulma

import (
	"github.com/ReanSn0w/gew/v2/pkg/view"
	"github.com/ReanSn0w/marker/pkg/marker"
)

func Field(title string) func(elements ...view.View) view.ModificationApplyer {
	return func(elements ...view.View) view.ModificationApplyer {
		return view.NewModificationApplyer(
			marker.Div(
				marker.Label(marker.Text(title))(marker.SetClass("label")),
				marker.Div(elements...)(marker.SetClass("control")),
			)(marker.SetClass("field")),
		)
	}
}

func Content(elements ...view.View) view.ModificationApplyer {
	return view.NewModificationApplyer(
		marker.Div(elements...)(marker.SetClass("content")),
	)
}

func Button(text string) view.ModificationApplyer {
	return view.NewModificationApplyer(
		marker.Button(marker.Text(text))(marker.SetClass("button")),
	)
}

func Link(text string) view.ModificationApplyer {
	return view.NewModificationApplyer(
		marker.A(marker.Text(text))(marker.SetClass("link")),
	)
}

func Icon(src string) view.ModificationApplyer {
	return view.NewModificationApplyer(
		marker.Img()(
			marker.SetAttribute("src", src),
		),
	)
}

func Figure(elements ...view.View) view.ModificationApplyer {
	return view.NewModificationApplyer(
		marker.Figure(elements...)(marker.SetClass("image")),
	)
}

// TODO: Add delete logic
func Notification(deletable bool) func(content ...view.View) view.ModificationApplyer {
	return func(content ...view.View) view.ModificationApplyer {
		return view.NewModificationApplyer(
			marker.Div(
				marker.Button()(marker.SetClass("delete")),
				view.Group(content...),
			)(marker.SetClass("notification")),
		)
	}
}

// Progress
//
// default value: 0
// default max: 100
func Pregress() view.ModificationApplyer {
	return view.NewModificationApplyer(
		marker.Progress()(
			marker.SetClass("progress"),
			marker.SetAttribute("value", "0"),
			marker.SetAttribute("max", "100"),
		),
	)
}

type BreadCrumbItem struct {
	Link  string
	Title string
	Icon  view.View
}

func Breadcrumb(items ...BreadCrumbItem) view.ModificationApplyer {
	return view.NewModificationApplyer(
		view.For(len(items), func(i int) view.View {
			return marker.A(
				view.If(items[i].Icon != nil)(view.Group(
					marker.Span(items[i].Icon)(marker.SetClass("icon is-small")),
					marker.Span(marker.Text(items[i].Title)),
				)),
				view.If(items[i].Icon == nil)(marker.Text(items[i].Title)),
			)(
				marker.SetAttribute("href", items[i].Link),
				marker.PlaceInContainer("li"),
				view.If(i == len(items)-1, marker.SetClass("is-active")),
			)
		})(
			marker.PlaceInContainer("ul"),
			marker.PlaceInContainer("nav"),
			marker.SetClass("breadcrumb"),
			marker.SetAttribute("aria-label", "breadcrumbs"),
		),
	)
}

const (
	MainCard        = ""              // the main container
	CardHeader      = "-header"       // a horizontal bar with a shadow
	CardHeaderTitle = "-header-title" // a left-aligned bold text
	CardHeaderIcon  = "-header-icon"  // a placeholder for an icon
	CardImage       = "-image"        // a fullwidth container for a responsive image
	CardContent     = "-content"      // a multi-purpose container for any other element
	CardFooter      = "-footer"       // a horizontal list of controls
	CardFooterItem  = "-footer-item"  // a repeatable list item
)

func Card(element string) func(elements ...view.View) view.ModificationApplyer {
	return func(elements ...view.View) view.ModificationApplyer {
		return view.NewModificationApplyer(
			marker.Div(elements...)(marker.SetClass("card" + element)),
		)
	}
}

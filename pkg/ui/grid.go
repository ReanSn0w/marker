package ui

import (
	"strings"

	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/mk"
	"github.com/ReanSn0w/marker/pkg/tag"
)

func NewGridPreference(v, h string, columns ...string) GridPreference {
	return GridPreference{
		Vertical:   v,
		Horizontal: h,
		Columns:    columns,
	}
}

type GridPreference struct {
	Vertical   string
	Horizontal string
	Columns    []string
}

func Grid(pref GridPreference) func(content ...view.View) view.Use {
	return func(content ...view.View) view.Use {
		return view.NewView(
			tag.Div(content...)(
				mk.WithStyle().
					WithValue("display", "grid").
					WithValue("grid-template-columns", strings.Join(pref.Columns, " ")).
					WithValue("grid-column-gap", pref.Horizontal).
					WithValue("grid-row-gap", pref.Vertical).
					Extract(),
			),
		)
	}
}

package ui

import (
	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/mk"
)

func Width(val string) view.Mod {
	return mk.WithStyle("", mk.Declarations{"width": mk.Value(val)})
}

func Height(val string) view.Mod {
	return mk.WithStyle("", mk.Declarations{"height": mk.Value(val)})
}

func BackgroundColor(val string) view.Mod {
	return mk.WithStyle("", mk.Declarations{"background-color": mk.Value(val)})
}

package ui

import "github.com/ReanSn0w/gew/v3/pkg/view"

func PlaceInBlock() view.Mod {
	return func(v view.View) view.View {
		return Block(v)
	}
}

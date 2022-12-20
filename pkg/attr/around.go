package attr

import (
	"github.com/ReanSn0w/gew/v2/pkg/view"
)

func If(condition bool, modificator view.Modificator) view.Modificator {
	if !condition {
		return nil
	}

	return modificator
}

func Group(modificators ...view.Modificator) view.Modificator {
	return func(tag view.View) view.View {
		for _, mod := range modificators {
			if mod == nil {
				continue
			}

			tag = mod(tag)
		}

		return tag
	}
}

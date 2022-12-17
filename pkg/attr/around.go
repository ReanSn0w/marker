package attr

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/marker/pkg/base"
)

func If(condition bool, modificator base.TagModificator) base.TagModificator {
	if !condition {
		return nil
	}

	return modificator
}

func Group(modificators ...base.TagModificator) base.TagModificator {
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

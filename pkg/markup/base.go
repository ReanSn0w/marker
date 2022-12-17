package markup

import (
	"fmt"

	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/marker/pkg/base"
)

func Doctype() view.View {
	return view.External("<!DOCTYPE html>")
}

func Comment(val string) view.View {
	return view.External("<!-- " + val + " -->")
}

func Text(val string) base.ContainerModificator {
	return func(modificators ...base.TagModificator) view.View {
		if len(modificators) == 0 {
			return view.External(val)
		}

		return base.Container("span", view.External(val))(modificators...)
	}
}

func Snippet(snippet string, values ...interface{}) view.View {
	return view.External(fmt.Sprintf(snippet, values...))
}

// Функция для создания контейнера с заданынм тегом, содержимым и необходимыми модифиикациями
func basicContainerTag(tag base.ContainerTagName, content ...view.View) base.ContainerModificator {
	return func(mod ...base.TagModificator) view.View {
		return base.Container(tag, view.Group(content...))(mod...)
	}
}

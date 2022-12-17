package base

import "github.com/ReanSn0w/gew/pkg/view"

type TagModificator func(view.View) view.View

func SetAttribute(key, value string) TagModificator {
	return func(v view.View) view.View {
		item := getAttributeEditor(v)
		item.Set(key, value)
		return v
	}
}

func AppendAttribute(key, value string) TagModificator {
	return func(v view.View) view.View {
		item := getAttributeEditor(v)
		item.Append(key, value)
		return v
	}
}

func PrepareAttribute(key string, prepare func(string) string) TagModificator {
	return func(v view.View) view.View {
		item := getAttributeEditor(v)
		current := item.Get(key)
		item.Set(key, prepare(current))
		return v
	}
}

func DeleteAttribute(key string) TagModificator {
	return func(v view.View) view.View {
		item := getAttributeEditor(v)
		item.Delete(key)
		return v
	}
}

func PlaceInContainer(tag ContainerTagName) TagModificator {
	return func(v view.View) view.View {
		return Container(tag, v)()
	}
}

func getAttributeEditor(v view.View) AttributeEditor {
	item, ok := v.(AttributeEditor)
	if !ok {
		item = Container("div", v)().(*containerTag)
	}

	return item
}

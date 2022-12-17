package attr

import "github.com/ReanSn0w/marker/pkg/base"

func AccessKey(val string) base.TagModificator {
	return base.SetAttribute("accesskey", val)
}

func Class(val string) base.TagModificator {
	return base.SetAttribute("class", val)
}

func AppendClass(val string) base.TagModificator {
	return base.AppendAttribute("class", val)
}

func ContentEditable(val string) base.TagModificator {
	return base.SetAttribute("contenteditable", val)
}

func ContextMenu(val string) base.TagModificator {
	return base.SetAttribute("contextmenu", val)
}

func Dir(val string) base.TagModificator {
	return base.SetAttribute("dir", val)
}

func Draggable(val string) base.TagModificator {
	return base.SetAttribute("draggable", val)
}

func DropZone(val string) base.TagModificator {
	return base.SetAttribute("dropzone", val)
}

func Hidden(val string) base.TagModificator {
	return base.SetAttribute("hidden", val)
}

func Id(val string) base.TagModificator {
	return base.SetAttribute("id", val)
}

func Lang(val string) base.TagModificator {
	return base.SetAttribute("lang", val)
}

func SpellCheck(val string) base.TagModificator {
	return base.SetAttribute("spellcheck", val)
}

func Style(val string) base.TagModificator {
	return base.SetAttribute("style", val+";")
}

func AppendStyle(val string) base.TagModificator {
	return base.AppendAttribute("style", val+";")
}

func TabIndex(val string) base.TagModificator {
	return base.SetAttribute("tabindex", val)
}

func Title(val string) base.TagModificator {
	return base.SetAttribute("title", val)
}

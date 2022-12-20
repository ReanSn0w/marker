package attr

import (
	"github.com/ReanSn0w/gew/v2/pkg/view"
	"github.com/ReanSn0w/marker/pkg/base"
)

func AccessKey(val string) view.Modificator {
	return base.SetAttribute("accesskey", val)
}

func Class(val string) view.Modificator {
	return base.SetAttribute("class", val)
}

func AppendClass(val string) view.Modificator {
	return base.AppendAttribute("class", val)
}

func ContentEditable(val string) view.Modificator {
	return base.SetAttribute("contenteditable", val)
}

func ContextMenu(val string) view.Modificator {
	return base.SetAttribute("contextmenu", val)
}

func Dir(val string) view.Modificator {
	return base.SetAttribute("dir", val)
}

func Draggable(val string) view.Modificator {
	return base.SetAttribute("draggable", val)
}

func DropZone(val string) view.Modificator {
	return base.SetAttribute("dropzone", val)
}

func Hidden(val string) view.Modificator {
	return base.SetAttribute("hidden", val)
}

func Id(val string) view.Modificator {
	return base.SetAttribute("id", val)
}

func Lang(val string) view.Modificator {
	return base.SetAttribute("lang", val)
}

func SpellCheck(val string) view.Modificator {
	return base.SetAttribute("spellcheck", val)
}

func Style(val string) view.Modificator {
	return base.SetAttribute("style", val+";")
}

func AppendStyle(val string) view.Modificator {
	return base.AppendAttribute("style", val+";")
}

func TabIndex(val string) view.Modificator {
	return base.SetAttribute("tabindex", val)
}

func Title(val string) view.Modificator {
	return base.SetAttribute("title", val)
}

package attr

import (
	"github.com/ReanSn0w/gew/v2/pkg/view"
	"github.com/ReanSn0w/marker/pkg/base"
)

func OnChange(val string) view.Modificator {
	return base.SetAttribute("onchange", val)
}

func OnClick(val string) view.Modificator {
	return base.SetAttribute("onclick", val)
}

func Ondblclick(val string) view.Modificator {
	return base.SetAttribute("ondblclick", val)
}

func OnFocus(val string) view.Modificator {
	return base.SetAttribute("onfocus", val)
}

func OnKeyDown(val string) view.Modificator {
	return base.SetAttribute("onkeydown", val)
}

func OnKeyPress(val string) view.Modificator {
	return base.SetAttribute("onkeypress", val)
}

func OnKeyUp(val string) view.Modificator {
	return base.SetAttribute("onkeyup", val)
}

func OnLoad(val string) view.Modificator {
	return base.SetAttribute("onload", val)
}

func OnMouseDown(val string) view.Modificator {
	return base.SetAttribute("onmousedown", val)
}

func OnMouseMove(val string) view.Modificator {
	return base.SetAttribute("onmousemove", val)
}

func OnMouseOut(val string) view.Modificator {
	return base.SetAttribute("onmouseout", val)
}

func OnMouseOver(val string) view.Modificator {
	return base.SetAttribute("onmouseover", val)
}

func OnMouseUp(val string) view.Modificator {
	return base.SetAttribute("onmouseup", val)
}

func OnReset(val string) view.Modificator {
	return base.SetAttribute("onreset", val)
}

func OnSelect(val string) view.Modificator {
	return base.SetAttribute("onselect", val)
}

func OnSubmit(val string) view.Modificator {
	return base.SetAttribute("onsubmit", val)
}

func OnUnload(val string) view.Modificator {
	return base.SetAttribute("onunload", val)
}

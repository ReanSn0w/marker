package attr

import "github.com/ReanSn0w/marker/pkg/base"

func OnChange(val string) base.TagModificator {
	return base.SetAttribute("onchange", val)
}

func OnClick(val string) base.TagModificator {
	return base.SetAttribute("onclick", val)
}

func Ondblclick(val string) base.TagModificator {
	return base.SetAttribute("ondblclick", val)
}

func OnFocus(val string) base.TagModificator {
	return base.SetAttribute("onfocus", val)
}

func OnKeyDown(val string) base.TagModificator {
	return base.SetAttribute("onkeydown", val)
}

func OnKeyPress(val string) base.TagModificator {
	return base.SetAttribute("onkeypress", val)
}

func OnKeyUp(val string) base.TagModificator {
	return base.SetAttribute("onkeyup", val)
}

func OnLoad(val string) base.TagModificator {
	return base.SetAttribute("onload", val)
}

func OnMouseDown(val string) base.TagModificator {
	return base.SetAttribute("onmousedown", val)
}

func OnMouseMove(val string) base.TagModificator {
	return base.SetAttribute("onmousemove", val)
}

func OnMouseOut(val string) base.TagModificator {
	return base.SetAttribute("onmouseout", val)
}

func OnMouseOver(val string) base.TagModificator {
	return base.SetAttribute("onmouseover", val)
}

func OnMouseUp(val string) base.TagModificator {
	return base.SetAttribute("onmouseup", val)
}

func OnReset(val string) base.TagModificator {
	return base.SetAttribute("onreset", val)
}

func OnSelect(val string) base.TagModificator {
	return base.SetAttribute("onselect", val)
}

func OnSubmit(val string) base.TagModificator {
	return base.SetAttribute("onsubmit", val)
}

func OnUnload(val string) base.TagModificator {
	return base.SetAttribute("onunload", val)
}

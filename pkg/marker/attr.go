package marker

import (
	"github.com/ReanSn0w/gew/v2/pkg/view"
)

// хранилище для атрибутов html тега
type Attributes map[string]string

// Функция строит строку атрибутов для html тега
func (a Attributes) build() string {
	attributes := ""

	for key, value := range a {
		if value == "" {
			attributes += key + " "
		} else {
			attributes += key + "=\"" + value + "\" "
		}
	}

	return attributes
}

// Интерфейс, которому должны удовлетворять структуры для хранения тегов
type AttributeEditor interface {
	Get(string) string
	Set(string, string)
	Append(string, string)
	Delete(string)
}

// MARK: Attribute Modificator
// Ниже описан ряд функций созданный для базового редактирования атрибутов html тега

// Функция для принудительной установки атрибута
// в случае если в у тега уже был установлен данный атрибут он будет перезаписан
func SetAttribute(key, value string) view.Modificator {
	return func(v view.View) view.View {
		item := getAttributeEditor(v)
		item.Set(key, value)
		return v
	}
}

// Добавление атрибута к уже существующим
// такое значение добавляется к текущему значению атрибута через пробел
func AppendAttribute(key, value string) view.Modificator {
	return func(v view.View) view.View {
		item := getAttributeEditor(v)
		item.Append(key, value)
		return v
	}
}

// Функция для подготовки атрибута
// в качестве параметра принимает функцию, которая принимает текущее значение атрибута
// и возвращает новое значение
func PrepareAttribute(key string, prepare func(string) string) view.Modificator {
	return func(v view.View) view.View {
		item := getAttributeEditor(v)
		current := item.Get(key)
		item.Set(key, prepare(current))
		return v
	}
}

// Функция для удаления атрибута
func DeleteAttribute(key string) view.Modificator {
	return func(v view.View) view.View {
		item := getAttributeEditor(v)
		item.Delete(key)
		return v
	}
}

func getAttributeEditor(v view.View) AttributeEditor {
	item, ok := v.(AttributeEditor)
	if !ok {
		item = Container("div", v)().(*containerTag)
	}

	return item
}

// MARK: - Common Attributes
// модификаторы для выставления общих атрибутов html тега

func SetAccessKey(val string) view.Modificator {
	return SetAttribute("accesskey", val)
}

func SetClass(val string) view.Modificator {
	return SetAttribute("class", val)
}

func AppendClass(val string) view.Modificator {
	return AppendAttribute("class", val)
}

func SetContentEditable(val string) view.Modificator {
	return SetAttribute("contenteditable", val)
}

func SetContextMenu(val string) view.Modificator {
	return SetAttribute("contextmenu", val)
}

func SetDir(val string) view.Modificator {
	return SetAttribute("dir", val)
}

func SetDraggable(val string) view.Modificator {
	return SetAttribute("draggable", val)
}

func SetDropZone(val string) view.Modificator {
	return SetAttribute("dropzone", val)
}

func SetHidden(val string) view.Modificator {
	return SetAttribute("hidden", val)
}

func SetId(val string) view.Modificator {
	return SetAttribute("id", val)
}

func SetLang(val string) view.Modificator {
	return SetAttribute("lang", val)
}

func SetSpellCheck(val string) view.Modificator {
	return SetAttribute("spellcheck", val)
}

func SetStyle(val string) view.Modificator {
	return SetAttribute("style", val+";")
}

func SetAppendStyle(val string) view.Modificator {
	return AppendAttribute("style", val+";")
}

func SetTabIndex(val string) view.Modificator {
	return SetAttribute("tabindex", val)
}

func SetTitle(val string) view.Modificator {
	return SetAttribute("title", val)
}

// MARK: - Events Attributes
// модификаторы для выставления атрибутов событий html тега

func SetOnChange(val string) view.Modificator {
	return SetAttribute("onchange", val)
}

func SetOnClick(val string) view.Modificator {
	return SetAttribute("onclick", val)
}

func SetOndblclick(val string) view.Modificator {
	return SetAttribute("ondblclick", val)
}

func SetOnFocus(val string) view.Modificator {
	return SetAttribute("onfocus", val)
}

func SetOnKeyDown(val string) view.Modificator {
	return SetAttribute("onkeydown", val)
}

func SetOnKeyPress(val string) view.Modificator {
	return SetAttribute("onkeypress", val)
}

func SetOnKeyUp(val string) view.Modificator {
	return SetAttribute("onkeyup", val)
}

func SetOnLoad(val string) view.Modificator {
	return SetAttribute("onload", val)
}

func SetOnMouseDown(val string) view.Modificator {
	return SetAttribute("onmousedown", val)
}

func SetOnMouseMove(val string) view.Modificator {
	return SetAttribute("onmousemove", val)
}

func SetOnMouseOut(val string) view.Modificator {
	return SetAttribute("onmouseout", val)
}

func SetOnMouseOver(val string) view.Modificator {
	return SetAttribute("onmouseover", val)
}

func SetOnMouseUp(val string) view.Modificator {
	return SetAttribute("onmouseup", val)
}

func SetOnReset(val string) view.Modificator {
	return SetAttribute("onreset", val)
}

func SetOnSelect(val string) view.Modificator {
	return SetAttribute("onselect", val)
}

func SetOnSubmit(val string) view.Modificator {
	return SetAttribute("onsubmit", val)
}

func SetOnUnload(val string) view.Modificator {
	return SetAttribute("onunload", val)
}

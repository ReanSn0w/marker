package primitive

import (
	"context"
	"io"
	"time"
)

var (
	globalContextKey = &globalDataCtx{}
)

func Get(ctx context.Context) *GlobalData {
	return ctx.Value(globalContextKey).(*GlobalData)
}

func Set(ctx context.Context) context.Context {
	return context.WithValue(ctx, globalContextKey, &GlobalData{
		stylesEnabled: true,
		styles:        make(atRules),
	})
}

type (
	globalDataCtx struct{}

	GlobalData struct {
		stylesEnabled bool
		styles        atRules
		element       *ElementData
	}

	ElementData struct {
		class      string
		attributes map[Key]Value
		style      atRules
	}
)

func (gd *GlobalData) ElementData() *ElementData {
	if gd.element == nil {
		gd.element = &ElementData{
			class:      newClass(time.Now()),
			attributes: make(map[Key]Value),
			style:      make(atRules),
		}
	}

	return gd.element
}

func (gd *GlobalData) ExtractData() map[Key]Value {
	if gd.stylesEnabled {
		gd.appendStyles(gd.element.class, gd.element.style)
	}

	gd.element.attributes["class"] = Value("G" + gd.element.class)
	res := gd.element.attributes
	gd.element = nil
	return res
}

func (gd *GlobalData) EnableStyle() {
	gd.stylesEnabled = true
}

func (gd *GlobalData) DisableStyle() {
	gd.stylesEnabled = false
}

func (gd *GlobalData) WriteStyle(wr io.Writer) {
	gd.styles.write(wr)
}

func (ed *ElementData) GetAttributes() string {
	result := ""

	for k, v := range ed.attributes {
		if v == "" {
			result += " " + string(k)
			continue
		} else {
			result += " " + string(k) + "=" + "\"" + string(v) + "\""
		}
	}

	return result
}

func (ed *ElementData) GetAttribute(k Key) Value {
	return ed.attributes[k]
}

func (ed *ElementData) SetAttribute(k Key, v Value) {
	ed.attributes[k] = v
}

func (ed *ElementData) DropAttribute(k Key) {
	delete(ed.attributes, k)
}

func (ed *ElementData) SetClass(class string) {
	ed.class = class
}

func (ed *ElementData) SetStyle(rule AtRule, selector Selector, key Key, value Value) {
	ed.style.append(rule, selector, key, value)
}

func (gd *GlobalData) appendStyles(class string, styles atRules) {
	for rule, selectors := range styles {
		for selector, declarations := range selectors {
			for key, value := range declarations {
				gd.styles.append(rule, Selector("G"+class)+selector, key, value)
			}
		}
	}
}

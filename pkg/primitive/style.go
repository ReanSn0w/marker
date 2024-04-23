package primitive

import (
	"fmt"
	"io"
)

const (
	emptyAtRule = AtRule("")
)

type (
	Selector  string
	AtRule    string
	selectors map[Selector]attributes
	atRules   map[AtRule]selectors
)

func (a atRules) append(rule AtRule, selector Selector, key Key, value Value) {
	_, ok := a[rule]
	if !ok {
		selectorStorage := make(selectors)
		a[rule] = selectorStorage
	}

	declarationsStorage, ok := a[rule][selector]
	if !ok {
		declarationsStorage = make(attributes)
		a[rule][selector] = declarationsStorage
	}

	declarationsStorage[key] = value
}

func (a atRules) write(wr io.Writer) {
	for atRule, selectors := range a {
		if atRule != emptyAtRule {
			wr.Write([]byte("@" + atRule + "{"))
		}

		for selector, declarations := range selectors {
			wr.Write([]byte(selector + "{"))

			for property, value := range declarations {
				wr.Write([]byte(fmt.Sprintf("%v: %v; ", property, value)))
			}

			wr.Write([]byte("}"))
		}

		if atRule != emptyAtRule {
			wr.Write([]byte("}"))
		}
	}
}

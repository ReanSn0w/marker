package markup

import (
	"fmt"

	"github.com/ReanSn0w/gew/v2/pkg/view"
)

func Doctype() view.View {
	return view.External("<!DOCTYPE html>")
}

func Comment(val string) view.View {
	return view.External("<!-- " + val + " -->")
}

func Text(val string) view.View {
	return view.External(val)()
}

func Snippet(snippet string, values ...interface{}) view.View {
	return view.External(fmt.Sprintf(snippet, values...))
}

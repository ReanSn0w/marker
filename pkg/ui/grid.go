package ui

import (
	"strings"

	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/mk"
)

func Grid() *grid {
	return &grid{}
}

type grid struct {
	rule     string
	selector string
	vGap     Size
	hGap     Size
	columns  []Size
}

func (g *grid) columnsString() string {
	sizes := []string{}

	for _, size := range g.columns {
		sizes = append(sizes, size.Size())
	}

	return strings.Join(sizes, " ")
}

func (g *grid) Columns(sizes ...Size) *grid {
	g.columns = append(g.columns, sizes...)
	return g
}

func (g *grid) Rule(rule string) *grid {
	g.rule = rule
	return g
}

func (g *grid) Selector(selector string) *grid {
	g.selector = selector
	return g
}

func (g *grid) VGap(size Size) *grid {
	g.vGap = size
	return g
}

func (g *grid) HGap(size Size) *grid {
	g.hGap = size
	return g
}

func (g *grid) Extract() view.Mod {
	return mk.WithStyle().
		Selector(g.selector).
		Rule(g.rule).
		Value("display", "grid").
		Value("grid-template-columns", g.columnsString()).
		Value("grid-column-gap", g.hGap.Size()).
		Value("grid-row-gap", g.vGap.Size()).
		Extract()
}

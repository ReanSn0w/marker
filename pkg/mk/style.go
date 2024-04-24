package mk

import (
	"context"

	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/common"
)

func WithStyle() *StyleConfig {
	return &StyleConfig{
		values: make(map[string]string),
	}
}

type StyleConfig struct {
	rule     string
	selector string
	values   map[string]string
}

func (s *StyleConfig) Rule(rule string) *StyleConfig {
	s.rule = rule
	return s
}

func (s *StyleConfig) Selector(selector string) *StyleConfig {
	s.selector = selector
	return s
}

func (s *StyleConfig) Value(key, val string) *StyleConfig {
	s.values[key] = val
	return s
}

func (s *StyleConfig) Extract() view.Mod {
	return view.ContextModificator(func(ctx context.Context) context.Context {
		elementData := common.Get(ctx).ElementData()

		for key, val := range s.values {
			elementData.SetStyle(
				common.AtRule(s.rule),
				common.Selector(s.selector),
				common.Key(key),
				common.Value(val))
		}

		return ctx
	})
}

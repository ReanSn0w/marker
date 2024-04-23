package mk

import (
	"context"

	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/primitive"
)

func WithStyle() *StyleConfig {
	return &StyleConfig{
		Value: make(map[string]string),
	}
}

type StyleConfig struct {
	Rule     string
	Selector string
	Value    map[string]string
}

func (s *StyleConfig) WithRule(rule string) *StyleConfig {
	s.Rule = rule
	return s
}

func (s *StyleConfig) WithSelector(selector string) *StyleConfig {
	s.Selector = selector
	return s
}

func (s *StyleConfig) WithValue(key, val string) *StyleConfig {
	s.Value[key] = val
	return s
}

func (s *StyleConfig) Extract() view.Mod {
	return view.ContextModificator(func(ctx context.Context) context.Context {
		elementData := primitive.Get(ctx).ElementData()

		for key, val := range s.Value {
			elementData.SetStyle(
				primitive.AtRule(s.Rule),
				primitive.Selector(s.Selector),
				primitive.Key(key),
				primitive.Value(val))
		}

		return ctx
	})
}

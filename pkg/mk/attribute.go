package mk

import (
	"context"

	"github.com/ReanSn0w/gew/v3/pkg/view"
	"github.com/ReanSn0w/marker/pkg/primitive"
)

// Class - Устанавливает класс для ближайшего
// ребенка в контекст выполнения
func Class(value string) view.Mod {
	return view.ContextModificator(func(ctx context.Context) context.Context {
		primitive.Get(ctx).ElementData().SetClass(value)
		return ctx
	})
}

// Set - Устанавливает атрибут к ближайшему
// ребенку по дереву построения
func Set(key string, val string) view.Mod {
	return view.ContextModificator(func(ctx context.Context) context.Context {
		primitive.Get(ctx).ElementData().SetAttribute(primitive.Key(key), primitive.Value(val))
		return ctx
	})
}

// Delete - Производит удаление атрибута
// для ближайшего ребенка в дереве построения
func Delete(key string) view.Mod {
	return view.ContextModificator(func(ctx context.Context) context.Context {
		primitive.Get(ctx).ElementData().DropAttribute(primitive.Key(key))
		return ctx
	})
}

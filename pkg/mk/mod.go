package mk

import (
	"context"
	"fmt"

	"github.com/ReanSn0w/gew/v3/pkg/view"
)

// Class - Устанавливает класс для ближайшего
// ребенка в контекст выполнения
func Class(value string) view.Mod {
	return view.ContextModificator(func(ctx context.Context) context.Context {
		getClassStorage(ctx).set(class(value))
		return ctx
	})
}

// Set - Устанавливает атрибут к ближайшему
// ребенку по дереву построения
func Set(key Key, val Value) view.Mod {
	return func(v view.View) view.View {
		return view.Closure(func(ctx context.Context) view.View {
			value := ctx.Value(attrCtxKey)
			attr, ok := value.(attributes)
			if !ok {
				attr = make(attributes)
				attr[key] = val
				view.Context(attrCtxKey, attr)(v)
			}

			attr[key] = val
			return v
		})
	}
}

// Push - Добавляет значение атрибута
// к ближайшему ребенку по дереву построения
func Push(key Key, val Value) view.Mod {
	return func(v view.View) view.View {
		return view.Closure(func(ctx context.Context) view.View {
			value := ctx.Value(attrCtxKey)
			attr, ok := value.(attributes)
			if !ok {
				attr = make(attributes)
				attr[key] = val
				view.Context(attrCtxKey, attr)(v)
			}

			attr[key] += val
			return v
		})
	}
}

// Delete - Производит удаление атрибута
// для ближайшего ребенка в дереве построения
func Delete(key Key) view.Mod {
	return func(v view.View) view.View {
		return view.Closure(func(ctx context.Context) view.View {
			value := ctx.Value(attrCtxKey)
			attr, ok := value.(attributes)

			if ok {
				delete(attr, key)
			}

			return v
		})
	}
}

// WithStyle - применяет стили к ближайшему дочернему
// элементу в DOM дереве
func WithStyle(selector string, dec Declarations) view.Mod {
	return WithRuleStyle(string(emptyAtRule), selector, dec)
}

// WithRuleStyle - применяет стили к ближайшему дочернему
// элементу в DOM дереве. Стили применяемые данной фенкцией
// могут быть вложенны в @ блоки
func WithRuleStyle(rule string, sel string, dec Declarations) view.Mod {
	return func(v view.View) view.View {
		return view.Closure(func(ctx context.Context) view.View {
			style := getStyle(ctx)
			class := getClassStorage(ctx).get()
			s := selector(fmt.Sprintf(".G%s%s", class, sel))

			for p, v := range dec {
				style.append(atRule(rule), s, p, v)
			}

			// Возврат view осуществляется с установкой класса
			// для того что бы сгенерированное
			// значение осталось за текущим элементом
			return view.NewView(v)(
				Class(string(class)),
			)
		})
	}
}

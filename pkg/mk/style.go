package mk

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"sync/atomic"
	"time"

	"github.com/ReanSn0w/gew/v3/pkg/view"
)

// MARK: Class
//
//
//
//
//

var (
	emptyClass      class = ""
	objectIDCounter       = readRandomUint32()
	classCtxKey           = "class_ctx_key"
)

type (
	class string

	ClassStorage struct {
		enabled bool
		current *class
	}
)

func addClassStorage(ctx context.Context) context.Context {
	return context.WithValue(ctx, classCtxKey, &ClassStorage{enabled: true})
}

func getClassStorage(ctx context.Context) *ClassStorage {
	value := ctx.Value(classCtxKey)
	return value.(*ClassStorage)
}

func (c *ClassStorage) get() class {
	if c.current == nil {
		return emptyClass
	}

	return *c.current
}

func (c *ClassStorage) set(class class) {
	c.current = &class
}

func (c *ClassStorage) reset() {
	cl := newClass(time.Now())
	c.current = &cl
}

func (c *ClassStorage) enable() {
	c.enabled = true
}

func (c *ClassStorage) disable() {
	c.enabled = false
}

// // getClass - Производит получение класса элемента из котекста
// // в случае отсутствия класс, генерируется новое уникальное
// // значение класса.
// //
// // [ВАЖНО!] если класс небыл установлен заранее, функция будет возвращать новое название
// // класса во время каждого вызова
// func getClass(ctx context.Context) class {
// 	val := ctx.Value(classCtxKey)
// 	c, ok := val.(class)

// 	if !ok || c == emptyClass {
// 		return newClass(time.Now())
// 	}

// 	return c
// }

// func ignoreClass(ctx context.Context) bool {
// 	_, ok := ctx.Value(ignoreClassCtxKey).(struct{})
// 	return ok
// }

func newClass(timestamp time.Time) class {
	var b [7]byte
	binary.BigEndian.PutUint32(b[0:4], uint32(timestamp.Unix()))
	putUint24(b[4:7], atomic.AddUint32(&objectIDCounter, 1))
	return class(hex.EncodeToString(b[:]))
}

func readRandomUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot initialize objectid package with crypto.rand.Reader: %w", err))
	}

	return (uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
}

func putUint24(b []byte, v uint32) {
	b[0] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[2] = byte(v)
}

// MARK: Attributes
const (
	EmptyValue Value = ""
)

var (
	attrCtxKey = "attribute_ctx_key"
)

type (
	attributes map[Key]Value
	Key        string
	Value      string
)

func (a attributes) write(wr io.Writer) {
	for key, value := range a {
		wr.Write([]byte(fmt.Sprintf(" %v", key)))
		if value != EmptyValue {
			wr.Write([]byte(fmt.Sprintf("=\"%v\"", value)))
		}
	}
}

func getAttributes(ctx context.Context) attributes {
	value := ctx.Value(attrCtxKey)
	attr, ok := value.(attributes)
	if !ok {
		attr = make(attributes)
	}

	return attr
}

// MARK: Style
//
//
//
//

const (
	emptyAtRule atRule = "empty"
)

var (
	styleCtxKey = "style_ctx_key"
)

func newStyle(ctx context.Context) (context.Context, *Style) {
	style := plainNewStyle()
	ctx = style.set(ctx)
	return ctx, style
}

func plainNewStyle() *Style {
	return &Style{styles: make(atRules)}
}

func getStyle(ctx context.Context) *Style {
	value := ctx.Value(styleCtxKey)
	style, ok := value.(*Style)
	if !ok {
		return plainNewStyle()
	}

	return style
}

type (
	Property string

	Style struct {
		styles atRules
	}

	selector     string
	atRule       string
	Declarations map[Property]Value
	selectors    map[selector]Declarations
	atRules      map[atRule]selectors
)

// Body implements view.View interface
func (s *Style) Body(ctx context.Context) view.View {
	return nil
}

func (s *Style) build(wr io.Writer) {
	wr.Write([]byte("<style>"))

	for atRule, selectors := range s.styles {
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

	wr.Write([]byte("</style>"))
}

func (s *Style) append(rule atRule, selector selector, prop Property, val Value) {
	_, ok := s.styles[rule]
	if !ok {
		s.styles[rule] = make(selectors)
	}

	_, ok = s.styles[rule][selector]
	if !ok {
		s.styles[rule][selector] = make(Declarations)
	}

	s.styles[rule][selector][prop] = val
}

func (s *Style) set(ctx context.Context) context.Context {
	return context.WithValue(ctx, styleCtxKey, s)
}

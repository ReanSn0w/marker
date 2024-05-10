package page

import (
	"github.com/ReanSn0w/gew/v3/pkg/view"
)

func New(title string) *Page {
	return &Page{
		lang:  Russian,
		title: title,
		head:  make([]view.View, 0),
		body:  make([]view.View, 0),
	}
}

type Page struct {
	lang  Lang
	title string

	head []view.View
	body []view.View
	mod  []view.Mod
}

func (p *Page) Lang(lang Lang) *Page {
	p.lang = lang
	return p
}

func (p *Page) Head(head ...view.View) *Page {
	p.head = head
	return p
}

func (p *Page) Content(body ...view.View) *Page {
	p.body = body
	return p
}

func (p *Page) Mod(m ...view.Mod) *Page {
	p.mod = m
	return p
}

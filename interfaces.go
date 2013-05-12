package gsnack

import (
	ħ "net/http"
)

type Container interface {
	AddJs([]byte)
	AddCss([]byte)
	AddHtml([]byte)
}

type Plugger interface {
	Plug(c Container, vals map[string][]byte)
	Id() string
	Class() string
	Test(vals map[string][]byte, template string) string
	TestHandler(vals map[string][]byte, template string) func(w ħ.ResponseWriter, r *ħ.Request)
	TestServe(vals map[string][]byte, template string)
}

type Snacker interface {
	Test(vals map[string][]byte, template string) string
	WithId(id string) Plugger
	WithClass(class string) Plugger
}

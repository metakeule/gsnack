package gsnack

import (
	ŧ "fmt"
	f "github.com/metakeule/fastreplace"
	ħ "net/http"
)

var delimiter = []byte("@@")

type instance struct {
	*Snack
	html_  *f.FReplace
	css_   *f.FReplace
	js_    *f.FReplace
	id     string
	class  string
	Values map[string]string
}

func (ø *instance) Id() string    { return ø.id }
func (ø *instance) Class() string { return ø.class }

func (ø *instance) Plug(target Container, vals map[string][]byte) {
	vs := map[string][]byte{}
	for k, v := range ø.Defaults {
		vs[k] = v
	}
	for k, v := range vals {
		vs[k] = v
	}
	if ø.id != "" {
		vs["id"] = []byte(ø.id)
	}
	if ø.class != "" {
		vs["class"] = []byte(ø.class)
	}

	if ø.id != "" {
		vs["cssSelector"] = []byte("#" + ø.id)
		vs["bodyAttr"] = []byte(`id="` + ø.id + `"`)
	}
	if ø.class != "" {
		vs["cssSelector"] = []byte("." + ø.class)
		vs["bodyAttr"] = []byte(`class="` + ø.class + `"`)
	}
	if ø.Initializer != nil {
		ø.Initializer(ø, vs)
	}

	target.AddCss(ø.css_.ReplaceBytes(vs))
	target.AddHtml(ø.html_.ReplaceBytes(vs))
	target.AddJs(ø.js_.ReplaceBytes(vs))
}

func (ø *instance) Test(vals map[string][]byte, template string) string {
	c := &testContainer{}
	ø.Plug(c, vals)
	return c.Html(ø.Libs, template)
}

func (ø *instance) TestHandler(vals map[string][]byte, template string) func(w ħ.ResponseWriter, r *ħ.Request) {
	c := &testContainer{}
	ø.Plug(c, vals)
	s := c.Html(ø.Libs, template)
	return func(w ħ.ResponseWriter, r *ħ.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		ŧ.Fprintln(w, s)
	}
}

func (ø *instance) TestServe(vals map[string][]byte, template string) {
	h := ø.TestHandler(vals, template)
	ħ.HandleFunc("/", h)
	ŧ.Println("Serving at localhost:8080")
	ſ := ħ.ListenAndServe(":8080", nil)
	if ſ != nil {
		ŧ.Printf("Could not start server: %s\n", ſ.Error())
	}
}

package gsnack

import (
	ŧ "fmt"
	f "github.com/metakeule/fastreplace"
	ħ "net/http"
)

type layout struct {
	css          []byte
	js           []byte
	replacer     *f.FReplace
	replacements map[string][]byte
	NextKey      string
}

func Layout(template []byte) (ø *layout) {
	repl, _ := f.New(delimiter, template)
	ø = &layout{
		replacer:     repl,
		replacements: map[string][]byte{},
	}
	return
}

func (ø *layout) Serve(libs []*JsLib, snacks map[string]Plugger, vals map[string]map[string][]byte) {
	for k, pl := range snacks {
		ø.NextKey = k
		pl.Plug(ø, vals[k])
	}

	s := ø.Html(libs)
	handler := func(w ħ.ResponseWriter, r *ħ.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		ŧ.Fprintln(w, s)
	}

	ħ.HandleFunc("/", handler)
	ŧ.Println("Serving at localhost:8080")
	ſ := ħ.ListenAndServe(":8080", nil)
	if ſ != nil {
		ŧ.Printf("Could not start server: %s\n", ſ.Error())
	}
}

func (ø *layout) AddCss(css []byte) {
	ø.css = append(ø.css, css...)
	ø.css = append(ø.css, []byte("\n")...)
}

func (ø *layout) AddJs(js []byte) {
	ø.js = append(ø.js, js...)
	ø.js = append(ø.js, []byte("\n")...)
}

func (ø *layout) AddHtml(html []byte) {
	ø.replacements[ø.NextKey] = html
}

func (ø *layout) Html(libs []*JsLib) string {
	ø.replacements["css"] = ø.css
	ø.replacements["libs"] = []byte(jslibs(libs).Include())
	ø.replacements["script"] = ø.js
	return ø.replacer.ReplaceString(ø.replacements)
}

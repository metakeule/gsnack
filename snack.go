package gsnack

import (
	ŧ "fmt"
	f "github.com/metakeule/fastreplace"
	"strings"
)

type Snack struct {
	Dependancies map[string]Version // key: jslib name value: version of jslib
	Html         []byte
	Css          []byte
	Js           []byte
	Defaults     map[string][]byte // key -> defaults
	Initializer  func(i Plugger, vals map[string][]byte)
	Libs         []*JsLib
}

func NewSnack(dependancies map[string]Version) (ø *Snack) {
	ø = &Snack{
		Dependancies: dependancies,
		Html:         []byte{},
		Js:           []byte{},
		Css:          []byte{},
		Defaults:     map[string][]byte{},
	}
	return
}

func (ø *Snack) new(id string, class string) (p Plugger) {
	i := &instance{Snack: ø}
	i.id = id
	i.class = class
	i.html_, _ = f.NewBytes(delimiter, ø.Html)
	i.css_, _ = f.NewBytes(delimiter, ø.Css)
	i.js_, _ = f.NewBytes(delimiter, ø.Js)
	p = i
	return
}

func (ø *Snack) WithId(id string) (p Plugger) {
	return ø.new(id, "")
}

func (ø *Snack) WithClass(class string) (p Plugger) {
	return ø.new("", class)
}

// checks if we got parse errors
func (ø *Snack) ParseErrors() (errs []error) {
	errs = []error{}
	_, err := f.NewBytes(delimiter, ø.Html)
	if err != nil {
		errs = append(errs, err)
	}
	_, err = f.NewBytes(delimiter, ø.Css)
	if err != nil {
		errs = append(errs, err)
	}
	_, err = f.NewBytes(delimiter, ø.Js)
	if err != nil {
		errs = append(errs, err)
	}
	return
}

// checks if the dependancies are fullfilled and returns the errors
func (ø *Snack) DependancyErrors(libs []*JsLib) (errs []error) {
	errs = []error{}
	l := jslibs(libs)
	for name, version := range ø.Dependancies {
		if !l.Matches(name, version) {
			errs = append(errs, ŧ.Errorf(`Dependancy not fullfilled: %s (%s)`, name, version.String()))
		}
	}
	return
}

func (ø *Snack) Test(vals map[string][]byte, template string) string {
	errs := ø.DependancyErrors(ø.Libs)
	if len(errs) > 0 {
		s := []string{}
		for _, e := range errs {
			s = append(s, e.Error())
		}
		panic(strings.Join(s, "\n"))
	}
	i := ø.WithId("tester")
	return i.Test(vals, template)
}

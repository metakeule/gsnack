package gsnack

import (
	ŧ "fmt"
	"strings"
)

type Includable interface {
	Include() string
}

type JsLib struct {
	Path    string
	Name    string
	Version Version
}

func (ø *JsLib) Include() string {
	return ŧ.Sprintf(`<script src="%s"></script>`, ŧ.Sprintf(ø.Path, ø.Version.String()))
}

func Jquery(major uint, minor uint, minuscle uint) (ø *JsLib) {
	ø = &JsLib{}
	ø.Name = "jquery"
	ø.Path = `//ajax.googleapis.com/ajax/libs/jquery/%s/jquery.min.js`
	ø.Version = Version{major, minor, minuscle}
	return
}

type jslibs []*JsLib

func (ø jslibs) Matches(name string, version Version) bool {
	for _, l := range ø {
		if l.Name == name && version.Matches(l.Version) {
			return true
		}
	}
	return false
}

func (ø jslibs) Include() string {
	is := []string{}
	for _, l := range ø {
		is = append(is, l.Include())
	}
	return strings.Join(is, "\n")
}

package gsnack

import (
	ŧ "fmt"
)

type Version struct {
	Major, Minor, Patch uint
}

func (ø Version) String() string {
	return ŧ.Sprintf(`%v.%v.%v`, ø.Major, ø.Minor, ø.Patch)
}

func (ø Version) Matches(o Version) bool {
	if ø.Major != o.Major {
		return false
	}

	if ø.Minor > o.Minor {
		return false
	}

	if ø.Patch > o.Patch {
		return false
	}
	return true
}

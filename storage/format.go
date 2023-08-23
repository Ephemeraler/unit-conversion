package capacity

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var unitMap = map[string]float64{
	"b":  float64(B),
	"kb": float64(KB),
	"mb": float64(MB),
	"gb": float64(GB),
	"tb": float64(TB),
	"pb": float64(PB),
}

func ParseCapacity(s string) (*Capacity, error) {

	orig := s

	s = strings.ToLower(s)
	t, _ := regexp.Compile("^([0-9]*[.]?[0-9]+)([kmgtp]?[b]{1})$")
	if !t.MatchString(s) {
		return nil, errors.New("storage: invalid capacity " + orig)
	}

	tmp := t.FindStringSubmatch(s)
	value, _ := strconv.ParseFloat(tmp[1], 64)
	unit := tmp[2]

	return &Capacity{value: value * unitMap[unit]}, nil
}

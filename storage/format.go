package storage

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var unitMap = map[string]float64{
	"B":  float64(B),
	"KB": float64(KB),
	"MB": float64(MB),
	"GB": float64(GB),
	"TB": float64(TB),
	"PB": float64(PB),
}

// ParseCapacity parsers capacity's string to Capacity.
//
//	@param s the capacity's string.
//	@return *Capacity
//	@return error
func ParseCapacity(s string) (*Capacity, error) {

	orig := s

	s = strings.ToUpper(s)
	t, _ := regexp.Compile(`^([0-9]*[.]?[0-9]+)[\s]*([KMGTP]?[B]{1})$`)
	if !t.MatchString(s) {
		return nil, errors.New("storage: invalid capacity " + orig)
	}

	tmp := t.FindStringSubmatch(s)
	value, _ := strconv.ParseFloat(tmp[1], 64)
	unit := tmp[2]

	return &Capacity{value: value * unitMap[unit]}, nil
}

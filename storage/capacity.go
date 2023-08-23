package capacity

import (
	"fmt"
	"math"
)

const (
	B  float64 = 1
	KB         = 1024 * B
	MB         = 1024 * KB
	GB         = 1024 * MB
	TB         = 1024 * GB
	PB         = 1024 * TB
)

var unit = []string{"b", "kb", "mb", "gb", "tb", "pb"}

type Capacity struct {
	value float64 // default unit is byte
}

func (c *Capacity) String() string {
	u := math.Floor(math.Log2(c.value) / 10)
	v := c.value / unitMap[unit[int64(u)]]
	return fmt.Sprintf("%.2f%s", v, unit[int64(u)])
}

func (c *Capacity) ToByte() (float64, string) {
	return c.value, fmt.Sprintf("%.2fB", c.value)
}

func (c *Capacity) ToKByte() (float64, string) {
	v := c.value / KB
	return v, fmt.Sprintf("%.2fKB", v)
}

func (c *Capacity) ToMByte() (float64, string) {
	v := c.value / MB
	return v, fmt.Sprintf("%.2fMB", v)
}

func (c *Capacity) ToGByte() (float64, string) {
	v := c.value / GB
	return v, fmt.Sprintf("%.2fGB", v)
}

func (c *Capacity) ToTByte() (float64, string) {
	v := c.value / GB
	return v, fmt.Sprintf("%.2fTB", v)
}

func (c *Capacity) ToPByte() (float64, string) {
	v := c.value / PB
	return v, fmt.Sprintf("%.2fPB", v)
}

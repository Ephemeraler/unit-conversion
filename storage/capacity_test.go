package storage

import (
	"fmt"
	"testing"
)

func TestParseCapacity(t *testing.T) {
	tds := []string{
		"123.45 kB",
		"5.67MB",
		"1024 GB",
		"1.5tb",
		"500", // Missing unit.
		"2.0  b",
		"0.25 pb",
		"32.5gb",
		"8.75kb",
		"64  b",
		"0.5 b",
		"1000 pb",
		"0 b",
		"10 MB",
		"3.14159 gb",
		"1000000 TB",
		"15.5 ZB", // Invalid unit
		"4.5 qB",  // Invalid unit
		"1.2 m",   // Missing 'b' at the end
	}

	for _, td := range tds {
		_, err := ParseCapacity(td)
		if err != nil {
			fmt.Println(td, err)
			continue
		}
	}
}

func TestString(t *testing.T) {
	tds := []string{
		"123.45 kB",
		"5.67MB",
		"1024 GB",
		"1.5tb",
		"2.0  b",
		"0.25 pb",
		"32.5gb",
		"8.75kb",
		"64  b",
		"0.5 b",
		"1000 pb",
		"0 b",
		"10 MB",
		"3.14159 gb",
		"1000000 TB",
		"2048 pb",
	}

	for _, td := range tds {
		c, _ := ParseCapacity(td)
		fmt.Println(c)
	}
}

func TestTox(t *testing.T) {
	tds := []string{
		"123.45 kB",
		"5.67MB",
		"1024 GB",
		"1.5tb",
		"2.0  b",
		"0.25 pb",
		"32.5gb",
		"8.75kb",
		"64  b",
		"0.5 b",
		"1000 pb",
		"0 b",
		"10 MB",
		"3.14159 gb",
		"1000000 TB",
		"2048 pb",
	}

	for _, td := range tds {
		c, _ := ParseCapacity(td)
		fmt.Println(c.ToByte())
		fmt.Println(c.ToKByte())
		fmt.Println(c.ToMByte())
		fmt.Println(c.ToGByte())
		fmt.Println(c.ToTByte())
		fmt.Println(c.ToPByte())
	}
}

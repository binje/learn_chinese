package dictionary

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	m := parseHsk1()
	if len(m) != 497 {
		t.Fail()
		fmt.Println(len(m))
	}
	if m["也"] != "also" {
		t.Fail()
		fmt.Println(m["也"])
	}
}

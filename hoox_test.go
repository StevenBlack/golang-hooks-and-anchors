package hoox

import (
	"testing"
)

func TestHoox(t *testing.T) {
	a := Hook{name: "foo"}
	b := Hook{name: "bar"}
	c := Hook{name: "baz"}
	a.Sethook(b)
	a.Sethook(c)

	th := thing{t: "Hello"}
	out := a.Process(&th)

	result := out.t
	expect := "Hellofoobarbaz"
	if result != expect {
		t.Errorf("Expect result to be equal %s, but %s\n", expect, result)
	}

}

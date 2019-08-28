package hooks

import (
	"testing"
)

func (h *Hook) execute(th *thing) *thing {
	th.Sett(th.t + h.name)
	return th
}

func TestHoox(t *testing.T) {
	a := Hook{name: "foo"}
	b := Hook{name: "bar"}
	c := Hook{name: "baz"}
	a.Sethook(b)
	a.Sethook(c)

	th := thing{t: "Hello"}
	a.Process(&th)

	result := th.t
	expect := "Hellofoobarbaz"
	if result != expect {
		t.Errorf("Expect result to be equal %s, but %s\n", expect, result)
	}

}

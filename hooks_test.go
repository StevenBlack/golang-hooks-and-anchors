package hooks

import (
	"testing"
)

// implementing the hook
func (h *Hook) execute(th *thing) *thing {
	th.SetText(th.text + h.name)
	return th
}

func TestHoox(t *testing.T) {
	a := Hook{name: "foo"}
	b := Hook{name: "bar"}
	c := Hook{name: "baz"}
	a.Sethook(b)
	a.Sethook(c)

	th := thing{text: "Hello"}
	a.Process(&th)

	result := th.text
	expect := "Hellofoobarbaz"
	if result != expect {
		t.Errorf("Expect result to be equal %s, but %s\n", expect, result)
	}

}

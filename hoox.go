package main

import "fmt"

type thing struct {
	t   string
	i   int
	i8  int8
	i16 int16
	i32 int32
	i64 int64
	u   uint
	u8  uint8
	u16 uint16
	u32 uint32
	u64 uint64
}

func (t *thing) Sett(tstr string) {
	t.t = tstr
}

type hooker interface {
	Sethook(*Hook) bool
	Process(*thing) *thing
	preprocess(*thing)
	execute(*thing) *thing
	postprocess(*thing)
}

type Hook struct {
	name  string
	ohook *Hook
}

func (h *Hook) Setname(nm string) {
	h.name = nm
}

func (h *Hook) Sethook(hk Hook) bool {
	if h.ohook == nil {
		h.ohook = &hk
		return true
	} else {
		h.ohook.Sethook(hk)
		return true
	}
}

func (h *Hook) Process(th *thing) *thing {
	ret := h.execute(th)
	if h.ohook == nil {
		return ret
	} else {
		return h.ohook.Process(ret)
	}
}

func (h *Hook) execute(th *thing) *thing {
	th.Sett(th.t + h.name)
	return th
}

func main() {
	a := Hook{name: "foo"}
	b := Hook{name: "bar"}
	c := Hook{name: "baz"}
	a.Sethook(b)
	a.Sethook(c)

	t := thing{t: "Hello"}
	out := a.Process(&t)
	fmt.Print(out.t)
}

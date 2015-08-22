package gozwave

// #cgo LDFLAGS: -L./ -lfoo
// #include "foo.h"
import "C"

type GoFoo struct {
	foo C.Foo
}

func New() GoFoo {
	var ret GoFoo
	ret.foo = C.FooInit()
	return ret
}
func (f GoFoo) Free() {
	C.FooFree(f.foo)
}
func (f GoFoo) Bar() {
	C.FooBar(f.foo)
}

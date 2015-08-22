package gozwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "manager.h"
import "C"

type Manager struct {
	manager C.manager_t
}

func NewManager() Manager {
	var m Manager
	m.manager = C.manager_init()
	return m
}
func (m Manager) Free() {
	C.manager_free(m.manager)
}
func (m Manager) Bar() {
	C.manager_bar(m.manager)
}

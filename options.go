package gozwave

// #cgo LDFLAGS: -lopenzwave -L/usr/local/lib
// #cgo CPPFLAGS: -I/usr/local/include -I/usr/local/include/openzwave
// #include "options.h"
// #include <stdlib.h>
import "C"
import "unsafe"

// Options is a container for the C++ OpenZWave library Options class.
type Options struct {
	options C.options_t
}

// CreateOptions creates an object to manage the program options.
func CreateOptions(configPath, userPath, commandLine string) *Options {
	o := &Options{}
	cConfigPath := C.CString(configPath)
	cUserPath := C.CString(userPath)
	cCommandLine := C.CString(commandLine)
	o.options = C.options_create(cConfigPath, cUserPath, cCommandLine)
	C.free(unsafe.Pointer(cConfigPath))
	C.free(unsafe.Pointer(cUserPath))
	C.free(unsafe.Pointer(cCommandLine))
	return o
}

// DestroyOptions deletes the Options and cleans up any associated objects. The
// application is responsible for destroying the Options object, but this must
// not be done until after the Manager object has been destroyed.
func DestroyOptions() bool {
	if C.bool(C.options_destroy()) {
		return true
	}
	return false
}

// GetOptions gets a pointer to the Options singleton object.
func GetOptions() *Options {
	o := &Options{}
	o.options = C.options_get()
	return o
}

// Lock locks the options. Reads in option values from the XML options file and
// command line string and marks the options as locked. Once locked, no more
// calls to AddOption can be made. The options must be locked before the
// Manager::Create method is called.
func (o *Options) Lock() bool {
	if C.bool(C.options_lock(o.options)) {
		return true
	}
	return false
}

// AddOptionBool add a boolean option to the program. Adds an option to the
// program whose value can then be read from a file or command line. All calls
// to AddOptionInt must be made before Lock.
func (o *Options) AddOptionBool(name string, value bool) bool {
	cName := C.CString(name)
	result := C.bool(C.options_addOptionBool(o.options, cName, C.bool(value)))
	C.free(unsafe.Pointer(cName))
	if result {
		return true
	}
	return false
}

// AddOptionInt add an integer option to the program. Adds an option to the
// program whose value can then be read from a file or command line. All calls
// to AddOptionInt must be made before Lock.
func (o *Options) AddOptionInt(name string, value int32) bool {
	cName := C.CString(name)
	result := C.bool(C.options_addOptionInt(o.options, cName, C.int32_t(value)))
	C.free(unsafe.Pointer(cName))
	if result {
		return true
	}
	return false
}

// AddOptionString add a string option to the program. Adds an option to the
// program whose value can then be read from a file or command line. All calls
// to AddOptionString must be made before Lock.
func (o *Options) AddOptionString(name string, value string, append bool) bool {
	cName := C.CString(name)
	cValue := C.CString(value)
	result := C.bool(C.options_addOptionString(o.options, cName, cValue, C.bool(append)))
	C.free(unsafe.Pointer(cName))
	C.free(unsafe.Pointer(cValue))
	if result {
		return true
	}
	return false
}

// GetOptionAsBool get the value of a boolean option.
func (o *Options) GetOptionAsBool(name string) (valueOut bool, ok bool) {
	cName := C.CString(name)
	var cValue *C.bool
	result := C.bool(C.options_getOptionAsBool(o.options, cName, cValue))
	C.free(unsafe.Pointer(cName))
	if *cValue {
		valueOut = true
	} else {
		valueOut = false
	}
	if result {
		ok = true
	} else {
		ok = false
	}
	return valueOut, ok
}

// GetOptionAsInt get the value of an integer option.
func (o *Options) GetOptionAsInt(name string) (valueOut int, ok bool) {
	cName := C.CString(name)
	var cValue *C.int32_t
	result := C.bool(C.options_getOptionAsInt(o.options, cName, cValue))
	C.free(unsafe.Pointer(cName))
	valueOut = int(*cValue)
	if result {
		ok = true
	} else {
		ok = false
	}
	return valueOut, ok
}

// GetOptionAsString get the value of a string option.
// func (o *Options) GetOptionAsString(name string) (valueOut int, ok bool) {
// 	cName := C.CString(name)
// 	var cValue *C.int
// 	ok = C.bool(C.options_getOptionAsInt(o.options, cName, cValue))
// 	C.free(unsafe.Pointer(cName))
// 	valueOut = C.GoInt(cValue)
// 	return
// }

// GetOptionType get the type of value stored in an option.
//////// OptionType options_getOptionType(string const &_name);

// AreLocked test whether the options have been locked.
func (o *Options) AreLocked() bool {
	if C.bool(C.options_areLocked(o.options)) {
		return true
	}
	return false
}

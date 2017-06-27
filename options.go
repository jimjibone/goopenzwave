package goopenzwave

// #include "gzw_options.h"
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
	return bool(C.options_lock(o.options))
}

// AddOptionBool add a boolean option to the program. Adds an option to the
// program whose value can then be read from a file or command line. All calls
// to AddOptionInt must be made before Lock.
func (o *Options) AddOptionBool(name string, value bool) bool {
	cName := C.CString(name)
	result := bool(C.options_addOptionBool(o.options, cName, C.bool(value)))
	C.free(unsafe.Pointer(cName))
	return result
}

// AddOptionInt add an integer option to the program. Adds an option to the
// program whose value can then be read from a file or command line. All calls
// to AddOptionInt must be made before Lock.
func (o *Options) AddOptionInt(name string, value int32) bool {
	cName := C.CString(name)
	result := bool(C.options_addOptionInt(o.options, cName, C.int32_t(value)))
	C.free(unsafe.Pointer(cName))
	return result
}

// AddOptionLogLevel add a log level option to the program. Adds an option to
// the program whose value can then be read from a file or command line. All
// calls to AddOptionLogLevel must be made before Lock.
func (o *Options) AddOptionLogLevel(name string, value LogLevel) bool {
	cName := C.CString(name)
	result := bool(C.options_addOptionLogLevel(o.options, cName, C.loglevel_t(value)))
	C.free(unsafe.Pointer(cName))
	return result
}

// AddOptionString add a string option to the program. Adds an option to the
// program whose value can then be read from a file or command line. All calls
// to AddOptionString must be made before Lock.
func (o *Options) AddOptionString(name string, value string, append bool) bool {
	cName := C.CString(name)
	cValue := C.CString(value)
	result := bool(C.options_addOptionString(o.options, cName, cValue, C.bool(append)))
	C.free(unsafe.Pointer(cName))
	C.free(unsafe.Pointer(cValue))
	return result
}

// GetOptionAsBool get the value of a boolean option.
func (o *Options) GetOptionAsBool(name string) (bool, bool) {
	cName := C.CString(name)
	var cValue C.bool
	result := bool(C.options_getOptionAsBool(o.options, cName, &cValue))
	C.free(unsafe.Pointer(cName))
	return result, bool(cValue)
}

// GetOptionAsInt get the value of an integer option.
func (o *Options) GetOptionAsInt(name string) (bool, int32) {
	cName := C.CString(name)
	var cValue C.int32_t
	result := bool(C.options_getOptionAsInt(o.options, cName, &cValue))
	C.free(unsafe.Pointer(cName))
	return result, int32(cValue)
}

// GetOptionAsString get the value of a string option.
func (o *Options) GetOptionAsString(name string) (bool, string) {
	cName := C.CString(name)
	var cstr *C.char
	result := bool(C.options_getOptionAsString(o.options, cName, &cstr))
	gostr := C.GoString(cstr)
	C.free(unsafe.Pointer(cstr))
	return result, gostr
}

// GetOptionType get the type of value stored in an option.
//TODO OptionType options_getOptionType(string const &_name);

// AreLocked test whether the options have been locked.
func (o *Options) AreLocked() bool {
	return bool(C.options_areLocked(o.options))
}

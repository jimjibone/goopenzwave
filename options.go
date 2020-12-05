package goopenzwave

// #cgo pkg-config: libopenzwave
// #include "options_wrap.h"
// #include <stdlib.h>
import "C"
import "unsafe"

// OptionsCreate creates the openzwave program options object.
// configPath is the path to the OpenZWave library config folder, which contains XML descriptions of Z-Wave manufacturers and products.
// userPath is the path to the application's user data folder where the OpenZWave should store the Z-Wave network configuration and state.
// commandLine is the program's command line options.
// Wraps `Options* OpenZWave::Options::Create(...)`.
func OptionsCreate(configPath, userPath, commandLine string) {
	// Options::Create( "../../../config/", "", "" );
	cConfigPath := C.CString(configPath)
	cUserPath := C.CString(userPath)
	cCommandLine := C.CString(commandLine)
	defer C.free(unsafe.Pointer(cConfigPath))
	defer C.free(unsafe.Pointer(cUserPath))
	defer C.free(unsafe.Pointer(cCommandLine))
	C.options_create(cConfigPath, cUserPath, cCommandLine)
}

// OptionsAddBool adds a boolean option to the program. Must be called before OptionsLock.
// name is the name of the option. Option names are case insensitive and must be unique.
// defaultval is the default value for this option.
// Wraps `bool OpenZWave::Options::AddOptionBool(...)`.
func OptionsAddBool(name string, defaultval bool) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	_ = C.options_add_bool(cname, C.bool(defaultval))
}

// OptionsAddInt adds an integer option to the program. Must be called before OptionsLock.
// name is the name of the option. Option names are case insensitive and must be unique.
// defaultval is the default value for this option.
// Wraps `bool OpenZWave::Options::AddOptionInt(...)`.
func OptionsAddInt(name string, defaultval int32) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	_ = C.options_add_int(cname, C.int32_t(defaultval))
}

// OptionsAddString adds a string option to the program. Must be called before OptionsLock.
// name is the name of the option. Option names are case insensitive and must be unique.
// defaultval is the default value for this option.
// append, if set to true, will cause values read from the command line or XML file to be concatenated into a comma delimited list. If false, newer values will overwrite older ones
// Wraps `bool OpenZWave::Options::AddOptionString(...)`.
func OptionsAddString(name string, defaultval string, append bool) {
	cname := C.CString(name)
	cdefaultval := C.CString(defaultval)
	defer C.free(unsafe.Pointer(cname))
	defer C.free(unsafe.Pointer(cdefaultval))
	_ = C.options_add_string(cname, cdefaultval, C.bool(append))
}

// OptionsLock reads in option values from the XML options file and command line string and marks the options as locked. Once locked, no more calls to OptionsAdd* can be made. The options must be locked before the ManagerCreate function is called.
// Wraps `bool OpenZWave::Options::Lock()`.
func OptionsLock() {
	_ = C.options_lock()
}

// OptionsDestroy deletes the Options and cleans up any associated objects. The application is responsible for destroying the Options object, but this must not be done until after the Manager object has been destroyed.
// Wraps `bool OpenZWave::Options::Destroy()`.
func OptionsDestroy() {
	_ = C.options_destroy()
}

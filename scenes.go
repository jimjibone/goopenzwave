package goopenzwave

// #include "gzw_manager.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// GetNumScenes returns the number of scenes that have been defined.
func GetNumScenes() uint8 {
	return uint8(C.manager_getNumScenes(cmanager))
}

// GetAllScenes Gets a list of all the SceneIds.
//TODO(jimjibone) func GetAllScenes(...) ...

// RemoveAllScenes removes all the SceneIds.
func RemoveAllScenes(homeID uint32) {
	C.manager_removeAllScenes(cmanager, C.uint32_t(homeID))
}

// CreateScene creates a new Scene and returns the scene ID.
func CreateScene() uint8 {
	return uint8(C.manager_createScene(cmanager))
}

// RemoveScene removes an existing Scene. Returns true if the scene was removed.
func RemoveScene(sceneID uint8) bool {
	return bool(C.manager_removeScene(cmanager, C.uint8_t(sceneID)))
}

// AddSceneValueBool adds a bool Value ID to an existing scene. Returns true if
// the Value ID was added.
func AddSceneValueBool(sceneID uint8, homeID uint32, valueID uint64, value bool) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_addSceneValueBool(cmanager, C.uint8_t(sceneID), cvalueid, C.bool(value)))
}

// AddSceneValueUint8 adds a bool Value ID to an existing scene. Returns true if
// the Value ID was added.
func AddSceneValueUint8(sceneID uint8, homeID uint32, valueID uint64, value uint8) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_addSceneValueUint8(cmanager, C.uint8_t(sceneID), cvalueid, C.uint8_t(value)))
}

// AddSceneValueFloat adds a decimal Value ID to an existing scene. Returns true
// if the Value ID was added.
func AddSceneValueFloat(sceneID uint8, homeID uint32, valueID uint64, value float32) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_addSceneValueFloat(cmanager, C.uint8_t(sceneID), cvalueid, C.float(value)))
}

// AddSceneValueInt32 adds a 32-bit signed integer Value ID to an existing
// scene. Returns true if the Value ID was added.
func AddSceneValueInt32(sceneID uint8, homeID uint32, valueID uint64, value int32) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_addSceneValueInt32(cmanager, C.uint8_t(sceneID), cvalueid, C.int32_t(value)))
}

// AddSceneValueInt16 adds a 16-bit signed integer Value ID to an existing
// scene. Returns true if the Value ID was added.
func AddSceneValueInt16(sceneID uint8, homeID uint32, valueID uint64, value int16) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_addSceneValueInt16(cmanager, C.uint8_t(sceneID), cvalueid, C.int16_t(value)))
}

// AddSceneValueString adds a string Value ID to an existing scene. Returns true
// if the Value ID was added.
func AddSceneValueString(sceneID uint8, homeID uint32, valueID uint64, value string) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.CString(value)
	result := bool(C.manager_addSceneValueString(cmanager, C.uint8_t(sceneID), cvalueid, cstring))
	C.free(unsafe.Pointer(cstring))
	return result
}

// AddSceneValueListSelectionString adds the selected item list Value ID to an
// existing scene (as a string). Returns true if the Value ID was added.
func AddSceneValueListSelectionString(sceneID uint8, homeID uint32, valueID uint64, value string) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.CString(value)
	result := bool(C.manager_addSceneValueListSelectionString(cmanager, C.uint8_t(sceneID), cvalueid, cstring))
	C.free(unsafe.Pointer(cstring))
	return result
}

// AddSceneValueListSelectionInt32 adds the selected item list Value ID to an
// existing scene (as a integer). Returns true if the Value ID was added.
func AddSceneValueListSelectionInt32(sceneID uint8, homeID uint32, valueID uint64, value int32) bool {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	return bool(C.manager_addSceneValueListSelectionInt32(cmanager, C.uint8_t(sceneID), cvalueid, C.int32_t(value)))
}

// RemoveSceneValue removes the Value ID from an existing scene.
// TODO: bool RemoveSceneValue (uint8 const _sceneId, ValueID const &_valueId) ...

// SceneGetValues retrieves the scene's list of values.
// TODO: int SceneGetValues (uint8 const _sceneId, vector< ValueID > *o_value) ...

// GetSceneValueAsBool returns a scene's value as a bool and returns an error if
// the value was not obtained.
func GetSceneValueAsBool(sceneID uint8, homeID uint32, valueID uint64) (bool, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cbool C.bool
	ok := bool(C.manager_sceneGetValueAsBool(cmanager, C.uint8_t(sceneID), cvalueid, &cbool))
	if ok == false {
		return bool(cbool), fmt.Errorf("bool value was not obtained")
	}
	return bool(cbool), nil
}

// GetSceneValueAsByte returns a scene's value as a byte and returns an error if
// the value was not obtained.
func GetSceneValueAsByte(sceneID uint8, homeID uint32, valueID uint64) (byte, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cbyte C.uint8_t
	ok := bool(C.manager_sceneGetValueAsByte(cmanager, C.uint8_t(sceneID), cvalueid, &cbyte))
	if ok == false {
		return byte(cbyte), fmt.Errorf("byte value was not obtained")
	}
	return byte(cbyte), nil
}

// GetSceneValueAsFloat returns a scene's value as a float and returns an error
// if the value was not obtained.
func GetSceneValueAsFloat(sceneID uint8, homeID uint32, valueID uint64) (float32, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cfloat C.float
	ok := bool(C.manager_sceneGetValueAsFloat(cmanager, C.uint8_t(sceneID), cvalueid, &cfloat))
	if ok == false {
		return float32(cfloat), fmt.Errorf("float value was not obtained")
	}
	return float32(cfloat), nil
}

// GetSceneValueAsInt returns a scene's value as a 32-bit signed integer and
// returns an error if the value was not obtained.
func GetSceneValueAsInt(sceneID uint8, homeID uint32, valueID uint64) (int32, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cint C.int32_t
	ok := bool(C.manager_sceneGetValueAsInt(cmanager, C.uint8_t(sceneID), cvalueid, &cint))
	if ok == false {
		return int32(cint), fmt.Errorf("int value was not obtained")
	}
	return int32(cint), nil
}

// GetSceneValueAsShort returns a scene's value as a 16-bit signed integer and
// returns an error if the value was not obtained.
func GetSceneValueAsShort(sceneID uint8, homeID uint32, valueID uint64) (int16, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cshort C.int16_t
	ok := bool(C.manager_sceneGetValueAsShort(cmanager, C.uint8_t(sceneID), cvalueid, &cshort))
	if ok == false {
		return int16(cshort), fmt.Errorf("short value was not obtained")
	}
	return int16(cshort), nil
}

// GetSceneValueAsString returns a scene's value as a string and returns an
// error if the value was not obtained.
func GetSceneValueAsString(sceneID uint8, homeID uint32, valueID uint64) (string, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cstr *C.char
	ok := bool(C.manager_sceneGetValueAsString(cmanager, C.uint8_t(sceneID), cvalueid, &cstr))
	gostr := C.GoString(cstr)
	C.free(unsafe.Pointer(cstr))
	if ok == false {
		return gostr, fmt.Errorf("string value was not obtained")
	}
	return gostr, nil
}

// GetSceneValueListSelectionString returns a scene's value list as a string and
// returns an error if the value was not obtained.
func GetSceneValueListSelectionString(sceneID uint8, homeID uint32, valueID uint64) (string, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cstr *C.char
	ok := bool(C.manager_sceneGetValueListSelectionString(cmanager, C.uint8_t(sceneID), cvalueid, &cstr))
	gostr := C.GoString(cstr)
	C.free(unsafe.Pointer(cstr))
	if ok == false {
		return gostr, fmt.Errorf("string list value was not obtained")
	}
	return gostr, nil
}

// GetSceneValueListSelectionInt32 returns a scene's value list as an integer
// and returns an error if the value was not obtained.
func GetSceneValueListSelectionInt32(sceneID uint8, homeID uint32, valueID uint64) (int32, error) {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	var cint C.int32_t
	ok := bool(C.manager_sceneGetValueListSelectionInt32(cmanager, C.uint8_t(sceneID), cvalueid, &cint))
	if ok == false {
		return int32(cint), fmt.Errorf("int list value was not obtained")
	}
	return int32(cint), nil
}

// SetSceneValueBool sets a bool Value ID to an existing scene's ValueID.
// Returns an error if the Value ID was not added.
func SetSceneValueBool(sceneID uint8, homeID uint32, valueID uint64, value bool) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setSceneValueBool(cmanager, C.uint8_t(sceneID), cvalueid, C.bool(value)))
	if ok == false {
		return fmt.Errorf("bool value was not added to scene")
	}
	return nil
}

// SetSceneValueUint8 sets a byte Value ID to an existing scene's ValueID.
// Returns an error if the Value ID was not added.
func SetSceneValueUint8(sceneID uint8, homeID uint32, valueID uint64, value uint8) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setSceneValueUint8(cmanager, C.uint8_t(sceneID), cvalueid, C.uint8_t(value)))
	if ok == false {
		return fmt.Errorf("byte value was not added to scene")
	}
	return nil
}

// SetSceneValueFloat sets a decimal Value ID to an existing scene's ValueID.
// Returns an error if the Value ID was not added.
func SetSceneValueFloat(sceneID uint8, homeID uint32, valueID uint64, value float32) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setSceneValueFloat(cmanager, C.uint8_t(sceneID), cvalueid, C.float(value)))
	if ok == false {
		return fmt.Errorf("float value was not added to scene")
	}
	return nil
}

// SetSceneValueInt32 sets a 32-bit signed integer Value ID to an existing
// scene's ValueID. Returns an error if the Value ID was not added.
func SetSceneValueInt32(sceneID uint8, homeID uint32, valueID uint64, value int32) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setSceneValueInt32(cmanager, C.uint8_t(sceneID), cvalueid, C.int32_t(value)))
	if ok == false {
		return fmt.Errorf("32-bit signed integer value was not added to scene")
	}
	return nil
}

// SetSceneValueInt16 sets a 16-bit integer Value ID to an existing scene's
// ValueID. Returns an error if the Value ID was not added.
func SetSceneValueInt16(sceneID uint8, homeID uint32, valueID uint64, value int16) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setSceneValueInt16(cmanager, C.uint8_t(sceneID), cvalueid, C.int16_t(value)))
	if ok == false {
		return fmt.Errorf("16-bit integer value was not added to scene")
	}
	return nil
}

// SetSceneValueString sets a string Value ID to an existing scene's ValueID.
// Returns an error if the Value ID was not added.
func SetSceneValueString(sceneID uint8, homeID uint32, valueID uint64, value string) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.CString(value)
	ok := bool(C.manager_setSceneValueString(cmanager, C.uint8_t(sceneID), cvalueid, cstring))
	C.free(unsafe.Pointer(cstring))
	if ok == false {
		return fmt.Errorf("string value was not added to scene")
	}
	return nil
}

// SetSceneValueListSelectionString sets the list selected item Value ID to an
// existing scene's ValueID (as a string). Returns an error if the Value ID was
// not added.
func SetSceneValueListSelectionString(sceneID uint8, homeID uint32, valueID uint64, value string) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	cstring := C.CString(value)
	ok := bool(C.manager_setSceneValueListSelectionString(cmanager, C.uint8_t(sceneID), cvalueid, cstring))
	C.free(unsafe.Pointer(cstring))
	if ok == false {
		return fmt.Errorf("string value list selection was not added to scene")
	}
	return nil
}

// SetSceneValueListSelectionInt32 sets the list selected item Value ID to an
// existing scene's ValueID (as a integer). Returns an error if the Value ID was
// not added.
func SetSceneValueListSelectionInt32(sceneID uint8, homeID uint32, valueID uint64, value int32) error {
	cvalueid := C.valueid_create(C.uint32_t(homeID), C.uint64_t(valueID))
	defer C.valueid_free(cvalueid)
	ok := bool(C.manager_setSceneValueListSelectionInt32(cmanager, C.uint8_t(sceneID), cvalueid, C.int32_t(value)))
	if ok == false {
		return fmt.Errorf("int value list selection was not added to scene")
	}
	return nil
}

// GetSceneLabel returns a label for the particular scene.
func GetSceneLabel(sceneID uint8) string {
	cstr := C.manager_getSceneLabel(cmanager, C.uint8_t(sceneID))
	gostr := C.GoString(cstr)
	C.free(unsafe.Pointer(cstr))
	return gostr
}

// SetSceneLabel sets a label for the particular scene.
func SetSceneLabel(sceneID uint8, value string) {
	cstring := C.CString(value)
	C.manager_setSceneLabel(cmanager, C.uint8_t(sceneID), cstring)
	C.free(unsafe.Pointer(cstring))
}

// SceneExists returns true if a Scene ID is defined.
func SceneExists(sceneID uint8) bool {
	return bool(C.manager_sceneExists(cmanager, C.uint8_t(sceneID)))
}

// ActivateScene activates a given scene to perform all its actions. Returns an
// error if the scene was not activated.
func ActivateScene(sceneID uint8) error {
	ok := bool(C.manager_activateScene(cmanager, C.uint8_t(sceneID)))
	if ok == false {
		return fmt.Errorf("failed to activate scene")
	}
	return nil
}

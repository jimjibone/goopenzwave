package goopenzwave

// #include "manager.h"
// #include <stdlib.h>
import "C"

// GetNumGroups returns the number of association groups reported by this node.
//
// In Z-Wave, groups are numbered starting from one. For example, if a call to
// GetNumGroups returns 4, the _groupIdx value to use in calls to
// GetAssociations, AddAssociation and RemoveAssociation will be a number
// between 1 and 4.
func GetNumGroups(homeID uint32, nodeID uint8) uint8 {
	return uint8(C.manager_getNumGroups(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID)))
}

// GetAssociations returns the associations for a group.
//
// Makes a copy of the list of associated nodes in the group, and returns it in
// an array of uint8's. The caller is responsible for freeing the array memory
// with a call to delete [].
//TODO func GetAssociations(homeID uint32, nodeID uint8, groupIDx uint8, uint8 **o_associations) ...

// GetAssociations returns the associations for a group.
//
// Makes a copy of the list of associated nodes in the group, and returns it in
// an array of InstanceAssociation's. The caller is responsible for freeing the
// array memory with a call to delete [].
//TODO func GetAssociations(homeID uint32, nodeID uint8, groupIDx uint8, InstanceAssociation **o_associations) ...

// GetMaxAssociations returns the maximum number of associations for a group.
func GetMaxAssociations(homeID uint32, nodeID uint8, groupIDx uint8) uint8 {
	return uint8(C.manager_getMaxAssociations(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(groupIDx)))
}

// GetGroupLabel returns a label for the particular group of a node. This label
// is populated by the device specific configuration files.
func GetGroupLabel(homeID uint32, nodeID uint8, groupIDx uint8) string {
	cString := C.manager_getGroupLabel(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(groupIDx))
	goString := C.GoString(cString.data)
	C.string_freeString(cString)
	return goString
}

// AddAssociation adds a node to an association group.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the association data held in this class is updated directly.
// This will be reverted by a future Association message from the device if the
// Z-Wave message actually failed to get through. Notification callbacks will be
// sent in both cases.
func AddAssociation(homeID uint32, nodeID uint8, groupIDx uint8, targetNodeID uint8, instance uint8) {
	C.manager_addAssociation(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(groupIDx), C.uint8_t(targetNodeID), C.uint8_t(instance))
}

// RemoveAssociation removes a node from an association group.
//
// Due to the possibility of a device being asleep, the command is assumed to
// succeed, and the association data held in this class is updated directly.
// This will be reverted by a future Association message from the device if the
// Z-Wave message actually failed to get through. Notification callbacks will be
// sent in both cases.
func RemoveAssociation(homeID uint32, nodeID uint8, groupIDx uint8, targetNodeID uint8, instance uint8) {
	C.manager_removeAssociation(cmanager, C.uint32_t(homeID), C.uint8_t(nodeID), C.uint8_t(groupIDx), C.uint8_t(targetNodeID), C.uint8_t(instance))
}

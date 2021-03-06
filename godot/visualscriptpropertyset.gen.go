package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewVisualScriptPropertySetFromPointer(ptr gdnative.Pointer) VisualScriptPropertySet {
func newVisualScriptPropertySetFromPointer(ptr gdnative.Pointer) VisualScriptPropertySet {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptPropertySet{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptPropertySet struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptPropertySet) BaseClass() string {
	return "VisualScriptPropertySet"
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/
func (o *VisualScriptPropertySet) X_GetTypeCache() gdnative.Dictionary {
	//log.Println("Calling VisualScriptPropertySet.X_GetTypeCache()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "_get_type_cache")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false type_cache Dictionary}], Returns: void
*/
func (o *VisualScriptPropertySet) X_SetTypeCache(typeCache gdnative.Dictionary) {
	//log.Println("Calling VisualScriptPropertySet.X_SetTypeCache()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(typeCache)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "_set_type_cache")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: enum.VisualScriptPropertySet::AssignOp
*/

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *VisualScriptPropertySet) GetBasePath() gdnative.NodePath {
	//log.Println("Calling VisualScriptPropertySet.GetBasePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "get_base_path")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptPropertySet) GetBaseScript() gdnative.String {
	//log.Println("Calling VisualScriptPropertySet.GetBaseScript()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "get_base_script")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptPropertySet) GetBaseType() gdnative.String {
	//log.Println("Calling VisualScriptPropertySet.GetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "get_base_type")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.Variant::Type
*/

/*
        Undocumented
	Args: [], Returns: enum.VisualScriptPropertySet::CallMode
*/

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptPropertySet) GetIndex() gdnative.String {
	//log.Println("Calling VisualScriptPropertySet.GetIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "get_index")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptPropertySet) GetProperty() gdnative.String {
	//log.Println("Calling VisualScriptPropertySet.GetProperty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "get_property")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false assign_op int}], Returns: void
*/
func (o *VisualScriptPropertySet) SetAssignOp(assignOp gdnative.Int) {
	//log.Println("Calling VisualScriptPropertySet.SetAssignOp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(assignOp)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "set_assign_op")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false base_path NodePath}], Returns: void
*/
func (o *VisualScriptPropertySet) SetBasePath(basePath gdnative.NodePath) {
	//log.Println("Calling VisualScriptPropertySet.SetBasePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(basePath)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "set_base_path")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false base_script String}], Returns: void
*/
func (o *VisualScriptPropertySet) SetBaseScript(baseScript gdnative.String) {
	//log.Println("Calling VisualScriptPropertySet.SetBaseScript()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(baseScript)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "set_base_script")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false base_type String}], Returns: void
*/
func (o *VisualScriptPropertySet) SetBaseType(baseType gdnative.String) {
	//log.Println("Calling VisualScriptPropertySet.SetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(baseType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "set_base_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false basic_type int}], Returns: void
*/
func (o *VisualScriptPropertySet) SetBasicType(basicType gdnative.Int) {
	//log.Println("Calling VisualScriptPropertySet.SetBasicType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(basicType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "set_basic_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *VisualScriptPropertySet) SetCallMode(mode gdnative.Int) {
	//log.Println("Calling VisualScriptPropertySet.SetCallMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "set_call_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false index String}], Returns: void
*/
func (o *VisualScriptPropertySet) SetIndex(index gdnative.String) {
	//log.Println("Calling VisualScriptPropertySet.SetIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "set_index")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false property String}], Returns: void
*/
func (o *VisualScriptPropertySet) SetProperty(property gdnative.String) {
	//log.Println("Calling VisualScriptPropertySet.SetProperty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(property)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertySet", "set_property")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptPropertySetImplementer is an interface that implements the methods
// of the VisualScriptPropertySet class.
type VisualScriptPropertySetImplementer interface {
	VisualScriptNodeImplementer
	X_GetTypeCache() gdnative.Dictionary
	X_SetTypeCache(typeCache gdnative.Dictionary)
	GetBasePath() gdnative.NodePath
	GetBaseScript() gdnative.String
	GetBaseType() gdnative.String
	GetIndex() gdnative.String
	GetProperty() gdnative.String
	SetAssignOp(assignOp gdnative.Int)
	SetBasePath(basePath gdnative.NodePath)
	SetBaseScript(baseScript gdnative.String)
	SetBaseType(baseType gdnative.String)
	SetBasicType(basicType gdnative.Int)
	SetCallMode(mode gdnative.Int)
	SetIndex(index gdnative.String)
	SetProperty(property gdnative.String)
}

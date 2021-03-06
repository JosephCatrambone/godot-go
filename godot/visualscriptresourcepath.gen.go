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

//func NewVisualScriptResourcePathFromPointer(ptr gdnative.Pointer) VisualScriptResourcePath {
func newVisualScriptResourcePathFromPointer(ptr gdnative.Pointer) VisualScriptResourcePath {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptResourcePath{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptResourcePath struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptResourcePath) BaseClass() string {
	return "VisualScriptResourcePath"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptResourcePath) GetResourcePath() gdnative.String {
	//log.Println("Calling VisualScriptResourcePath.GetResourcePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptResourcePath", "get_resource_path")

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
	Args: [{ false path String}], Returns: void
*/
func (o *VisualScriptResourcePath) SetResourcePath(path gdnative.String) {
	//log.Println("Calling VisualScriptResourcePath.SetResourcePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptResourcePath", "set_resource_path")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptResourcePathImplementer is an interface that implements the methods
// of the VisualScriptResourcePath class.
type VisualScriptResourcePathImplementer interface {
	VisualScriptNodeImplementer
	GetResourcePath() gdnative.String
	SetResourcePath(path gdnative.String)
}

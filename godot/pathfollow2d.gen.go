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

//func NewPathFollow2DFromPointer(ptr gdnative.Pointer) PathFollow2D {
func newPathFollow2DFromPointer(ptr gdnative.Pointer) PathFollow2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PathFollow2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This node takes its parent [Path2D], and returns the coordinates of a point within it, given a distance from the first vertex. It is useful for making other nodes follow a path, without coding the movement pattern. For that, the nodes must be descendants of this node. Then, when setting an offset in this node, the descendant nodes will move accordingly.
*/
type PathFollow2D struct {
	Node2D
	owner gdnative.Object
}

func (o *PathFollow2D) BaseClass() string {
	return "PathFollow2D"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *PathFollow2D) GetCubicInterpolation() gdnative.Bool {
	//log.Println("Calling PathFollow2D.GetCubicInterpolation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "get_cubic_interpolation")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *PathFollow2D) GetHOffset() gdnative.Float {
	//log.Println("Calling PathFollow2D.GetHOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "get_h_offset")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *PathFollow2D) GetLookahead() gdnative.Float {
	//log.Println("Calling PathFollow2D.GetLookahead()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "get_lookahead")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *PathFollow2D) GetOffset() gdnative.Float {
	//log.Println("Calling PathFollow2D.GetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "get_offset")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *PathFollow2D) GetUnitOffset() gdnative.Float {
	//log.Println("Calling PathFollow2D.GetUnitOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "get_unit_offset")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *PathFollow2D) GetVOffset() gdnative.Float {
	//log.Println("Calling PathFollow2D.GetVOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "get_v_offset")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *PathFollow2D) HasLoop() gdnative.Bool {
	//log.Println("Calling PathFollow2D.HasLoop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "has_loop")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *PathFollow2D) IsRotating() gdnative.Bool {
	//log.Println("Calling PathFollow2D.IsRotating()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "is_rotating")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *PathFollow2D) SetCubicInterpolation(enable gdnative.Bool) {
	//log.Println("Calling PathFollow2D.SetCubicInterpolation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "set_cubic_interpolation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false h_offset float}], Returns: void
*/
func (o *PathFollow2D) SetHOffset(hOffset gdnative.Float) {
	//log.Println("Calling PathFollow2D.SetHOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(hOffset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "set_h_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false lookahead float}], Returns: void
*/
func (o *PathFollow2D) SetLookahead(lookahead gdnative.Float) {
	//log.Println("Calling PathFollow2D.SetLookahead()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(lookahead)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "set_lookahead")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false loop bool}], Returns: void
*/
func (o *PathFollow2D) SetLoop(loop gdnative.Bool) {
	//log.Println("Calling PathFollow2D.SetLoop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(loop)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "set_loop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false offset float}], Returns: void
*/
func (o *PathFollow2D) SetOffset(offset gdnative.Float) {
	//log.Println("Calling PathFollow2D.SetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "set_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *PathFollow2D) SetRotate(enable gdnative.Bool) {
	//log.Println("Calling PathFollow2D.SetRotate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "set_rotate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false unit_offset float}], Returns: void
*/
func (o *PathFollow2D) SetUnitOffset(unitOffset gdnative.Float) {
	//log.Println("Calling PathFollow2D.SetUnitOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(unitOffset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "set_unit_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false v_offset float}], Returns: void
*/
func (o *PathFollow2D) SetVOffset(vOffset gdnative.Float) {
	//log.Println("Calling PathFollow2D.SetVOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(vOffset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow2D", "set_v_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// PathFollow2DImplementer is an interface that implements the methods
// of the PathFollow2D class.
type PathFollow2DImplementer interface {
	Node2DImplementer
	GetCubicInterpolation() gdnative.Bool
	GetHOffset() gdnative.Float
	GetLookahead() gdnative.Float
	GetOffset() gdnative.Float
	GetUnitOffset() gdnative.Float
	GetVOffset() gdnative.Float
	HasLoop() gdnative.Bool
	IsRotating() gdnative.Bool
	SetCubicInterpolation(enable gdnative.Bool)
	SetHOffset(hOffset gdnative.Float)
	SetLookahead(lookahead gdnative.Float)
	SetLoop(loop gdnative.Bool)
	SetOffset(offset gdnative.Float)
	SetRotate(enable gdnative.Bool)
	SetUnitOffset(unitOffset gdnative.Float)
	SetVOffset(vOffset gdnative.Float)
}

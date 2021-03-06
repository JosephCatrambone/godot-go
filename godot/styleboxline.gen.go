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

//func NewStyleBoxLineFromPointer(ptr gdnative.Pointer) StyleBoxLine {
func newStyleBoxLineFromPointer(ptr gdnative.Pointer) StyleBoxLine {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := StyleBoxLine{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type StyleBoxLine struct {
	StyleBox
	owner gdnative.Object
}

func (o *StyleBoxLine) BaseClass() string {
	return "StyleBoxLine"
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *StyleBoxLine) GetColor() gdnative.Color {
	//log.Println("Calling StyleBoxLine.GetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxLine", "get_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *StyleBoxLine) GetGrow() gdnative.Float {
	//log.Println("Calling StyleBoxLine.GetGrow()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxLine", "get_grow")

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
	Args: [], Returns: int
*/
func (o *StyleBoxLine) GetThickness() gdnative.Int {
	//log.Println("Calling StyleBoxLine.GetThickness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxLine", "get_thickness")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *StyleBoxLine) IsVertical() gdnative.Bool {
	//log.Println("Calling StyleBoxLine.IsVertical()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxLine", "is_vertical")

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
	Args: [{ false color Color}], Returns: void
*/
func (o *StyleBoxLine) SetColor(color gdnative.Color) {
	//log.Println("Calling StyleBoxLine.SetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxLine", "set_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false grow float}], Returns: void
*/
func (o *StyleBoxLine) SetGrow(grow gdnative.Float) {
	//log.Println("Calling StyleBoxLine.SetGrow()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(grow)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxLine", "set_grow")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false thickness int}], Returns: void
*/
func (o *StyleBoxLine) SetThickness(thickness gdnative.Int) {
	//log.Println("Calling StyleBoxLine.SetThickness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(thickness)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxLine", "set_thickness")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false vertical bool}], Returns: void
*/
func (o *StyleBoxLine) SetVertical(vertical gdnative.Bool) {
	//log.Println("Calling StyleBoxLine.SetVertical()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(vertical)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxLine", "set_vertical")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// StyleBoxLineImplementer is an interface that implements the methods
// of the StyleBoxLine class.
type StyleBoxLineImplementer interface {
	StyleBoxImplementer
	GetColor() gdnative.Color
	GetGrow() gdnative.Float
	GetThickness() gdnative.Int
	IsVertical() gdnative.Bool
	SetColor(color gdnative.Color)
	SetGrow(grow gdnative.Float)
	SetThickness(thickness gdnative.Int)
	SetVertical(vertical gdnative.Bool)
}

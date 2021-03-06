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

//func NewPath2DFromPointer(ptr gdnative.Pointer) Path2D {
func newPath2DFromPointer(ptr gdnative.Pointer) Path2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Path2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Can have [PathFollow2D] child-nodes moving along the [Curve2D]. See [PathFollow2D] for more information on this usage.
*/
type Path2D struct {
	Node2D
	owner gdnative.Object
}

func (o *Path2D) BaseClass() string {
	return "Path2D"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Path2D) X_CurveChanged() {
	//log.Println("Calling Path2D.X_CurveChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Path2D", "_curve_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Curve2D
*/
func (o *Path2D) GetCurve() Curve2DImplementer {
	//log.Println("Calling Path2D.GetCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Path2D", "get_curve")

	// Call the parent method.
	// Curve2D
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newCurve2DFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(Curve2DImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false curve Curve2D}], Returns: void
*/
func (o *Path2D) SetCurve(curve Curve2D) {
	//log.Println("Calling Path2D.SetCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(curve.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Path2D", "set_curve")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// Path2DImplementer is an interface that implements the methods
// of the Path2D class.
type Path2DImplementer interface {
	Node2DImplementer
	X_CurveChanged()
	GetCurve() Curve2DImplementer
	SetCurve(curve Curve2D)
}

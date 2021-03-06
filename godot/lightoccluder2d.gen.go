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

//func NewLightOccluder2DFromPointer(ptr gdnative.Pointer) LightOccluder2D {
func newLightOccluder2DFromPointer(ptr gdnative.Pointer) LightOccluder2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := LightOccluder2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Occludes light cast by a Light2D, casting shadows. The LightOccluder2D must be provided with an [OccluderPolygon2D] in order for the shadow to be computed.
*/
type LightOccluder2D struct {
	Node2D
	owner gdnative.Object
}

func (o *LightOccluder2D) BaseClass() string {
	return "LightOccluder2D"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *LightOccluder2D) X_PolyChanged() {
	//log.Println("Calling LightOccluder2D.X_PolyChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LightOccluder2D", "_poly_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *LightOccluder2D) GetOccluderLightMask() gdnative.Int {
	//log.Println("Calling LightOccluder2D.GetOccluderLightMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LightOccluder2D", "get_occluder_light_mask")

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
	Args: [], Returns: OccluderPolygon2D
*/
func (o *LightOccluder2D) GetOccluderPolygon() OccluderPolygon2DImplementer {
	//log.Println("Calling LightOccluder2D.GetOccluderPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LightOccluder2D", "get_occluder_polygon")

	// Call the parent method.
	// OccluderPolygon2D
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newOccluderPolygon2DFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(OccluderPolygon2DImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *LightOccluder2D) SetOccluderLightMask(mask gdnative.Int) {
	//log.Println("Calling LightOccluder2D.SetOccluderLightMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LightOccluder2D", "set_occluder_light_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false polygon OccluderPolygon2D}], Returns: void
*/
func (o *LightOccluder2D) SetOccluderPolygon(polygon OccluderPolygon2D) {
	//log.Println("Calling LightOccluder2D.SetOccluderPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(polygon.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LightOccluder2D", "set_occluder_polygon")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// LightOccluder2DImplementer is an interface that implements the methods
// of the LightOccluder2D class.
type LightOccluder2DImplementer interface {
	Node2DImplementer
	X_PolyChanged()
	GetOccluderLightMask() gdnative.Int
	GetOccluderPolygon() OccluderPolygon2DImplementer
	SetOccluderLightMask(mask gdnative.Int)
	SetOccluderPolygon(polygon OccluderPolygon2D)
}

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

//func NewStaticBodyFromPointer(ptr gdnative.Pointer) StaticBody {
func newStaticBodyFromPointer(ptr gdnative.Pointer) StaticBody {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := StaticBody{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Static body for 3D Physics. A static body is a simple body that is not intended to move. They don't consume any CPU resources in contrast to a [RigidBody3D] so they are great for scenario collision. A static body can also be animated by using simulated motion mode. This is useful for implementing functionalities such as moving platforms. When this mode is active the body can be animated and automatically computes linear and angular velocity to apply in that frame and to influence other bodies. Alternatively, a constant linear or angular velocity can be set for the static body, so even if it doesn't move, it affects other bodies as if it was moving (this is useful for simulating conveyor belts or conveyor wheels).
*/
type StaticBody struct {
	PhysicsBody
	owner gdnative.Object
}

func (o *StaticBody) BaseClass() string {
	return "StaticBody"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *StaticBody) GetBounce() gdnative.Float {
	//log.Println("Calling StaticBody.GetBounce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StaticBody", "get_bounce")

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
	Args: [], Returns: Vector3
*/
func (o *StaticBody) GetConstantAngularVelocity() gdnative.Vector3 {
	//log.Println("Calling StaticBody.GetConstantAngularVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StaticBody", "get_constant_angular_velocity")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *StaticBody) GetConstantLinearVelocity() gdnative.Vector3 {
	//log.Println("Calling StaticBody.GetConstantLinearVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StaticBody", "get_constant_linear_velocity")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *StaticBody) GetFriction() gdnative.Float {
	//log.Println("Calling StaticBody.GetFriction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StaticBody", "get_friction")

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
	Args: [{ false bounce float}], Returns: void
*/
func (o *StaticBody) SetBounce(bounce gdnative.Float) {
	//log.Println("Calling StaticBody.SetBounce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(bounce)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StaticBody", "set_bounce")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false vel Vector3}], Returns: void
*/
func (o *StaticBody) SetConstantAngularVelocity(vel gdnative.Vector3) {
	//log.Println("Calling StaticBody.SetConstantAngularVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(vel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StaticBody", "set_constant_angular_velocity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false vel Vector3}], Returns: void
*/
func (o *StaticBody) SetConstantLinearVelocity(vel gdnative.Vector3) {
	//log.Println("Calling StaticBody.SetConstantLinearVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(vel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StaticBody", "set_constant_linear_velocity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false friction float}], Returns: void
*/
func (o *StaticBody) SetFriction(friction gdnative.Float) {
	//log.Println("Calling StaticBody.SetFriction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(friction)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StaticBody", "set_friction")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// StaticBodyImplementer is an interface that implements the methods
// of the StaticBody class.
type StaticBodyImplementer interface {
	PhysicsBodyImplementer
	GetBounce() gdnative.Float
	GetConstantAngularVelocity() gdnative.Vector3
	GetConstantLinearVelocity() gdnative.Vector3
	GetFriction() gdnative.Float
	SetBounce(bounce gdnative.Float)
	SetConstantAngularVelocity(vel gdnative.Vector3)
	SetConstantLinearVelocity(vel gdnative.Vector3)
	SetFriction(friction gdnative.Float)
}

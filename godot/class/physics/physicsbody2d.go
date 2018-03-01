package physics

import (
	"log"
	"reflect"

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

/*
PhysicsBody2D is an abstract base class for implementing a physics body. All *Body2D types inherit from it.
*/
type PhysicsBody2D struct {
	CollisionObject2D
}

func (o *PhysicsBody2D) BaseClass() string {
	return "PhysicsBody2D"
}

/*
   Undocumented
*/
func (o *PhysicsBody2D) X_GetLayers() gdnative.Int {
	log.Println("Calling PhysicsBody2D.X_GetLayers()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "_get_layers", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PhysicsBody2D) X_SetLayers(mask gdnative.Int) {
	log.Println("Calling PhysicsBody2D.X_SetLayers()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_set_layers", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a body to the list of bodies that this body can't collide with.
*/
func (o *PhysicsBody2D) AddCollisionExceptionWith(body *Object) {
	log.Println("Calling PhysicsBody2D.AddCollisionExceptionWith()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(body)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_collision_exception_with", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PhysicsBody2D) GetCollisionLayer() gdnative.Int {
	log.Println("Calling PhysicsBody2D.GetCollisionLayer()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_layer", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return an individual bit on the collision mask.
*/
func (o *PhysicsBody2D) GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool {
	log.Println("Calling PhysicsBody2D.GetCollisionLayerBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bit)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_layer_bit", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *PhysicsBody2D) GetCollisionMask() gdnative.Int {
	log.Println("Calling PhysicsBody2D.GetCollisionMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_mask", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return an individual bit on the collision mask.
*/
func (o *PhysicsBody2D) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	log.Println("Calling PhysicsBody2D.GetCollisionMaskBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bit)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_mask_bit", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Removes a body from the list of bodies that this body can't collide with.
*/
func (o *PhysicsBody2D) RemoveCollisionExceptionWith(body *Object) {
	log.Println("Calling PhysicsBody2D.RemoveCollisionExceptionWith()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(body)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "remove_collision_exception_with", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PhysicsBody2D) SetCollisionLayer(layer gdnative.Int) {
	log.Println("Calling PhysicsBody2D.SetCollisionLayer()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(layer)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_layer", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set/clear individual bits on the layer mask. This makes getting a body in/out of only one layer easier.
*/
func (o *PhysicsBody2D) SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool) {
	log.Println("Calling PhysicsBody2D.SetCollisionLayerBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(bit)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_layer_bit", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *PhysicsBody2D) SetCollisionMask(mask gdnative.Int) {
	log.Println("Calling PhysicsBody2D.SetCollisionMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set/clear individual bits on the collision mask. This makes selecting the areas scanned easier.
*/
func (o *PhysicsBody2D) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	log.Println("Calling PhysicsBody2D.SetCollisionMaskBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(bit)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_mask_bit", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   PhysicsBody2DImplementer is an interface for PhysicsBody2D objects.
*/
type PhysicsBody2DImplementer interface {
	Class
}
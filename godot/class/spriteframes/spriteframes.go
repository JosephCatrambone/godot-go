package spriteframes

import (
	"log"
	"reflect"
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
Sprite frame library for [AnimatedSprite]. Contains frames and animation data for playback.
*/
type SpriteFrames struct {
	Resource
}

func (o *SpriteFrames) BaseClass() string {
	return "SpriteFrames"
}

/*
   Undocumented
*/
func (o *SpriteFrames) X_GetAnimations() *Array {
	log.Println("Calling SpriteFrames.X_GetAnimations()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "_get_animations", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *SpriteFrames) X_GetFrames() *Array {
	log.Println("Calling SpriteFrames.X_GetFrames()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "_get_frames", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *SpriteFrames) X_SetAnimations(arg0 *Array) {
	log.Println("Calling SpriteFrames.X_SetAnimations()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_set_animations", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *SpriteFrames) X_SetFrames(arg0 *Array) {
	log.Println("Calling SpriteFrames.X_SetFrames()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_set_frames", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a new animation to the the library.
*/
func (o *SpriteFrames) AddAnimation(anim gdnative.String) {
	log.Println("Calling SpriteFrames.AddAnimation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(anim)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_animation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a frame to the given animation.
*/
func (o *SpriteFrames) AddFrame(anim gdnative.String, frame *Texture, atPosition gdnative.Int) {
	log.Println("Calling SpriteFrames.AddFrame()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(anim)
	goArguments[1] = reflect.ValueOf(frame)
	goArguments[2] = reflect.ValueOf(atPosition)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_frame", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes all frames from the given animation.
*/
func (o *SpriteFrames) Clear(anim gdnative.String) {
	log.Println("Calling SpriteFrames.Clear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(anim)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "clear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes all animations. A "default" animation will be created.
*/
func (o *SpriteFrames) ClearAll() {
	log.Println("Calling SpriteFrames.ClearAll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "clear_all", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If [code]true[/code] the given animation will loop.
*/
func (o *SpriteFrames) GetAnimationLoop(anim gdnative.String) gdnative.Bool {
	log.Println("Calling SpriteFrames.GetAnimationLoop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(anim)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_animation_loop", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   The animation's speed in frames per second.
*/
func (o *SpriteFrames) GetAnimationSpeed(anim gdnative.String) gdnative.Float {
	log.Println("Calling SpriteFrames.GetAnimationSpeed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(anim)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_animation_speed", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the animation's selected frame.
*/
func (o *SpriteFrames) GetFrame(anim gdnative.String, idx gdnative.Int) *Texture {
	log.Println("Calling SpriteFrames.GetFrame()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(anim)
	goArguments[1] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_frame", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the number of frames in the animation.
*/
func (o *SpriteFrames) GetFrameCount(anim gdnative.String) gdnative.Int {
	log.Println("Calling SpriteFrames.GetFrameCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(anim)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_frame_count", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   If [code]true[/code] the named animation exists.
*/
func (o *SpriteFrames) HasAnimation(anim gdnative.String) gdnative.Bool {
	log.Println("Calling SpriteFrames.HasAnimation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(anim)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "has_animation", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Removes the given animation.
*/
func (o *SpriteFrames) RemoveAnimation(anim gdnative.String) {
	log.Println("Calling SpriteFrames.RemoveAnimation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(anim)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "remove_animation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes the animation's selected frame.
*/
func (o *SpriteFrames) RemoveFrame(anim gdnative.String, idx gdnative.Int) {
	log.Println("Calling SpriteFrames.RemoveFrame()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(anim)
	goArguments[1] = reflect.ValueOf(idx)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "remove_frame", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Changes the animation's name to [code]newname[/code].
*/
func (o *SpriteFrames) RenameAnimation(anim gdnative.String, newname gdnative.String) {
	log.Println("Calling SpriteFrames.RenameAnimation()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(anim)
	goArguments[1] = reflect.ValueOf(newname)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "rename_animation", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   If [code]true[/code] the animation will loop.
*/
func (o *SpriteFrames) SetAnimationLoop(anim gdnative.String, loop gdnative.Bool) {
	log.Println("Calling SpriteFrames.SetAnimationLoop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(anim)
	goArguments[1] = reflect.ValueOf(loop)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_animation_loop", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   The animation's speed in frames per second.
*/
func (o *SpriteFrames) SetAnimationSpeed(anim gdnative.String, speed gdnative.Float) {
	log.Println("Calling SpriteFrames.SetAnimationSpeed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(anim)
	goArguments[1] = reflect.ValueOf(speed)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_animation_speed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the texture of the given frame.
*/
func (o *SpriteFrames) SetFrame(anim gdnative.String, idx gdnative.Int, txt *Texture) {
	log.Println("Calling SpriteFrames.SetFrame()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(anim)
	goArguments[1] = reflect.ValueOf(idx)
	goArguments[2] = reflect.ValueOf(txt)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_frame", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   SpriteFramesImplementer is an interface for SpriteFrames objects.
*/
type SpriteFramesImplementer interface {
	Class
}
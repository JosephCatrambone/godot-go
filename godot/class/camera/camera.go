package camera

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
Camera is a special node that displays what is visible from its current location. Cameras register themselves in the nearest [Viewport] node (when ascending the tree). Only one camera can be active per viewport. If no viewport is available ascending the tree, the Camera will register in the global viewport. In other words, a Camera just provides [i]3D[/i] display capabilities to a [Viewport], and, without one, a scene registered in that [Viewport] (or higher viewports) can't be displayed.
*/
type Camera struct {
	Spatial
}

func (o *Camera) BaseClass() string {
	return "Camera"
}

/*
   If this is the current Camera, remove it from being current. If it is inside the node tree, request to make the next Camera current, if any.
*/
func (o *Camera) ClearCurrent() {
	log.Println("Calling Camera.ClearCurrent()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "clear_current", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Get the camera transform. Subclassed cameras (such as CharacterCamera) may provide different transforms than the [Node] transform.
*/
func (o *Camera) GetCameraTransform() *Transform {
	log.Println("Calling Camera.GetCameraTransform()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_camera_transform", goArguments, "*Transform")

	returnValue := goRet.Interface().(*Transform)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) GetCullMask() gdnative.Int {
	log.Println("Calling Camera.GetCullMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_cull_mask", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) GetDopplerTracking() gdnative.Int {
	log.Println("Calling Camera.GetDopplerTracking()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_doppler_tracking", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) GetEnvironment() *Environment {
	log.Println("Calling Camera.GetEnvironment()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_environment", goArguments, "*Environment")

	returnValue := goRet.Interface().(*Environment)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) GetFov() gdnative.Float {
	log.Println("Calling Camera.GetFov()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_fov", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) GetHOffset() gdnative.Float {
	log.Println("Calling Camera.GetHOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_h_offset", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) GetKeepAspectMode() gdnative.Int {
	log.Println("Calling Camera.GetKeepAspectMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_keep_aspect_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) GetProjection() gdnative.Int {
	log.Println("Calling Camera.GetProjection()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_projection", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) GetSize() gdnative.Float {
	log.Println("Calling Camera.GetSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_size", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) GetVOffset() gdnative.Float {
	log.Println("Calling Camera.GetVOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_v_offset", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) GetZfar() gdnative.Float {
	log.Println("Calling Camera.GetZfar()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_zfar", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) GetZnear() gdnative.Float {
	log.Println("Calling Camera.GetZnear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_znear", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) IsCurrent() gdnative.Bool {
	log.Println("Calling Camera.IsCurrent()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_current", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the given position is behind the Camera.
*/
func (o *Camera) IsPositionBehind(worldPoint *Vector3) gdnative.Bool {
	log.Println("Calling Camera.IsPositionBehind()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(worldPoint)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_position_behind", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Make this camera the current Camera for the [Viewport] (see class description). If the Camera Node is outside the scene tree, it will attempt to become current once it's added.
*/
func (o *Camera) MakeCurrent() {
	log.Println("Calling Camera.MakeCurrent()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "make_current", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns a normal vector from the screen point location directed along the camera. Orthogonal cameras are normalized. Perspective cameras account for perspective, screen width/height, etc.
*/
func (o *Camera) ProjectLocalRayNormal(screenPoint *Vector2) *Vector3 {
	log.Println("Calling Camera.ProjectLocalRayNormal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(screenPoint)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "project_local_ray_normal", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns how a 2D coordinate in the Viewport rectangle maps to a 3D point in worldspace.
*/
func (o *Camera) ProjectPosition(screenPoint *Vector2) *Vector3 {
	log.Println("Calling Camera.ProjectPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(screenPoint)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "project_position", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a normal vector in worldspace, that is the result of projecting a point on the [Viewport] rectangle by the camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
func (o *Camera) ProjectRayNormal(screenPoint *Vector2) *Vector3 {
	log.Println("Calling Camera.ProjectRayNormal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(screenPoint)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "project_ray_normal", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a 3D position in worldspace, that is the result of projecting a point on the [Viewport] rectangle by the camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
func (o *Camera) ProjectRayOrigin(screenPoint *Vector2) *Vector3 {
	log.Println("Calling Camera.ProjectRayOrigin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(screenPoint)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "project_ray_origin", goArguments, "*Vector3")

	returnValue := goRet.Interface().(*Vector3)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Camera) SetCullMask(mask gdnative.Int) {
	log.Println("Calling Camera.SetCullMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_cull_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Camera) SetCurrent(arg0 gdnative.Bool) {
	log.Println("Calling Camera.SetCurrent()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_current", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Camera) SetDopplerTracking(mode gdnative.Int) {
	log.Println("Calling Camera.SetDopplerTracking()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_doppler_tracking", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Camera) SetEnvironment(env *Environment) {
	log.Println("Calling Camera.SetEnvironment()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(env)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_environment", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Camera) SetFov(arg0 gdnative.Float) {
	log.Println("Calling Camera.SetFov()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_fov", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Camera) SetHOffset(ofs gdnative.Float) {
	log.Println("Calling Camera.SetHOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(ofs)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_h_offset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Camera) SetKeepAspectMode(mode gdnative.Int) {
	log.Println("Calling Camera.SetKeepAspectMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_keep_aspect_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the camera projection to orthogonal mode, by specifying a width and the [i]near[/i] and [i]far[/i] clip planes in worldspace units. (As a hint, 2D games often use this projection, with values specified in pixels)
*/
func (o *Camera) SetOrthogonal(size gdnative.Float, zNear gdnative.Float, zFar gdnative.Float) {
	log.Println("Calling Camera.SetOrthogonal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(size)
	goArguments[1] = reflect.ValueOf(zNear)
	goArguments[2] = reflect.ValueOf(zFar)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_orthogonal", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set the camera projection to perspective mode, by specifying a [i]FOV[/i] Y angle in degrees (FOV means Field of View), and the [i]near[/i] and [i]far[/i] clip planes in worldspace units.
*/
func (o *Camera) SetPerspective(fov gdnative.Float, zNear gdnative.Float, zFar gdnative.Float) {
	log.Println("Calling Camera.SetPerspective()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(fov)
	goArguments[1] = reflect.ValueOf(zNear)
	goArguments[2] = reflect.ValueOf(zFar)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_perspective", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Camera) SetProjection(arg0 gdnative.Int) {
	log.Println("Calling Camera.SetProjection()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_projection", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Camera) SetSize(arg0 gdnative.Float) {
	log.Println("Calling Camera.SetSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Camera) SetVOffset(ofs gdnative.Float) {
	log.Println("Calling Camera.SetVOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(ofs)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_v_offset", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Camera) SetZfar(arg0 gdnative.Float) {
	log.Println("Calling Camera.SetZfar()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_zfar", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Camera) SetZnear(arg0 gdnative.Float) {
	log.Println("Calling Camera.SetZnear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_znear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns how a 3D point in worldspace maps to a 2D coordinate in the [Viewport] rectangle.
*/
func (o *Camera) UnprojectPosition(worldPoint *Vector3) *Vector2 {
	log.Println("Calling Camera.UnprojectPosition()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(worldPoint)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "unproject_position", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   CameraImplementer is an interface for Camera objects.
*/
type CameraImplementer interface {
	Class
}
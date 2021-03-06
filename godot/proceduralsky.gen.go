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

//func NewProceduralSkyFromPointer(ptr gdnative.Pointer) ProceduralSky {
func newProceduralSkyFromPointer(ptr gdnative.Pointer) ProceduralSky {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ProceduralSky{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type ProceduralSky struct {
	Sky
	owner gdnative.Object
}

func (o *ProceduralSky) BaseClass() string {
	return "ProceduralSky"
}

/*
        Undocumented
	Args: [{ false image Image}], Returns: void
*/
func (o *ProceduralSky) X_ThreadDone(image Image) {
	//log.Println("Calling ProceduralSky.X_ThreadDone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(image.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "_thread_done")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ProceduralSky) X_UpdateSky() {
	//log.Println("Calling ProceduralSky.X_UpdateSky()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "_update_sky")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *ProceduralSky) GetGroundBottomColor() gdnative.Color {
	//log.Println("Calling ProceduralSky.GetGroundBottomColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_ground_bottom_color")

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
func (o *ProceduralSky) GetGroundCurve() gdnative.Float {
	//log.Println("Calling ProceduralSky.GetGroundCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_ground_curve")

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
func (o *ProceduralSky) GetGroundEnergy() gdnative.Float {
	//log.Println("Calling ProceduralSky.GetGroundEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_ground_energy")

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
	Args: [], Returns: Color
*/
func (o *ProceduralSky) GetGroundHorizonColor() gdnative.Color {
	//log.Println("Calling ProceduralSky.GetGroundHorizonColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_ground_horizon_color")

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
func (o *ProceduralSky) GetSkyCurve() gdnative.Float {
	//log.Println("Calling ProceduralSky.GetSkyCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sky_curve")

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
func (o *ProceduralSky) GetSkyEnergy() gdnative.Float {
	//log.Println("Calling ProceduralSky.GetSkyEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sky_energy")

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
	Args: [], Returns: Color
*/
func (o *ProceduralSky) GetSkyHorizonColor() gdnative.Color {
	//log.Println("Calling ProceduralSky.GetSkyHorizonColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sky_horizon_color")

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
	Args: [], Returns: Color
*/
func (o *ProceduralSky) GetSkyTopColor() gdnative.Color {
	//log.Println("Calling ProceduralSky.GetSkyTopColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sky_top_color")

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
func (o *ProceduralSky) GetSunAngleMax() gdnative.Float {
	//log.Println("Calling ProceduralSky.GetSunAngleMax()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_angle_max")

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
func (o *ProceduralSky) GetSunAngleMin() gdnative.Float {
	//log.Println("Calling ProceduralSky.GetSunAngleMin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_angle_min")

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
	Args: [], Returns: Color
*/
func (o *ProceduralSky) GetSunColor() gdnative.Color {
	//log.Println("Calling ProceduralSky.GetSunColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_color")

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
func (o *ProceduralSky) GetSunCurve() gdnative.Float {
	//log.Println("Calling ProceduralSky.GetSunCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_curve")

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
func (o *ProceduralSky) GetSunEnergy() gdnative.Float {
	//log.Println("Calling ProceduralSky.GetSunEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_energy")

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
func (o *ProceduralSky) GetSunLatitude() gdnative.Float {
	//log.Println("Calling ProceduralSky.GetSunLatitude()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_latitude")

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
func (o *ProceduralSky) GetSunLongitude() gdnative.Float {
	//log.Println("Calling ProceduralSky.GetSunLongitude()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "get_sun_longitude")

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
	Args: [], Returns: enum.ProceduralSky::TextureSize
*/

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ProceduralSky) SetGroundBottomColor(color gdnative.Color) {
	//log.Println("Calling ProceduralSky.SetGroundBottomColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_ground_bottom_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false curve float}], Returns: void
*/
func (o *ProceduralSky) SetGroundCurve(curve gdnative.Float) {
	//log.Println("Calling ProceduralSky.SetGroundCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(curve)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_ground_curve")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false energy float}], Returns: void
*/
func (o *ProceduralSky) SetGroundEnergy(energy gdnative.Float) {
	//log.Println("Calling ProceduralSky.SetGroundEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(energy)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_ground_energy")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ProceduralSky) SetGroundHorizonColor(color gdnative.Color) {
	//log.Println("Calling ProceduralSky.SetGroundHorizonColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_ground_horizon_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false curve float}], Returns: void
*/
func (o *ProceduralSky) SetSkyCurve(curve gdnative.Float) {
	//log.Println("Calling ProceduralSky.SetSkyCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(curve)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sky_curve")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false energy float}], Returns: void
*/
func (o *ProceduralSky) SetSkyEnergy(energy gdnative.Float) {
	//log.Println("Calling ProceduralSky.SetSkyEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(energy)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sky_energy")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ProceduralSky) SetSkyHorizonColor(color gdnative.Color) {
	//log.Println("Calling ProceduralSky.SetSkyHorizonColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sky_horizon_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ProceduralSky) SetSkyTopColor(color gdnative.Color) {
	//log.Println("Calling ProceduralSky.SetSkyTopColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sky_top_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *ProceduralSky) SetSunAngleMax(degrees gdnative.Float) {
	//log.Println("Calling ProceduralSky.SetSunAngleMax()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_angle_max")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *ProceduralSky) SetSunAngleMin(degrees gdnative.Float) {
	//log.Println("Calling ProceduralSky.SetSunAngleMin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_angle_min")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ProceduralSky) SetSunColor(color gdnative.Color) {
	//log.Println("Calling ProceduralSky.SetSunColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false curve float}], Returns: void
*/
func (o *ProceduralSky) SetSunCurve(curve gdnative.Float) {
	//log.Println("Calling ProceduralSky.SetSunCurve()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(curve)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_curve")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false energy float}], Returns: void
*/
func (o *ProceduralSky) SetSunEnergy(energy gdnative.Float) {
	//log.Println("Calling ProceduralSky.SetSunEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(energy)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_energy")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *ProceduralSky) SetSunLatitude(degrees gdnative.Float) {
	//log.Println("Calling ProceduralSky.SetSunLatitude()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_latitude")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *ProceduralSky) SetSunLongitude(degrees gdnative.Float) {
	//log.Println("Calling ProceduralSky.SetSunLongitude()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_sun_longitude")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false size int}], Returns: void
*/
func (o *ProceduralSky) SetTextureSize(size gdnative.Int) {
	//log.Println("Calling ProceduralSky.SetTextureSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ProceduralSky", "set_texture_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ProceduralSkyImplementer is an interface that implements the methods
// of the ProceduralSky class.
type ProceduralSkyImplementer interface {
	SkyImplementer
	X_ThreadDone(image Image)
	X_UpdateSky()
	GetGroundBottomColor() gdnative.Color
	GetGroundCurve() gdnative.Float
	GetGroundEnergy() gdnative.Float
	GetGroundHorizonColor() gdnative.Color
	GetSkyCurve() gdnative.Float
	GetSkyEnergy() gdnative.Float
	GetSkyHorizonColor() gdnative.Color
	GetSkyTopColor() gdnative.Color
	GetSunAngleMax() gdnative.Float
	GetSunAngleMin() gdnative.Float
	GetSunColor() gdnative.Color
	GetSunCurve() gdnative.Float
	GetSunEnergy() gdnative.Float
	GetSunLatitude() gdnative.Float
	GetSunLongitude() gdnative.Float
	SetGroundBottomColor(color gdnative.Color)
	SetGroundCurve(curve gdnative.Float)
	SetGroundEnergy(energy gdnative.Float)
	SetGroundHorizonColor(color gdnative.Color)
	SetSkyCurve(curve gdnative.Float)
	SetSkyEnergy(energy gdnative.Float)
	SetSkyHorizonColor(color gdnative.Color)
	SetSkyTopColor(color gdnative.Color)
	SetSunAngleMax(degrees gdnative.Float)
	SetSunAngleMin(degrees gdnative.Float)
	SetSunColor(color gdnative.Color)
	SetSunCurve(curve gdnative.Float)
	SetSunEnergy(energy gdnative.Float)
	SetSunLatitude(degrees gdnative.Float)
	SetSunLongitude(degrees gdnative.Float)
	SetTextureSize(size gdnative.Int)
}

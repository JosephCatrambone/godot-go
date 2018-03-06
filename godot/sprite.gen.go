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

//func NewSpriteFromPointer(ptr gdnative.Pointer) Sprite {
func newSpriteFromPointer(ptr gdnative.Pointer) Sprite {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Sprite{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A node that displays a 2D texture. The texture displayed can be a region from a larger atlas texture, or a frame from a sprite sheet animation.
*/
type Sprite struct {
	Node2D
	owner gdnative.Object
}

func (o *Sprite) BaseClass() string {
	return "Sprite"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Sprite) GetFrame() gdnative.Int {
	//log.Println("Calling Sprite.GetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "get_frame")

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
	Args: [], Returns: int
*/
func (o *Sprite) GetHframes() gdnative.Int {
	//log.Println("Calling Sprite.GetHframes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "get_hframes")

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
	Args: [], Returns: Texture
*/
func (o *Sprite) GetNormalMap() TextureImplementer {
	//log.Println("Calling Sprite.GetNormalMap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "get_normal_map")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *Sprite) GetOffset() gdnative.Vector2 {
	//log.Println("Calling Sprite.GetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "get_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Rect2
*/
func (o *Sprite) GetRegionRect() gdnative.Rect2 {
	//log.Println("Calling Sprite.GetRegionRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "get_region_rect")

	// Call the parent method.
	// Rect2
	retPtr := gdnative.NewEmptyRect2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRect2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *Sprite) GetTexture() TextureImplementer {
	//log.Println("Calling Sprite.GetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "get_texture")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Sprite) GetVframes() gdnative.Int {
	//log.Println("Calling Sprite.GetVframes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "get_vframes")

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
func (o *Sprite) IsCentered() gdnative.Bool {
	//log.Println("Calling Sprite.IsCentered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "is_centered")

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
func (o *Sprite) IsFlippedH() gdnative.Bool {
	//log.Println("Calling Sprite.IsFlippedH()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "is_flipped_h")

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
func (o *Sprite) IsFlippedV() gdnative.Bool {
	//log.Println("Calling Sprite.IsFlippedV()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "is_flipped_v")

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
func (o *Sprite) IsRegion() gdnative.Bool {
	//log.Println("Calling Sprite.IsRegion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "is_region")

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
func (o *Sprite) IsRegionFilterClipEnabled() gdnative.Bool {
	//log.Println("Calling Sprite.IsRegionFilterClipEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "is_region_filter_clip_enabled")

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
	Args: [{ false centered bool}], Returns: void
*/
func (o *Sprite) SetCentered(centered gdnative.Bool) {
	//log.Println("Calling Sprite.SetCentered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(centered)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_centered")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flip_h bool}], Returns: void
*/
func (o *Sprite) SetFlipH(flipH gdnative.Bool) {
	//log.Println("Calling Sprite.SetFlipH()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(flipH)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_flip_h")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flip_v bool}], Returns: void
*/
func (o *Sprite) SetFlipV(flipV gdnative.Bool) {
	//log.Println("Calling Sprite.SetFlipV()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(flipV)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_flip_v")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false frame int}], Returns: void
*/
func (o *Sprite) SetFrame(frame gdnative.Int) {
	//log.Println("Calling Sprite.SetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(frame)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_frame")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false hframes int}], Returns: void
*/
func (o *Sprite) SetHframes(hframes gdnative.Int) {
	//log.Println("Calling Sprite.SetHframes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(hframes)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_hframes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false normal_map Texture}], Returns: void
*/
func (o *Sprite) SetNormalMap(normalMap Texture) {
	//log.Println("Calling Sprite.SetNormalMap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(normalMap.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_normal_map")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false offset Vector2}], Returns: void
*/
func (o *Sprite) SetOffset(offset gdnative.Vector2) {
	//log.Println("Calling Sprite.SetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Sprite) SetRegion(enabled gdnative.Bool) {
	//log.Println("Calling Sprite.SetRegion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_region")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Sprite) SetRegionFilterClip(enabled gdnative.Bool) {
	//log.Println("Calling Sprite.SetRegionFilterClip()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_region_filter_clip")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rect Rect2}], Returns: void
*/
func (o *Sprite) SetRegionRect(rect gdnative.Rect2) {
	//log.Println("Calling Sprite.SetRegionRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRect2(rect)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_region_rect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *Sprite) SetTexture(texture Texture) {
	//log.Println("Calling Sprite.SetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false vframes int}], Returns: void
*/
func (o *Sprite) SetVframes(vframes gdnative.Int) {
	//log.Println("Calling Sprite.SetVframes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(vframes)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite", "set_vframes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// SpriteImplementer is an interface that implements the methods
// of the Sprite class.
type SpriteImplementer interface {
	Node2DImplementer
	GetFrame() gdnative.Int
	GetHframes() gdnative.Int
	GetNormalMap() TextureImplementer
	GetOffset() gdnative.Vector2
	GetRegionRect() gdnative.Rect2
	GetTexture() TextureImplementer
	GetVframes() gdnative.Int
	IsCentered() gdnative.Bool
	IsFlippedH() gdnative.Bool
	IsFlippedV() gdnative.Bool
	IsRegion() gdnative.Bool
	IsRegionFilterClipEnabled() gdnative.Bool
	SetCentered(centered gdnative.Bool)
	SetFlipH(flipH gdnative.Bool)
	SetFlipV(flipV gdnative.Bool)
	SetFrame(frame gdnative.Int)
	SetHframes(hframes gdnative.Int)
	SetNormalMap(normalMap Texture)
	SetOffset(offset gdnative.Vector2)
	SetRegion(enabled gdnative.Bool)
	SetRegionFilterClip(enabled gdnative.Bool)
	SetRegionRect(rect gdnative.Rect2)
	SetTexture(texture Texture)
	SetVframes(vframes gdnative.Int)
}
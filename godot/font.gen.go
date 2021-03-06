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

//func NewFontFromPointer(ptr gdnative.Pointer) Font {
func newFontFromPointer(ptr gdnative.Pointer) Font {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Font{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Font contains a unicode compatible character set, as well as the ability to draw it with variable width, ascent, descent and kerning. For creating fonts from TTF files (or other font formats), see the editor support for fonts. TODO check wikipedia for graph of ascent/baseline/descent/height/etc.
*/
type Font struct {
	Resource
	owner gdnative.Object
}

func (o *Font) BaseClass() string {
	return "Font"
}

/*
        Draw "string" into a canvas item using the font at a given position, with "modulate" color, and optionally clipping the width. "position" specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
	Args: [{ false canvas_item RID} { false position Vector2} { false string String} {1,1,1,1 true modulate Color} {-1 true clip_w int}], Returns: void
*/
func (o *Font) Draw(canvasItem gdnative.Rid, position gdnative.Vector2, string gdnative.String, modulate gdnative.Color, clipW gdnative.Int) {
	//log.Println("Calling Font.Draw()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromRid(canvasItem)
	ptrArguments[1] = gdnative.NewPointerFromVector2(position)
	ptrArguments[2] = gdnative.NewPointerFromString(string)
	ptrArguments[3] = gdnative.NewPointerFromColor(modulate)
	ptrArguments[4] = gdnative.NewPointerFromInt(clipW)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Font", "draw")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Draw character "char" into a canvas item using the font at a given position, with "modulate" color, and optionally kerning if "next" is passed. clipping the width. "position" specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis. The width used by the character is returned, making this function useful for drawing strings character by character.
	Args: [{ false canvas_item RID} { false position Vector2} { false char int} {-1 true next int} {1,1,1,1 true modulate Color}], Returns: float
*/
func (o *Font) DrawChar(canvasItem gdnative.Rid, position gdnative.Vector2, char gdnative.Int, next gdnative.Int, modulate gdnative.Color) gdnative.Float {
	//log.Println("Calling Font.DrawChar()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromRid(canvasItem)
	ptrArguments[1] = gdnative.NewPointerFromVector2(position)
	ptrArguments[2] = gdnative.NewPointerFromInt(char)
	ptrArguments[3] = gdnative.NewPointerFromInt(next)
	ptrArguments[4] = gdnative.NewPointerFromColor(modulate)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Font", "draw_char")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Return the font ascent (number of pixels above the baseline).
	Args: [], Returns: float
*/
func (o *Font) GetAscent() gdnative.Float {
	//log.Println("Calling Font.GetAscent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Font", "get_ascent")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Return the font descent (number of pixels below the baseline).
	Args: [], Returns: float
*/
func (o *Font) GetDescent() gdnative.Float {
	//log.Println("Calling Font.GetDescent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Font", "get_descent")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Return the total font height (ascent plus descent) in pixels.
	Args: [], Returns: float
*/
func (o *Font) GetHeight() gdnative.Float {
	//log.Println("Calling Font.GetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Font", "get_height")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Return the size of a string, taking kerning and advance into account.
	Args: [{ false string String}], Returns: Vector2
*/
func (o *Font) GetStringSize(string gdnative.String) gdnative.Vector2 {
	//log.Println("Calling Font.GetStringSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(string)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Font", "get_string_size")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: bool
*/
func (o *Font) IsDistanceFieldHint() gdnative.Bool {
	//log.Println("Calling Font.IsDistanceFieldHint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Font", "is_distance_field_hint")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        After editing a font (changing size, ascent, char rects, etc.). Call this function to propagate changes to controls that might use it.
	Args: [], Returns: void
*/
func (o *Font) UpdateChanges() {
	//log.Println("Calling Font.UpdateChanges()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Font", "update_changes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// FontImplementer is an interface that implements the methods
// of the Font class.
type FontImplementer interface {
	ResourceImplementer
	Draw(canvasItem gdnative.Rid, position gdnative.Vector2, string gdnative.String, modulate gdnative.Color, clipW gdnative.Int)
	DrawChar(canvasItem gdnative.Rid, position gdnative.Vector2, char gdnative.Int, next gdnative.Int, modulate gdnative.Color) gdnative.Float
	GetAscent() gdnative.Float
	GetDescent() gdnative.Float
	GetHeight() gdnative.Float
	GetStringSize(string gdnative.String) gdnative.Vector2
	IsDistanceFieldHint() gdnative.Bool
	UpdateChanges()
}

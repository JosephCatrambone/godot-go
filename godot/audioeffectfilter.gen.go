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

//func NewAudioEffectFilterFromPointer(ptr gdnative.Pointer) AudioEffectFilter {
func newAudioEffectFilterFromPointer(ptr gdnative.Pointer) AudioEffectFilter {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectFilter{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Allows frequencies other than the [member cutoff_hz] to pass.
*/
type AudioEffectFilter struct {
	AudioEffect
	owner gdnative.Object
}

func (o *AudioEffectFilter) BaseClass() string {
	return "AudioEffectFilter"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectFilter) GetCutoff() gdnative.Float {
	//log.Println("Calling AudioEffectFilter.GetCutoff()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectFilter", "get_cutoff")

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
	Args: [], Returns: enum.AudioEffectFilter::FilterDB
*/

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectFilter) GetGain() gdnative.Float {
	//log.Println("Calling AudioEffectFilter.GetGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectFilter", "get_gain")

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
func (o *AudioEffectFilter) GetResonance() gdnative.Float {
	//log.Println("Calling AudioEffectFilter.GetResonance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectFilter", "get_resonance")

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
	Args: [{ false freq float}], Returns: void
*/
func (o *AudioEffectFilter) SetCutoff(freq gdnative.Float) {
	//log.Println("Calling AudioEffectFilter.SetCutoff()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(freq)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectFilter", "set_cutoff")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount int}], Returns: void
*/
func (o *AudioEffectFilter) SetDb(amount gdnative.Int) {
	//log.Println("Calling AudioEffectFilter.SetDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectFilter", "set_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectFilter) SetGain(amount gdnative.Float) {
	//log.Println("Calling AudioEffectFilter.SetGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectFilter", "set_gain")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *AudioEffectFilter) SetResonance(amount gdnative.Float) {
	//log.Println("Calling AudioEffectFilter.SetResonance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectFilter", "set_resonance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioEffectFilterImplementer is an interface that implements the methods
// of the AudioEffectFilter class.
type AudioEffectFilterImplementer interface {
	AudioEffectImplementer
	GetCutoff() gdnative.Float
	GetGain() gdnative.Float
	GetResonance() gdnative.Float
	SetCutoff(freq gdnative.Float)
	SetDb(amount gdnative.Int)
	SetGain(amount gdnative.Float)
	SetResonance(amount gdnative.Float)
}

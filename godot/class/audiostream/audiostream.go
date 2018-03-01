package audiostream

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
Base class for audio streams. Audio streams are used for music playback, or other types of streamed sounds that don't fit or require more flexibility than a [Sample].
*/
type AudioStream struct {
	Resource
}

func (o *AudioStream) BaseClass() string {
	return "AudioStream"
}

/*

 */
func (o *AudioStream) GetLength() gdnative.Float {
	log.Println("Calling AudioStream.GetLength()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_length", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   AudioStreamImplementer is an interface for AudioStream objects.
*/
type AudioStreamImplementer interface {
	Class
}
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

//func NewEditorScenePostImportFromPointer(ptr gdnative.Pointer) EditorScenePostImport {
func newEditorScenePostImportFromPointer(ptr gdnative.Pointer) EditorScenePostImport {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorScenePostImport{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type EditorScenePostImport struct {
	Reference
	owner gdnative.Object
}

func (o *EditorScenePostImport) BaseClass() string {
	return "EditorScenePostImport"
}

/*

	Args: [{ false scene Object}], Returns: void
*/
func (o *EditorScenePostImport) PostImport(scene Object) {
	//log.Println("Calling EditorScenePostImport.PostImport()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(scene.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorScenePostImport", "post_import")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EditorScenePostImportImplementer is an interface that implements the methods
// of the EditorScenePostImport class.
type EditorScenePostImportImplementer interface {
	ReferenceImplementer
	PostImport(scene Object)
}

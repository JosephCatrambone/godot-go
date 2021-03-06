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

//func NewMeshFromPointer(ptr gdnative.Pointer) Mesh {
func newMeshFromPointer(ptr gdnative.Pointer) Mesh {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Mesh{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Mesh is a type of [Resource] that contains vertex-array based geometry, divided in [i]surfaces[/i]. Each surface contains a completely separate array and a material used to draw it. Design wise, a mesh with multiple surfaces is preferred to a single surface, because objects created in 3D editing software commonly contain multiple materials.
*/
type Mesh struct {
	Resource
	owner gdnative.Object
}

func (o *Mesh) BaseClass() string {
	return "Mesh"
}

/*
        Calculate a [ConvexPolygonShape] from the mesh.
	Args: [], Returns: Shape
*/
func (o *Mesh) CreateConvexShape() ShapeImplementer {
	//log.Println("Calling Mesh.CreateConvexShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Mesh", "create_convex_shape")

	// Call the parent method.
	// Shape
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newShapeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ShapeImplementer)
	}

	return &ret
}

/*
        Calculate an outline mesh at a defined offset (margin) from the original mesh. Note: Typically returns the vertices in reverse order (e.g. clockwise to anti-clockwise).
	Args: [{ false margin float}], Returns: Mesh
*/
func (o *Mesh) CreateOutline(margin gdnative.Float) MeshImplementer {
	//log.Println("Calling Mesh.CreateOutline()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Mesh", "create_outline")

	// Call the parent method.
	// Mesh
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newMeshFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(MeshImplementer)
	}

	return &ret
}

/*
        Calculate a [ConcavePolygonShape] from the mesh.
	Args: [], Returns: Shape
*/
func (o *Mesh) CreateTrimeshShape() ShapeImplementer {
	//log.Println("Calling Mesh.CreateTrimeshShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Mesh", "create_trimesh_shape")

	// Call the parent method.
	// Shape
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newShapeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ShapeImplementer)
	}

	return &ret
}

/*
        Generate a [TriangleMesh] from the mesh.
	Args: [], Returns: TriangleMesh
*/
func (o *Mesh) GenerateTriangleMesh() TriangleMeshImplementer {
	//log.Println("Calling Mesh.GenerateTriangleMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Mesh", "generate_triangle_mesh")

	// Call the parent method.
	// TriangleMesh
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTriangleMeshFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TriangleMeshImplementer)
	}

	return &ret
}

/*
        Returns all the vertices that make up the faces of the mesh. Each three vertices represent one triangle.
	Args: [], Returns: PoolVector3Array
*/
func (o *Mesh) GetFaces() gdnative.PoolVector3Array {
	//log.Println("Calling Mesh.GetFaces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Mesh", "get_faces")

	// Call the parent method.
	// PoolVector3Array
	retPtr := gdnative.NewEmptyPoolVector3Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector3ArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *Mesh) GetLightmapSizeHint() gdnative.Vector2 {
	//log.Println("Calling Mesh.GetLightmapSizeHint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Mesh", "get_lightmap_size_hint")

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
	Args: [{ false size Vector2}], Returns: void
*/
func (o *Mesh) SetLightmapSizeHint(size gdnative.Vector2) {
	//log.Println("Calling Mesh.SetLightmapSizeHint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Mesh", "set_lightmap_size_hint")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// MeshImplementer is an interface that implements the methods
// of the Mesh class.
type MeshImplementer interface {
	ResourceImplementer
	CreateConvexShape() ShapeImplementer
	CreateOutline(margin gdnative.Float) MeshImplementer
	CreateTrimeshShape() ShapeImplementer
	GenerateTriangleMesh() TriangleMeshImplementer
	GetFaces() gdnative.PoolVector3Array
	GetLightmapSizeHint() gdnative.Vector2
	SetLightmapSizeHint(size gdnative.Vector2)
}

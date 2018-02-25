package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include "gdnative.gen.h"
#include <gdnative/pool_arrays.h>
*/
import "C"

type PoolArrayReadAccess struct {
	base *C.godot_pool_array_read_access
}

func (gdt PoolArrayReadAccess) getBase() *C.godot_pool_array_read_access {
	return gdt.base
}

type PoolByteArrayReadAccess PoolArrayReadAccess

type PoolIntArrayReadAccess PoolArrayReadAccess

type PoolRealArrayReadAccess PoolArrayReadAccess

type PoolStringArrayReadAccess PoolArrayReadAccess

type PoolVector2ArrayReadAccess PoolArrayReadAccess

type PoolVector3ArrayReadAccess PoolArrayReadAccess

type PoolColorArrayReadAccess PoolArrayReadAccess

type PoolArrayWriteAccess struct {
	base *C.godot_pool_array_write_access
}

func (gdt PoolArrayWriteAccess) getBase() *C.godot_pool_array_write_access {
	return gdt.base
}

type PoolByteArrayWriteAccess PoolArrayWriteAccess

type PoolIntArrayWriteAccess PoolArrayWriteAccess

type PoolRealArrayWriteAccess PoolArrayWriteAccess

type PoolStringArrayWriteAccess PoolArrayWriteAccess

type PoolVector2ArrayWriteAccess PoolArrayWriteAccess

type PoolVector3ArrayWriteAccess PoolArrayWriteAccess

type PoolColorArrayWriteAccess PoolArrayWriteAccess

type PoolByteArray struct {
	base *C.godot_pool_byte_array
}

func (gdt PoolByteArray) getBase() *C.godot_pool_byte_array {
	return gdt.base
}

// Append godot_pool_byte_array_append [[godot_pool_byte_array * p_self] [const uint8_t p_data]] void
func (gdt *PoolByteArray) Append(data Uint8T) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_byte_array_append(GDNative.api, arg0, arg1)
}

// AppendArray godot_pool_byte_array_append_array [[godot_pool_byte_array * p_self] [const godot_pool_byte_array * p_array]] void
func (gdt *PoolByteArray) AppendArray(array PoolByteArray) {
	arg0 := gdt.getBase()
	arg1 := array.getBase()

	C.go_godot_pool_byte_array_append_array(GDNative.api, arg0, arg1)
}

// Insert godot_pool_byte_array_insert [[godot_pool_byte_array * p_self] [const godot_int p_idx] [const uint8_t p_data]] godot_error
func (gdt *PoolByteArray) Insert(idx Int, data Uint8T) Error {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	ret := C.go_godot_pool_byte_array_insert(GDNative.api, arg0, arg1, arg2)

	return Error(ret)
}

// Invert godot_pool_byte_array_invert [[godot_pool_byte_array * p_self]] void
func (gdt *PoolByteArray) Invert() {
	arg0 := gdt.getBase()

	C.go_godot_pool_byte_array_invert(GDNative.api, arg0)
}

// PushBack godot_pool_byte_array_push_back [[godot_pool_byte_array * p_self] [const uint8_t p_data]] void
func (gdt *PoolByteArray) PushBack(data Uint8T) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_byte_array_push_back(GDNative.api, arg0, arg1)
}

// Remove godot_pool_byte_array_remove [[godot_pool_byte_array * p_self] [const godot_int p_idx]] void
func (gdt *PoolByteArray) Remove(idx Int) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	C.go_godot_pool_byte_array_remove(GDNative.api, arg0, arg1)
}

// Resize godot_pool_byte_array_resize [[godot_pool_byte_array * p_self] [const godot_int p_size]] void
func (gdt *PoolByteArray) Resize(size Int) {
	arg0 := gdt.getBase()
	arg1 := size.getBase()

	C.go_godot_pool_byte_array_resize(GDNative.api, arg0, arg1)
}

// Read godot_pool_byte_array_read [[const godot_pool_byte_array * p_self]] godot_pool_byte_array_read_access *
func (gdt *PoolByteArray) Read() PoolByteArrayReadAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_byte_array_read(GDNative.api, arg0)

	return PoolByteArrayReadAccess{base: ret}
}

// Write godot_pool_byte_array_write [[godot_pool_byte_array * p_self]] godot_pool_byte_array_write_access *
func (gdt *PoolByteArray) Write() PoolByteArrayWriteAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_byte_array_write(GDNative.api, arg0)

	return PoolByteArrayWriteAccess{base: ret}
}

// Set godot_pool_byte_array_set [[godot_pool_byte_array * p_self] [const godot_int p_idx] [const uint8_t p_data]] void
func (gdt *PoolByteArray) Set(idx Int, data Uint8T) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	C.go_godot_pool_byte_array_set(GDNative.api, arg0, arg1, arg2)
}

// Get godot_pool_byte_array_get [[const godot_pool_byte_array * p_self] [const godot_int p_idx]] uint8_t
func (gdt *PoolByteArray) Get(idx Int) Uint8T {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_pool_byte_array_get(GDNative.api, arg0, arg1)

	return Uint8T(ret)
}

// Size godot_pool_byte_array_size [[const godot_pool_byte_array * p_self]] godot_int
func (gdt *PoolByteArray) Size() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_byte_array_size(GDNative.api, arg0)

	return Int(ret)
}

// Destroy godot_pool_byte_array_destroy [[godot_pool_byte_array * p_self]] void
func (gdt *PoolByteArray) Destroy() {
	arg0 := gdt.getBase()

	C.go_godot_pool_byte_array_destroy(GDNative.api, arg0)
}

type PoolIntArray struct {
	base *C.godot_pool_int_array
}

func (gdt PoolIntArray) getBase() *C.godot_pool_int_array {
	return gdt.base
}

// Append godot_pool_int_array_append [[godot_pool_int_array * p_self] [const godot_int p_data]] void
func (gdt *PoolIntArray) Append(data Int) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_int_array_append(GDNative.api, arg0, arg1)
}

// AppendArray godot_pool_int_array_append_array [[godot_pool_int_array * p_self] [const godot_pool_int_array * p_array]] void
func (gdt *PoolIntArray) AppendArray(array PoolIntArray) {
	arg0 := gdt.getBase()
	arg1 := array.getBase()

	C.go_godot_pool_int_array_append_array(GDNative.api, arg0, arg1)
}

// Insert godot_pool_int_array_insert [[godot_pool_int_array * p_self] [const godot_int p_idx] [const godot_int p_data]] godot_error
func (gdt *PoolIntArray) Insert(idx Int, data Int) Error {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	ret := C.go_godot_pool_int_array_insert(GDNative.api, arg0, arg1, arg2)

	return Error(ret)
}

// Invert godot_pool_int_array_invert [[godot_pool_int_array * p_self]] void
func (gdt *PoolIntArray) Invert() {
	arg0 := gdt.getBase()

	C.go_godot_pool_int_array_invert(GDNative.api, arg0)
}

// PushBack godot_pool_int_array_push_back [[godot_pool_int_array * p_self] [const godot_int p_data]] void
func (gdt *PoolIntArray) PushBack(data Int) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_int_array_push_back(GDNative.api, arg0, arg1)
}

// Remove godot_pool_int_array_remove [[godot_pool_int_array * p_self] [const godot_int p_idx]] void
func (gdt *PoolIntArray) Remove(idx Int) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	C.go_godot_pool_int_array_remove(GDNative.api, arg0, arg1)
}

// Resize godot_pool_int_array_resize [[godot_pool_int_array * p_self] [const godot_int p_size]] void
func (gdt *PoolIntArray) Resize(size Int) {
	arg0 := gdt.getBase()
	arg1 := size.getBase()

	C.go_godot_pool_int_array_resize(GDNative.api, arg0, arg1)
}

// Read godot_pool_int_array_read [[const godot_pool_int_array * p_self]] godot_pool_int_array_read_access *
func (gdt *PoolIntArray) Read() PoolIntArrayReadAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_int_array_read(GDNative.api, arg0)

	return PoolIntArrayReadAccess{base: ret}
}

// Write godot_pool_int_array_write [[godot_pool_int_array * p_self]] godot_pool_int_array_write_access *
func (gdt *PoolIntArray) Write() PoolIntArrayWriteAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_int_array_write(GDNative.api, arg0)

	return PoolIntArrayWriteAccess{base: ret}
}

// Set godot_pool_int_array_set [[godot_pool_int_array * p_self] [const godot_int p_idx] [const godot_int p_data]] void
func (gdt *PoolIntArray) Set(idx Int, data Int) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	C.go_godot_pool_int_array_set(GDNative.api, arg0, arg1, arg2)
}

// Get godot_pool_int_array_get [[const godot_pool_int_array * p_self] [const godot_int p_idx]] godot_int
func (gdt *PoolIntArray) Get(idx Int) Int {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_pool_int_array_get(GDNative.api, arg0, arg1)

	return Int(ret)
}

// Size godot_pool_int_array_size [[const godot_pool_int_array * p_self]] godot_int
func (gdt *PoolIntArray) Size() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_int_array_size(GDNative.api, arg0)

	return Int(ret)
}

// Destroy godot_pool_int_array_destroy [[godot_pool_int_array * p_self]] void
func (gdt *PoolIntArray) Destroy() {
	arg0 := gdt.getBase()

	C.go_godot_pool_int_array_destroy(GDNative.api, arg0)
}

type PoolRealArray struct {
	base *C.godot_pool_real_array
}

func (gdt PoolRealArray) getBase() *C.godot_pool_real_array {
	return gdt.base
}

// Append godot_pool_real_array_append [[godot_pool_real_array * p_self] [const godot_real p_data]] void
func (gdt *PoolRealArray) Append(data Real) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_real_array_append(GDNative.api, arg0, arg1)
}

// AppendArray godot_pool_real_array_append_array [[godot_pool_real_array * p_self] [const godot_pool_real_array * p_array]] void
func (gdt *PoolRealArray) AppendArray(array PoolRealArray) {
	arg0 := gdt.getBase()
	arg1 := array.getBase()

	C.go_godot_pool_real_array_append_array(GDNative.api, arg0, arg1)
}

// Insert godot_pool_real_array_insert [[godot_pool_real_array * p_self] [const godot_int p_idx] [const godot_real p_data]] godot_error
func (gdt *PoolRealArray) Insert(idx Int, data Real) Error {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	ret := C.go_godot_pool_real_array_insert(GDNative.api, arg0, arg1, arg2)

	return Error(ret)
}

// Invert godot_pool_real_array_invert [[godot_pool_real_array * p_self]] void
func (gdt *PoolRealArray) Invert() {
	arg0 := gdt.getBase()

	C.go_godot_pool_real_array_invert(GDNative.api, arg0)
}

// PushBack godot_pool_real_array_push_back [[godot_pool_real_array * p_self] [const godot_real p_data]] void
func (gdt *PoolRealArray) PushBack(data Real) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_real_array_push_back(GDNative.api, arg0, arg1)
}

// Remove godot_pool_real_array_remove [[godot_pool_real_array * p_self] [const godot_int p_idx]] void
func (gdt *PoolRealArray) Remove(idx Int) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	C.go_godot_pool_real_array_remove(GDNative.api, arg0, arg1)
}

// Resize godot_pool_real_array_resize [[godot_pool_real_array * p_self] [const godot_int p_size]] void
func (gdt *PoolRealArray) Resize(size Int) {
	arg0 := gdt.getBase()
	arg1 := size.getBase()

	C.go_godot_pool_real_array_resize(GDNative.api, arg0, arg1)
}

// Read godot_pool_real_array_read [[const godot_pool_real_array * p_self]] godot_pool_real_array_read_access *
func (gdt *PoolRealArray) Read() PoolRealArrayReadAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_real_array_read(GDNative.api, arg0)

	return PoolRealArrayReadAccess{base: ret}
}

// Write godot_pool_real_array_write [[godot_pool_real_array * p_self]] godot_pool_real_array_write_access *
func (gdt *PoolRealArray) Write() PoolRealArrayWriteAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_real_array_write(GDNative.api, arg0)

	return PoolRealArrayWriteAccess{base: ret}
}

// Set godot_pool_real_array_set [[godot_pool_real_array * p_self] [const godot_int p_idx] [const godot_real p_data]] void
func (gdt *PoolRealArray) Set(idx Int, data Real) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	C.go_godot_pool_real_array_set(GDNative.api, arg0, arg1, arg2)
}

// Get godot_pool_real_array_get [[const godot_pool_real_array * p_self] [const godot_int p_idx]] godot_real
func (gdt *PoolRealArray) Get(idx Int) Real {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_pool_real_array_get(GDNative.api, arg0, arg1)

	return Real(ret)
}

// Size godot_pool_real_array_size [[const godot_pool_real_array * p_self]] godot_int
func (gdt *PoolRealArray) Size() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_real_array_size(GDNative.api, arg0)

	return Int(ret)
}

// Destroy godot_pool_real_array_destroy [[godot_pool_real_array * p_self]] void
func (gdt *PoolRealArray) Destroy() {
	arg0 := gdt.getBase()

	C.go_godot_pool_real_array_destroy(GDNative.api, arg0)
}

type PoolStringArray struct {
	base *C.godot_pool_string_array
}

func (gdt PoolStringArray) getBase() *C.godot_pool_string_array {
	return gdt.base
}

// Append godot_pool_string_array_append [[godot_pool_string_array * p_self] [const godot_string * p_data]] void
func (gdt *PoolStringArray) Append(data String) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_string_array_append(GDNative.api, arg0, arg1)
}

// AppendArray godot_pool_string_array_append_array [[godot_pool_string_array * p_self] [const godot_pool_string_array * p_array]] void
func (gdt *PoolStringArray) AppendArray(array PoolStringArray) {
	arg0 := gdt.getBase()
	arg1 := array.getBase()

	C.go_godot_pool_string_array_append_array(GDNative.api, arg0, arg1)
}

// Insert godot_pool_string_array_insert [[godot_pool_string_array * p_self] [const godot_int p_idx] [const godot_string * p_data]] godot_error
func (gdt *PoolStringArray) Insert(idx Int, data String) Error {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	ret := C.go_godot_pool_string_array_insert(GDNative.api, arg0, arg1, arg2)

	return Error(ret)
}

// Invert godot_pool_string_array_invert [[godot_pool_string_array * p_self]] void
func (gdt *PoolStringArray) Invert() {
	arg0 := gdt.getBase()

	C.go_godot_pool_string_array_invert(GDNative.api, arg0)
}

// PushBack godot_pool_string_array_push_back [[godot_pool_string_array * p_self] [const godot_string * p_data]] void
func (gdt *PoolStringArray) PushBack(data String) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_string_array_push_back(GDNative.api, arg0, arg1)
}

// Remove godot_pool_string_array_remove [[godot_pool_string_array * p_self] [const godot_int p_idx]] void
func (gdt *PoolStringArray) Remove(idx Int) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	C.go_godot_pool_string_array_remove(GDNative.api, arg0, arg1)
}

// Resize godot_pool_string_array_resize [[godot_pool_string_array * p_self] [const godot_int p_size]] void
func (gdt *PoolStringArray) Resize(size Int) {
	arg0 := gdt.getBase()
	arg1 := size.getBase()

	C.go_godot_pool_string_array_resize(GDNative.api, arg0, arg1)
}

// Read godot_pool_string_array_read [[const godot_pool_string_array * p_self]] godot_pool_string_array_read_access *
func (gdt *PoolStringArray) Read() PoolStringArrayReadAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_string_array_read(GDNative.api, arg0)

	return PoolStringArrayReadAccess{base: ret}
}

// Write godot_pool_string_array_write [[godot_pool_string_array * p_self]] godot_pool_string_array_write_access *
func (gdt *PoolStringArray) Write() PoolStringArrayWriteAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_string_array_write(GDNative.api, arg0)

	return PoolStringArrayWriteAccess{base: ret}
}

// Set godot_pool_string_array_set [[godot_pool_string_array * p_self] [const godot_int p_idx] [const godot_string * p_data]] void
func (gdt *PoolStringArray) Set(idx Int, data String) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	C.go_godot_pool_string_array_set(GDNative.api, arg0, arg1, arg2)
}

// Get godot_pool_string_array_get [[const godot_pool_string_array * p_self] [const godot_int p_idx]] godot_string
func (gdt *PoolStringArray) Get(idx Int) String {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_pool_string_array_get(GDNative.api, arg0, arg1)

	return String{base: ret}
}

// Size godot_pool_string_array_size [[const godot_pool_string_array * p_self]] godot_int
func (gdt *PoolStringArray) Size() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_string_array_size(GDNative.api, arg0)

	return Int(ret)
}

// Destroy godot_pool_string_array_destroy [[godot_pool_string_array * p_self]] void
func (gdt *PoolStringArray) Destroy() {
	arg0 := gdt.getBase()

	C.go_godot_pool_string_array_destroy(GDNative.api, arg0)
}

type PoolVector2Array struct {
	base *C.godot_pool_vector2_array
}

func (gdt PoolVector2Array) getBase() *C.godot_pool_vector2_array {
	return gdt.base
}

// Append godot_pool_vector2_array_append [[godot_pool_vector2_array * p_self] [const godot_vector2 * p_data]] void
func (gdt *PoolVector2Array) Append(data Vector2) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_vector2_array_append(GDNative.api, arg0, arg1)
}

// AppendArray godot_pool_vector2_array_append_array [[godot_pool_vector2_array * p_self] [const godot_pool_vector2_array * p_array]] void
func (gdt *PoolVector2Array) AppendArray(array PoolVector2Array) {
	arg0 := gdt.getBase()
	arg1 := array.getBase()

	C.go_godot_pool_vector2_array_append_array(GDNative.api, arg0, arg1)
}

// Insert godot_pool_vector2_array_insert [[godot_pool_vector2_array * p_self] [const godot_int p_idx] [const godot_vector2 * p_data]] godot_error
func (gdt *PoolVector2Array) Insert(idx Int, data Vector2) Error {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	ret := C.go_godot_pool_vector2_array_insert(GDNative.api, arg0, arg1, arg2)

	return Error(ret)
}

// Invert godot_pool_vector2_array_invert [[godot_pool_vector2_array * p_self]] void
func (gdt *PoolVector2Array) Invert() {
	arg0 := gdt.getBase()

	C.go_godot_pool_vector2_array_invert(GDNative.api, arg0)
}

// PushBack godot_pool_vector2_array_push_back [[godot_pool_vector2_array * p_self] [const godot_vector2 * p_data]] void
func (gdt *PoolVector2Array) PushBack(data Vector2) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_vector2_array_push_back(GDNative.api, arg0, arg1)
}

// Remove godot_pool_vector2_array_remove [[godot_pool_vector2_array * p_self] [const godot_int p_idx]] void
func (gdt *PoolVector2Array) Remove(idx Int) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	C.go_godot_pool_vector2_array_remove(GDNative.api, arg0, arg1)
}

// Resize godot_pool_vector2_array_resize [[godot_pool_vector2_array * p_self] [const godot_int p_size]] void
func (gdt *PoolVector2Array) Resize(size Int) {
	arg0 := gdt.getBase()
	arg1 := size.getBase()

	C.go_godot_pool_vector2_array_resize(GDNative.api, arg0, arg1)
}

// Read godot_pool_vector2_array_read [[const godot_pool_vector2_array * p_self]] godot_pool_vector2_array_read_access *
func (gdt *PoolVector2Array) Read() PoolVector2ArrayReadAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_vector2_array_read(GDNative.api, arg0)

	return PoolVector2ArrayReadAccess{base: ret}
}

// Write godot_pool_vector2_array_write [[godot_pool_vector2_array * p_self]] godot_pool_vector2_array_write_access *
func (gdt *PoolVector2Array) Write() PoolVector2ArrayWriteAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_vector2_array_write(GDNative.api, arg0)

	return PoolVector2ArrayWriteAccess{base: ret}
}

// Set godot_pool_vector2_array_set [[godot_pool_vector2_array * p_self] [const godot_int p_idx] [const godot_vector2 * p_data]] void
func (gdt *PoolVector2Array) Set(idx Int, data Vector2) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	C.go_godot_pool_vector2_array_set(GDNative.api, arg0, arg1, arg2)
}

// Get godot_pool_vector2_array_get [[const godot_pool_vector2_array * p_self] [const godot_int p_idx]] godot_vector2
func (gdt *PoolVector2Array) Get(idx Int) Vector2 {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_pool_vector2_array_get(GDNative.api, arg0, arg1)

	return Vector2{base: ret}
}

// Size godot_pool_vector2_array_size [[const godot_pool_vector2_array * p_self]] godot_int
func (gdt *PoolVector2Array) Size() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_vector2_array_size(GDNative.api, arg0)

	return Int(ret)
}

// Destroy godot_pool_vector2_array_destroy [[godot_pool_vector2_array * p_self]] void
func (gdt *PoolVector2Array) Destroy() {
	arg0 := gdt.getBase()

	C.go_godot_pool_vector2_array_destroy(GDNative.api, arg0)
}

type PoolVector3Array struct {
	base *C.godot_pool_vector3_array
}

func (gdt PoolVector3Array) getBase() *C.godot_pool_vector3_array {
	return gdt.base
}

// Append godot_pool_vector3_array_append [[godot_pool_vector3_array * p_self] [const godot_vector3 * p_data]] void
func (gdt *PoolVector3Array) Append(data Vector3) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_vector3_array_append(GDNative.api, arg0, arg1)
}

// AppendArray godot_pool_vector3_array_append_array [[godot_pool_vector3_array * p_self] [const godot_pool_vector3_array * p_array]] void
func (gdt *PoolVector3Array) AppendArray(array PoolVector3Array) {
	arg0 := gdt.getBase()
	arg1 := array.getBase()

	C.go_godot_pool_vector3_array_append_array(GDNative.api, arg0, arg1)
}

// Insert godot_pool_vector3_array_insert [[godot_pool_vector3_array * p_self] [const godot_int p_idx] [const godot_vector3 * p_data]] godot_error
func (gdt *PoolVector3Array) Insert(idx Int, data Vector3) Error {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	ret := C.go_godot_pool_vector3_array_insert(GDNative.api, arg0, arg1, arg2)

	return Error(ret)
}

// Invert godot_pool_vector3_array_invert [[godot_pool_vector3_array * p_self]] void
func (gdt *PoolVector3Array) Invert() {
	arg0 := gdt.getBase()

	C.go_godot_pool_vector3_array_invert(GDNative.api, arg0)
}

// PushBack godot_pool_vector3_array_push_back [[godot_pool_vector3_array * p_self] [const godot_vector3 * p_data]] void
func (gdt *PoolVector3Array) PushBack(data Vector3) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_vector3_array_push_back(GDNative.api, arg0, arg1)
}

// Remove godot_pool_vector3_array_remove [[godot_pool_vector3_array * p_self] [const godot_int p_idx]] void
func (gdt *PoolVector3Array) Remove(idx Int) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	C.go_godot_pool_vector3_array_remove(GDNative.api, arg0, arg1)
}

// Resize godot_pool_vector3_array_resize [[godot_pool_vector3_array * p_self] [const godot_int p_size]] void
func (gdt *PoolVector3Array) Resize(size Int) {
	arg0 := gdt.getBase()
	arg1 := size.getBase()

	C.go_godot_pool_vector3_array_resize(GDNative.api, arg0, arg1)
}

// Read godot_pool_vector3_array_read [[const godot_pool_vector3_array * p_self]] godot_pool_vector3_array_read_access *
func (gdt *PoolVector3Array) Read() PoolVector3ArrayReadAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_vector3_array_read(GDNative.api, arg0)

	return PoolVector3ArrayReadAccess{base: ret}
}

// Write godot_pool_vector3_array_write [[godot_pool_vector3_array * p_self]] godot_pool_vector3_array_write_access *
func (gdt *PoolVector3Array) Write() PoolVector3ArrayWriteAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_vector3_array_write(GDNative.api, arg0)

	return PoolVector3ArrayWriteAccess{base: ret}
}

// Set godot_pool_vector3_array_set [[godot_pool_vector3_array * p_self] [const godot_int p_idx] [const godot_vector3 * p_data]] void
func (gdt *PoolVector3Array) Set(idx Int, data Vector3) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	C.go_godot_pool_vector3_array_set(GDNative.api, arg0, arg1, arg2)
}

// Get godot_pool_vector3_array_get [[const godot_pool_vector3_array * p_self] [const godot_int p_idx]] godot_vector3
func (gdt *PoolVector3Array) Get(idx Int) Vector3 {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_pool_vector3_array_get(GDNative.api, arg0, arg1)

	return Vector3{base: ret}
}

// Size godot_pool_vector3_array_size [[const godot_pool_vector3_array * p_self]] godot_int
func (gdt *PoolVector3Array) Size() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_vector3_array_size(GDNative.api, arg0)

	return Int(ret)
}

// Destroy godot_pool_vector3_array_destroy [[godot_pool_vector3_array * p_self]] void
func (gdt *PoolVector3Array) Destroy() {
	arg0 := gdt.getBase()

	C.go_godot_pool_vector3_array_destroy(GDNative.api, arg0)
}

type PoolColorArray struct {
	base *C.godot_pool_color_array
}

func (gdt PoolColorArray) getBase() *C.godot_pool_color_array {
	return gdt.base
}

// Append godot_pool_color_array_append [[godot_pool_color_array * p_self] [const godot_color * p_data]] void
func (gdt *PoolColorArray) Append(data Color) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_color_array_append(GDNative.api, arg0, arg1)
}

// AppendArray godot_pool_color_array_append_array [[godot_pool_color_array * p_self] [const godot_pool_color_array * p_array]] void
func (gdt *PoolColorArray) AppendArray(array PoolColorArray) {
	arg0 := gdt.getBase()
	arg1 := array.getBase()

	C.go_godot_pool_color_array_append_array(GDNative.api, arg0, arg1)
}

// Insert godot_pool_color_array_insert [[godot_pool_color_array * p_self] [const godot_int p_idx] [const godot_color * p_data]] godot_error
func (gdt *PoolColorArray) Insert(idx Int, data Color) Error {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	ret := C.go_godot_pool_color_array_insert(GDNative.api, arg0, arg1, arg2)

	return Error(ret)
}

// Invert godot_pool_color_array_invert [[godot_pool_color_array * p_self]] void
func (gdt *PoolColorArray) Invert() {
	arg0 := gdt.getBase()

	C.go_godot_pool_color_array_invert(GDNative.api, arg0)
}

// PushBack godot_pool_color_array_push_back [[godot_pool_color_array * p_self] [const godot_color * p_data]] void
func (gdt *PoolColorArray) PushBack(data Color) {
	arg0 := gdt.getBase()
	arg1 := data.getBase()

	C.go_godot_pool_color_array_push_back(GDNative.api, arg0, arg1)
}

// Remove godot_pool_color_array_remove [[godot_pool_color_array * p_self] [const godot_int p_idx]] void
func (gdt *PoolColorArray) Remove(idx Int) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	C.go_godot_pool_color_array_remove(GDNative.api, arg0, arg1)
}

// Resize godot_pool_color_array_resize [[godot_pool_color_array * p_self] [const godot_int p_size]] void
func (gdt *PoolColorArray) Resize(size Int) {
	arg0 := gdt.getBase()
	arg1 := size.getBase()

	C.go_godot_pool_color_array_resize(GDNative.api, arg0, arg1)
}

// Read godot_pool_color_array_read [[const godot_pool_color_array * p_self]] godot_pool_color_array_read_access *
func (gdt *PoolColorArray) Read() PoolColorArrayReadAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_color_array_read(GDNative.api, arg0)

	return PoolColorArrayReadAccess{base: ret}
}

// Write godot_pool_color_array_write [[godot_pool_color_array * p_self]] godot_pool_color_array_write_access *
func (gdt *PoolColorArray) Write() PoolColorArrayWriteAccess {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_color_array_write(GDNative.api, arg0)

	return PoolColorArrayWriteAccess{base: ret}
}

// Set godot_pool_color_array_set [[godot_pool_color_array * p_self] [const godot_int p_idx] [const godot_color * p_data]] void
func (gdt *PoolColorArray) Set(idx Int, data Color) {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()
	arg2 := data.getBase()

	C.go_godot_pool_color_array_set(GDNative.api, arg0, arg1, arg2)
}

// Get godot_pool_color_array_get [[const godot_pool_color_array * p_self] [const godot_int p_idx]] godot_color
func (gdt *PoolColorArray) Get(idx Int) Color {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_pool_color_array_get(GDNative.api, arg0, arg1)

	return Color{base: ret}
}

// Size godot_pool_color_array_size [[const godot_pool_color_array * p_self]] godot_int
func (gdt *PoolColorArray) Size() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_pool_color_array_size(GDNative.api, arg0)

	return Int(ret)
}

// Destroy godot_pool_color_array_destroy [[godot_pool_color_array * p_self]] void
func (gdt *PoolColorArray) Destroy() {
	arg0 := gdt.getBase()

	C.go_godot_pool_color_array_destroy(GDNative.api, arg0)
}

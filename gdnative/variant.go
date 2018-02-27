package gdnative

/*
#include <gdnative/string.h>
#include <gdnative/variant.h>
#include "gdnative.gen.h"
#include "variant.h"
*/
import "C"

func NewVariantWithString(str *String) Variant {
	var variant C.godot_variant
	C.go_godot_variant_new_string(GDNative.api, &variant, str.base)

	return Variant{base: &variant}
}

type VariantArray struct {
	base  **C.godot_variant
	array []Variant
}

func (gdt *VariantArray) getBase() **C.godot_variant {
	variantArray := C.go_godot_variant_build_array(C.int(len(gdt.array)))
	for i, variant := range gdt.array {
		C.go_godot_variant_add_element(variantArray, variant.getBase(), C.int(i))
	}

	return variantArray
}
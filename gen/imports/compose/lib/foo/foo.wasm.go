// Code generated by wit-bindgen-go. DO NOT EDIT.

package foo

// This file contains wasmimport and wasmexport declarations for "compose:lib@0.0.1".

//go:wasmimport compose:lib/foo@0.0.1 [resource-drop]foo-resource
//go:noescape
func wasmimport_FooResourceResourceDrop(self0 uint32)

//go:wasmimport compose:lib/foo@0.0.1 [constructor]foo-resource
//go:noescape
func wasmimport_NewFooResource() (result0 uint32)

//go:wasmimport compose:lib/foo@0.0.1 [method]foo-resource.foo
//go:noescape
func wasmimport_FooResourceFoo(self0 uint32, result *string)

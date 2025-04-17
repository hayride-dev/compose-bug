// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package foo represents the imported interface "compose:lib/foo@0.0.1".
package foo

import (
	"go.bytecodealliance.org/cm"
)

// FooResource represents the imported resource "compose:lib/foo@0.0.1#foo-resource".
//
//	resource foo-resource
type FooResource cm.Resource

// ResourceDrop represents the imported resource-drop for resource "foo-resource".
//
// Drops a resource handle.
//
//go:nosplit
func (self FooResource) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_FooResourceResourceDrop((uint32)(self0))
	return
}

// NewFooResource represents the imported constructor for resource "foo-resource".
//
//	constructor()
//
//go:nosplit
func NewFooResource() (result FooResource) {
	result0 := wasmimport_NewFooResource()
	result = cm.Reinterpret[FooResource]((uint32)(result0))
	return
}

// Foo represents the imported method "foo".
//
//	foo: func() -> string
//
//go:nosplit
func (self FooResource) Foo() (result string) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_FooResourceFoo((uint32)(self0), &result)
	return
}

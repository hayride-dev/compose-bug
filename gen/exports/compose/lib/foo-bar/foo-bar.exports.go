// Code generated by wit-bindgen-go. DO NOT EDIT.

package foobar

import (
	"go.bytecodealliance.org/cm"
)

// Exports represents the caller-defined exports from "compose:lib/foo-bar@0.0.1".
var Exports struct {
	// FooBarResource represents the caller-defined exports for resource "compose:lib/foo-bar@0.0.1#foo-bar-resource".
	FooBarResource struct {
		// Destructor represents the caller-defined, exported destructor for resource "foo-bar-resource".
		//
		// Resource destructor.
		Destructor func(self cm.Rep)

		// Constructor represents the caller-defined, exported constructor for resource "foo-bar-resource".
		//
		//	constructor()
		Constructor func() (result FooBarResource)

		// FooBar represents the caller-defined, exported method "foo-bar".
		//
		//	foo-bar: func(f: borrow<foo-resource>) -> string
		FooBar func(self cm.Rep, f cm.Rep) (result string)
	}
}

package main

import (
	foobar "github.com/hayride-dev/compose-bug/gen/exports/compose/lib/foo-bar"
	"github.com/hayride-dev/compose-bug/gen/imports/compose/lib/foo"
	"go.bytecodealliance.org/cm"
)

func init() {
	foobar.Exports.FooBarResource.Constructor = constructorFunc
	foobar.Exports.FooBarResource.FooBar = foobarFunc

}

func constructorFunc() foobar.FooBarResource {
	return foobar.FooBarResourceResourceNew(0)
}

// func(self cm.Rep, f FooResource) (result string)
func foobarFunc(self cm.Rep, f cm.Rep) string {
	fooResource := foo.FooResource(f)
	value := fooResource.Foo() + "bar"
	fooResource.ResourceDrop()
	return value
}

func main() {}

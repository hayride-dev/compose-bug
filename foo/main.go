package main

import (
	"github.com/hayride-dev/compose-bug/gen/exports/compose/lib/foo"
	"go.bytecodealliance.org/cm"
)

func init() {
	foo.Exports.FooResource.Constructor = constructorFunc
	foo.Exports.FooResource.Foo = fooFunc
}

func constructorFunc() foo.FooResource {
	return foo.FooResourceResourceNew(0)
}

func fooFunc(self cm.Rep) string {
	return "foo"
}

func main() {}

package main

import (
	"fmt"

	"github.com/hayride-dev/compose-bug/gen/imports/compose/lib/foo"
	foobar "github.com/hayride-dev/compose-bug/gen/imports/compose/lib/foo-bar"
)

func main() {
	fooResource := foo.NewFooResource()

	fmt.Println(fooResource.Foo())

	foobarResource := foobar.NewFooBarResource()

	fmt.Println(foobarResource.FooBar(fooResource))

	fmt.Println(foobarResource.FooBar(fooResource) + "baz")
}

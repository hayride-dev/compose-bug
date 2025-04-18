# Compose Bug

This repo was created a minimal reproducible example of the compose bug reference by issue: https://github.com/bytecodealliance/wac/issues/149

Requires [tinygo](https://github.com/tinygo-org/tinygo) and [wac](https://github.com/bytecodealliance/wac) to run.

Use `make` to run build and wac to replicate error.

```
make
```

TinyGo Build commands:
```
tinygo build -target wasip2 --wit-package ./foo/wit/ --wit-world foo -o foo.wasm ./foo/main.go
tinygo build -target wasip2 --wit-package ./foo-bar/wit/ --wit-world foo-bar -o foo-bar.wasm ./foo-bar/main.go
tinygo build -target wasip2 --wit-package ./foo-bar-baz/wit/ --wit-world foo-bar-baz -o foo-bar-baz.wasm ./foo-bar-baz/main.go
```

Wac Commands:

```
wac plug foo-bar.wasm --plug foo.wasm -o foo-bar-composed.wasm
wac plug foo-bar-baz.wasm --plug foo-bar-composed.wasm -o foo-bar-baz-composed.wasm
```

Currently produces the following error on wac plug:

```
Caused by:
    type mismatch for import `compose:lib/foo-bar@0.0.1`
    type mismatch in instance export `foo-resource`
    resource types are not the same (ResourceId { globally_unique_id: 2, contextually_unique_id: 51 } vs. ResourceId { globally_unique_id: 2, contextually_unique_id: 1 }) (at offset 0x15d26f)
make: *** [wac-foo-bar-baz] Error 1
```

## SOLUTION

Using a compose file and the `wac compose` command you can build this be providing the foo resource to each component separately:
```
// compose.wac
package example:composition;

let foo = new my:foo { ... };
let foo-bar = new my:foo-bar {
  foo: foo.foo,
  ...
};
let foo-bar-baz = new my:foo-bar-baz {
  foo: foo.foo,
  foo-bar: foo-bar.foo-bar,
  ...
};

export foo-bar-baz...;
```

Build with the compose command:
```
wac compose compose.wac --dep my:foo=./foo.wasm --dep my:foo-bar=./foo-bar.wasm --dep my:foo-bar-baz=./foo-bar-baz.wasm -o composed.wasm
```

## Wit Deps

wit-deps used to generate deps dir in each component:

```
wit-deps update
```

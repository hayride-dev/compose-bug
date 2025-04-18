.PHONY: all build build-foo build-foo-bar build-foo-bar-baz wac gen-imports gen-exports gen

all: build wac-foo-bar wac-foo-bar-baz

build: build-foo build-foo-bar build-foo-bar-baz

build-foo:
	tinygo build -target wasip2 --wit-package ./foo/wit/ --wit-world foo -o foo.wasm ./foo/main.go

build-foo-bar:
	tinygo build -target wasip2 --wit-package ./foo-bar/wit/ --wit-world foo-bar -o foo-bar.wasm ./foo-bar/main.go

build-foo-bar-baz:
	tinygo build -target wasip2 --wit-package ./foo-bar-baz/wit/ --wit-world foo-bar-baz -o foo-bar-baz.wasm ./foo-bar-baz/main.go

wac-foo-bar: 
	wac plug foo-bar.wasm --plug foo.wasm -o foo-bar-composed.wasm

wac-foo-bar-baz:
	wac plug foo-bar-baz.wasm --plug foo-bar-composed.wasm -o foo-bar-baz-composed.wasm

compose:
	wac compose compose.wac --dep my:foo=./foo.wasm --dep my:foo-bar=./foo-bar.wasm --dep my:foo-bar-baz=./foo-bar-baz.wasm -o composed.wasm

gen-imports:
	wit-bindgen-go generate --world compose:lib/imports --out ./gen/imports ./wit

gen-exports:
	wit-bindgen-go generate --world compose:lib/exports --out ./gen/exports ./wit

gen: gen-imports gen-exports

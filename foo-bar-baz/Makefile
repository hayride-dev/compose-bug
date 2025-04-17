.PHONY: deps gen

deps:
	wit-deps update

gen:
	wit-bindgen-go generate --world compose:socket/socket --out ./gen ./wit

build:
	tinygo build -target wasip2 --wit-package ./wit/ --wit-world foo-bar-baz

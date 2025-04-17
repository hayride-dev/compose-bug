.PHONY: deps gen

deps:
	wit-deps update

build:
	tinygo build -target wasip2 --wit-package ./wit/ --wit-world foo-bar

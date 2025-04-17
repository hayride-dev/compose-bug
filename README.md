# Compose Bug

This repo was created a minimal reproducible example of the compose bug reference by issue: https://github.com/bytecodealliance/wac/issues/149

Requires [tinygo](https://github.com/tinygo-org/tinygo) and [wac](https://github.com/bytecodealliance/wac) to run.

Use `make` to run build and wac to replicate error.

Currently produces the following error on wac plug:

```
Caused by:
    type mismatch for import `compose:lib/foo-bar@0.0.1`
    type mismatch in instance export `foo-resource`
    resource types are not the same (ResourceId { globally_unique_id: 2, contextually_unique_id: 51 } vs. ResourceId { globally_unique_id: 2, contextually_unique_id: 1 }) (at offset 0x15d26f)
make: *** [wac-foo-bar-baz] Error 1
```

## Wit Deps

wit-deps used to generate deps dir in each component:

```
wit-deps update
```

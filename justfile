# use with https://github.com/casey/just

_default:
    @just --list --unsorted --justfile {{justfile()}}

# Runs the raytracer
run:
    go run ./cmd/raytracer

# Runs the gradient test
gradient:
    go run ./cmd/gradient

# Deletes all .ppm images
clear:
    rm *.ppm

# Clears the cache
clean-cache:
    go clean -cache

# Clears the test cache and runs the tests
test:
    go clean -testcache
    go test -v ./...


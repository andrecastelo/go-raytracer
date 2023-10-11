# Runs the raytracer
run:
    go run ./cmd/raytracer

# Clears the cache
clean:
    go clean -cache

# Clears the test cache and runs the tests
test:
    go clean -testcache
    go test -v ./...


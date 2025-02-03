# useful zero value

One way of making unitialized structs useful without needing a `NewStruct()` method. You can leave `NewStruct()`-like methods for when there are parameters to be passed or more complicated initialization.

Inspiration: https://github.com/robpike/ivy/blob/58ca01e/config/config.go#L64

Run `go run .`

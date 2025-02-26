```
$ GOGC=off go run .
HeapAlloc = 65821216 B
TotalAlloc = 65821216 B
Sys = 77813792 B
NumGC = 0
```

```
$ go run .
HeapAlloc = 119888 B
TotalAlloc = 65830592 B
Sys = 78469152 B
NumGC = 3
```

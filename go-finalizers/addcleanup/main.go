package main

import (
	"fmt"
	"runtime"
	"time"
)

type Blob struct {
	buf []byte
}

func (b Blob) String() string {
	return fmt.Sprintf("Blob(%d KiB)", len(b.buf)/1024)
}

func newBlob(size int) *Blob {
	b := make([]byte, size)
	for i := range size {
		b[i] = 0
	}
	return &Blob{buf: b}
}

// current heap size in KiB.
func alloc() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / 1024
}

func main() {
	b := newBlob(5 * 1024 * 1024)
	// this is preferred over runtime.SetFinalizer
	runtime.AddCleanup(b, cleanup, time.Now())

	fmt.Printf("before GC, alloc size: %d KiB, blob: %s\n", alloc(), b)

	time.Sleep(10 * time.Millisecond)
	b = nil
	runtime.GC()

	// cleanup() runs after GC but there is no guarantee that it'll run before these
	// next lines. If the timeout is too small (like 1us) it might run after
	// the Printf at the end.
	//
	// It is not guaranteed to run at all. From `go doc runtime.AddCleanup`:
	//
	// * The cleanup(arg) call is not always guaranteed to run; in particular it is
	//   not guaranteed to run before program exit.
	// * Cleanups are not guaranteed to run if the size of T is zero bytes
	// * It is not guaranteed that a cleanup will run for objects allocated in
	//   initializers for package-level variables.

	time.Sleep(10 * time.Millisecond)
	fmt.Printf("after GC, alloc size: %d KiB, blob: %s\n", alloc(), b)
}

func cleanup(created time.Time) {
	fmt.Printf("GC ran, lifetime of object: %dms\n", time.Since(created)/time.Millisecond)
}

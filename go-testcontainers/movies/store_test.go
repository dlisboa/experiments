//go:build integration

package movies

import (
	"demo/internal/testdb"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	testdb.Create()
	testdb.Migrate()
	testdb.Seed()

	t := time.Now()
	code := m.Run()
	fmt.Fprintf(os.Stdout, "===== tests took: %v\n", time.Since(t))

	testdb.Terminate()
	os.Exit(code)
}

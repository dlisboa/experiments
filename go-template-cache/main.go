package main

// Example of caching templates with a map.
//
// Makes deliberate use of panic, but in production code they should return
// errors.

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//go:embed templates
var FS embed.FS

// NewTemplateCache creates a map of templates, keyed by their file names
// relative to basedir. Parses files recursively including sub-directories,
// e.g. basedir/foo/bar/baz.html. Includes basedir/base.html and
// basedir/partials.html automatically; panics if these files do not exist.
func newTemplateCache(resources fs.FS, basedir string) map[string]*template.Template {
	cache := make(map[string]*template.Template)

	walk := func(path string, d fs.DirEntry, err error) error {
		// skip non-HTML templates
		if !strings.Contains(d.Name(), ".html") {
			return nil
		}

		t, err := template.ParseFS(
			resources,
			filepath.Join(basedir, "base.html"),
			filepath.Join(basedir, "partials.html"),
			path,
		)
		if err != nil {
			return err
		}

		name := strings.TrimPrefix(path, basedir+"/")
		cache[name] = t
		return nil
	}
	err := fs.WalkDir(resources, basedir, walk)
	if err != nil {
		e := fmt.Errorf("template cache: fs.WalkDir: %w", err)
		panic(e)
	}

	return cache
}

func main() {
	cache := newTemplateCache(FS, "templates")
	fmt.Printf("cache: %+v\n", cache)

	pages := []string{"page1.html", "page2.html", "foo/page1.html"}
	for _, p := range pages {
		t, ok := cache[p]
		if !ok {
			panic("cannot find: " + p)
		}
		t.ExecuteTemplate(os.Stdout, "base", nil)
		println("----")
	}
}

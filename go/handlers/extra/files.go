package extra

import (
	"fmt"
	"log"
	"net/http"
	p "path/filepath"
	"strings"
	"time"

	"g.sacerb.com/sacer/go/entry"
	"g.sacerb.com/sacer/go/entry/tools"
	"g.sacerb.com/sacer/go/entry/types/set"
	"g.sacerb.com/sacer/go/entry/types/tree"
	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
	"g.sacerb.com/sacer/go/server/paths"
)

/*

	TODO: This package should be simplified.

*/

func ServeFile(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta, path *paths.Path) {
	err := serveFile(s, w, r, m, path)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
	}
}

func serveFile(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta, path *paths.Path) error {
	section := path.Section()
	tree := s.Trees[section].Access(m.Auth.Subscriber)[m.Lang]

	e, err := getEntry(tree, path)
	if err != nil {
		return err
	}

	col, ok := e.(entry.Collection)

	if !ok {
		return serveSingleBlob(w, r, e, path)
	}

	return serveCollectionBlob(w, r, col, path)
}

func serveSingleBlob(w http.ResponseWriter, r *http.Request, e entry.Entry, path *paths.Path) error {
	blob, ok := e.(entry.Blob)
	if !ok {
		return fmt.Errorf("file to serve (%v) is no blob", e.File().Name())
	}

	location, err := blob.Location(path.File.Ext, path.File.Option)
	if err != nil {
		return err
	}
	serveStatic(w, r, location)
	return nil
}

func serveCollectionBlob(w http.ResponseWriter, r *http.Request, col entry.Collection, path *paths.Path) error {
	name := baseName(path.File.Name)
	for _, e := range col.Entries() {
		if baseName(e.File().Name()) == name {
			return serveSingleBlob(w, r, e, path)
		}
	}

	if name := path.File.Name; len(name) > 5 && name[:5] == "cover" {
		set, ok := col.(*set.Set)
		if ok && set.Cover != nil {
			return serveSingleBlob(w, r, set.Cover, path)
		}
		t, ok := col.(*tree.Tree)
		if ok && t.Cover != nil {
			return serveSingleBlob(w, r, t.Cover, path)
		}
		return fmt.Errorf("serveCollectionBlob: Cover %v not found", path.File.Name)
	}

	return fmt.Errorf("serveCollectionBlob: File %v not found", path.File.Name)
}

func baseName(name string) string {
	name = stripBlur(name)
	return stripSize(name)
}

func stripSize(name string) string {
	i := strings.LastIndex(name, "-")
	if i > 0 {
		return name[:i]
	}
	return name
}

func stripBlur(name string) string {
	name = tools.StripExt(p.Base(name))
	if l := len(name); l > 4 && name[l-4] == '_' {
		return name[:l-4]
	}
	return name
}

func getEntry(t *tree.Tree, path *paths.Path) (entry.Entry, error) {
	hash := path.Hash
	if hash == "" {
		h, err := getMonthHash(path)
		if err != nil {
			return nil, err
		}
		hash = h
	}
	return t.LookupEntryHash(hash)
}

func getMonthHash(path *paths.Path) (string, error) {
	if len(path.Chain) != 3 {
		return "", fmt.Errorf("getMonthEntry: wrong month format. %v", path.Raw)
	}

	slug := path.Slug
	if paths.IsMergedMonths(path.Slug) {
		slug = slug[:2]
	}

	date, err := time.Parse("200601--150405", path.Chain[2]+slug+"--000001")
	if err != nil {
		return "", err
	}

	return tools.ToB16(date), nil
}

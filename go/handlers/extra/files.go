package extra

import (
	"fmt"
	"log"
	"net/http"

	//"path/filepath"
	"stferal/go/entry"
	"stferal/go/paths"
	"stferal/go/server"
)

// Another way to do it could be to go through Recents.
func ServeFile(s *server.Server, w http.ResponseWriter, r *http.Request, path *paths.Path) {
	err := serveFile(s, w, r, path)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
	}
}

func serveFile(s *server.Server, w http.ResponseWriter, r *http.Request, path *paths.Path) error {
	e, err := s.Trees[path.Section()].LookupEntryHash(path.Hash)
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
		return fmt.Errorf("File to serve (%v) is no blob.", e.File().Name())
	}
	serveStatic(w, r, blob.Location(path.SubFile.Size))
	return nil
}

func serveCollectionBlob(w http.ResponseWriter, r *http.Request, col entry.Collection, path *paths.Path) error {
	for _, e := range col.Entries() {
		if e.File().Name() == path.SubFile.Name {
			return serveSingleBlob(w, r, e, path)
		}
	}
	return fmt.Errorf("serveCollectionBlob: File %v not found.", path.SubFile.Name)
}

/*
func serveStandalone(w http.ResponseWriter, r *http.Request, e entry.Entry, subpath string) {
	e.Location(subpath)¶
	if e.Type() == "image" {
		return e.
	}
	//filename, size := paths.SplitSubpath(subpath)
	var abs string
	var err error

	switch e.Type() {
	case *entry.Image:
		abs, err = eh.(*entry.Image).ImageAbs(size), nil
	case "set":
		abs, err = findSetFile(eh.(*entry.Set), filename, size)
	case *entry.Hold:
		abs, err = findHoldFile(eh.(*entry.Hold), filename, size)
	default:
		err = fmt.Errorf("Cannot search cache file in #%v#. %v", entry.Type(eh), eh)
	}

	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	serveStatic(w, r, abs)
}

func serveStandalone(w http.ResponseWriter, r *http.Request, e entry.Entry, subpath string) {
func findHoldFile(h *entry.Hold, name, size string) (string, error) {
	//if name == "cover.jpg" && set.Cover != nil {
	//	return set.Cover.ImageAbs(size), nil
	//}

	for _, e := range h.Els {
		switch e.(type) {
		case *entry.Image:
			if e.(*entry.Image).File.Base() == name {
				return e.(*entry.Image).ImageAbs(size), nil
			}
		}
	}

	return "", fmt.Errorf("Could not find cache file (%v) in Hold (%v)", name, h)
}

func findSetFile(set *entry.Set, name, size string) (string, error) {
	if name == "cover.jpg" && set.Cover != nil {
		return set.Cover.ImageAbs(size), nil
	}

	for _, e := range set.Els {
		switch e.(type) {
		case *entry.Image:
			if e.(*entry.Image).File.Base() == name {
				return e.(*entry.Image).ImageAbs(size), nil
			}
		case *entry.Audio:
			if e.(*entry.Audio).File.Base() == name {
				return e.(*entry.Audio).File.Path, nil
			}
		case *entry.Video:
			if e.(*entry.Video).File.Base() == name {
				return e.(*entry.Video).File.Path, nil
			}
		}
	}

	return "", fmt.Errorf("Could not find cache file (%v) in Set (%v)", name, set)
}

*/
	// serve file with subpath directly.
	/*
		if p.Type == "files" {
			f, err := entry.ElFileSafe(eh)
			if err != nil {
				s.Debug(err)
				http.NotFound(w, r)
				return
			}
			serveStatic(w, r, filepath.Join(f.Hold.File.Path, p.Descriptor))
			return
		}
	*/


package kine

import (
	"fmt"
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
	"sacer/go/server/paths"
	"strings"
	"time"
)

type kineSingle struct {
	Head	  *head.Head
	Entry	  entry.Entry
	Neighbors []entry.Entry
}

func ServeSingle(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth, p *paths.Path) {
	lang := head.Lang(r.Host)
	kine := s.Trees["kine"].Access(a.Subscriber)[lang]
	e, err := kine.LookupEntryHash(p.Hash)
	if err != nil {
		http.Redirect(w, r, "/kine", 301)
		return
	}

	perma := e.Perma(lang)
	if r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	head := &head.Head{
		Title:   getTitle(e, head.Lang(r.Host)),
		Section: "kine",
		Path:    r.URL.Path,
		Host:    r.Host,
		Entry:   e,
		Options: head.GetOptions(r),
	}

	err = head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}


	err = s.ExecuteTemplate(w, "kine-single", &kineSingle{
		Head:	   head,
		Entry:     e,
		Neighbors: getNeighbors(s.Recents["kine"].Access(a.Subscriber)[lang], p.Hash),
	})
	if err != nil {
		log.Println(err)
	}
}

func getDate(d time.Time, lang string) string {
	return fmt.Sprintf(d.Format("02 %v 2006"), tools.Abbr(tools.MonthLang(d, lang)))
}

func getTitle(e entry.Entry, lang string) string {
	return fmt.Sprintf("%v - %v - %v", e.Title(lang), getDate(e.Date(), lang), strings.Title(tools.KineName[lang]))
}

func getNeighbors(es entry.Entries, hash string) []entry.Entry {
	cpy := make([]entry.Entry, len(es))
	copy(cpy, es)
	for i, e := range cpy {
		if e.Hash() == hash {
			l := len(es)
			j, k := limits(l, i)
			d := i + 1
			if d > l {
				d = i
			}
			return append(cpy[j:i], cpy[d:k]...)
		}
	}
	return nil
}

// TODO: verbose... should be simplified
func limits(l, i int) (int, int) {
	// number of neighors left and right
	number := 2
	j, k := 0, l
	// set start position
	if x := i-number; x > 0 {
		j = x
	}

	if left := i-j; left < 2 {
		number = number + (number - left)
	}

	if y := i+1+number; y < l {
		k = y
	} else {
		k = l
	}

	if right := k-i-1; right < 2 {
		if j - (number - right) > 0 {
			j = j - (number - right)
		} else {
			println(right)
			j = 0
		}
	}

	return j, k
}

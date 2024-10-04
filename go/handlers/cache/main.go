package cache

import (
	//"fmt"
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/tools"
	"g.rg-s.com/sera/go/entry/types/tree"
	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

type cacheMain struct {
	Meta    *meta.Meta
	Tree    *tree.Tree
	Entries entry.Entries
}

func Main(w http.ResponseWriter, r *http.Request, m *meta.Meta) {

	t := server.Store.Trees["cache"].Access(m.Auth.Subscriber)[m.Lang]

	m.Title = tools.Title(tools.KineName[m.Lang])
	m.Section = "cache"
	m.Desc = t.Info().Field("description", m.Lang)

	err := m.Process(t)
	if err != nil {
		log.Println(err)
		return
	}

	entries := server.Store.Recents["cache"].Access(m.Auth.Subscriber)[m.Lang].Limit(10)

	err = server.Store.ExecuteTemplate(w, "cache-main", &cacheMain{
		Meta:    m,
		Tree:    t,
		Entries: entries,
	})
	if err != nil {
		log.Println(err)
	}
}

package front

import (
	"log"
	"net/http"
	"stferal/pkg/entry"
	"stferal/pkg/head"
	"stferal/pkg/server"
)

type frontMain struct {
	Head  *head.Head
	Index entry.Els
	Graph entry.Els
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	head := &head.Head{
		Title:   "",
		Section: "home",
		Path:    "/",
		Host:    r.Host,
		El:      nil,
		Desc:    s.Vars.Lang("site", head.Lang(r.Host)),
		Night:   head.NightMode(r),
		Large:   head.TypeMode(r),
	}
	err := head.Make()
	if err != nil {
		return
	}

	err = s.ExecuteTemplate(w, "front", &frontMain{
		Head:  head,
		Index: s.Recents["index"],
		Graph: s.Recents["graph"],
	})
	if err != nil {
		log.Println(err)
	}
}

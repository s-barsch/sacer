package head

import (
	"fmt"
	"stferal/go/entry"
)

type Head struct {
	Title   string
	Section string
	Path    string
	Host    string

	Entry   entry.Entry
	Options map[string]bool

	/*
		Local   bool
		Dark    bool
		Large   bool
		NoLog   bool
	*/

	Nav   Nav
	Lang  string
	Langs Langs

	Desc   string
	Schema *Schema
}

func (h *Head) Process() error {
	// TODO: check for nil entry?
	h.Lang = Lang(h.Host)
	h.Desc = h.GetDesc()
	h.Langs = h.MakeLangs()
	h.Nav = h.MakeNav()

	return nil
}

func (h *Head) PageTitle() string {
	if h.Title == "" {
		return "Stef Feral"
	}
	switch h.Lang {
	case "en":
		return fmt.Sprintf("%v - Stef Feral - English", h.Title)
	default:
		return fmt.Sprintf("%v - Stef Feral", h.Title)
	}
}

func (h *Head) PageURL() string {
	if l := h.Langs.Hreflang(h.Lang); l != nil {
		return l.Href
	}
	return ""
}

func (h *Head) DontIndex() bool {
	switch h.Path {
	case "/impressum", "/legal", "/privacy", "/datenschutz":
		return true
	}
	return false
}

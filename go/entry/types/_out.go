// +build ignore

{{define "out" -}}
// Code generated by go generate; DO NOT EDIT.

package {{.}}

import (
{{if not (isTree .)}}
	"fmt"
{{end}}
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

func {{receiver .}} Type() string {
	return "{{.}}"
}

func {{receiver .}} Parent() entry.Entry {
	return e.parent
}

func {{receiver .}} File() *file.File {
	return e.file
}

func {{receiver .}} Id() int64 {
	return e.date.Unix()
}

func {{receiver .}} Timestamp() string {
	return e.date.Format(helper.Timestamp)
}

func {{receiver .}} Hash() string {
	return helper.ToB16(e.date)
}

func {{receiver .}} HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func {{receiver .}} Date() time.Time {
	return e.date
}

func {{receiver .}} Info() info.Info {
	return e.info
}

func {{receiver .}} Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func {{receiver .}} Slug(lang string) string {
	if slug := e.info.Slug(lang); slug != "" {
		return slug
	}
	return helper.Normalize(e.info.Title(lang))
}

func {{receiver .}} IsBlob() bool {
	return entry.IsBlob(e)
}

func {{receiver .}} MediaType() string {
	if e.Type() == "audio" {
		return "tob"
	}
	if entry.IsBlob(e) {
		return  "blob"
	}
	return "tob"
}

func {{receiver .}} SetParent(parent entry.Entry) {
	e.parent = parent
}

func {{receiver .}} SetInfo(inf info.Info) {
	e.info = inf
}

{{if not (isMedia .)}}
func {{receiver .}} Entries() entry.Entries {
	return e.entries
}
{{end}}

{{if not (isTree .)}}
func {{receiver .}} Path(lang string) string {
	return fmt.Sprintf("%v/%v", e.parent.Path(lang), e.Slug(lang))
}

// This recursive function call will be caught by a Tree type. For now, all 
// further up parent entries are exclusively of type Tree.
func {{receiver .}} Section() string {
	return e.Parent().Section()
}

func {{receiver .}} Perma(lang string) string {
	if e.parent.Type() == "set" {
		return e.parent.Perma(lang)
	}

	name := e.Hash()
	if slug := e.Slug(lang); slug != "" {
		name = fmt.Sprintf("%v-%v", slug, e.Hash())
	}

	switch e.Section() {
	case "index":
			return fmt.Sprintf("%v#%v", e.parent.Perma(lang), helper.Normalize(e.Title(lang)))
	case "kine":
			return fmt.Sprintf(
				"/%v/%v/%v",
				helper.KineName[lang],
				e.Date().Format("06-01"),
				fmt.Sprintf("%v-%v", e.Date().Format("02"), name),
			)
	}

	return fmt.Sprintf("%v/%v", e.parent.Path(lang), name)
}
{{end}}
{{end}}

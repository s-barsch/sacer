// Code generated by go generate; DO NOT EDIT.

package audio

import (
	"fmt"

	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

func (e *Audio) Type() string {
	return "audio"
}

func (e *Audio) Parent() entry.Entry {
	return e.parent
}

func (e *Audio) File() *file.File {
	return e.file
}

func (e *Audio) Id() int64 {
	return e.date.Unix()
}

func (e *Audio) Timestamp() string {
	return e.date.Format(helper.Timestamp)
}

func (e *Audio) Hash() string {
	return helper.ToB16(e.date)
}

func (e *Audio) HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func (e *Audio) Date() time.Time {
	return e.date
}

func (e *Audio) Info() info.Info {
	return e.info
}

func (e *Audio) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Audio) Slug(lang string) string {
	if slug := e.info.Slug(lang); slug != "" {
		return slug
	}
	return helper.Normalize(e.info.Title(lang))
}

func (e *Audio) IsBlob() bool {
	return entry.IsBlob(e)
}

func (e *Audio) MediaType() string {
	if e.Type() == "audio" {
		return "tob"
	}
	if entry.IsBlob(e) {
		return "blob"
	}
	return "tob"
}

func (e *Audio) SetParent(parent entry.Entry) {
	e.parent = parent
}

func (e *Audio) SetInfo(inf info.Info) {
	e.info = inf
}

func (e *Audio) Path(lang string) string {
	return fmt.Sprintf("%v/%v", e.parent.Path(lang), e.Slug(lang))
}

// This recursive function call will be caught by a Tree type. For now, all
// further up parent entries are exclusively of type Tree.
func (e *Audio) Section() string {
	return e.Parent().Section()
}

func (e *Audio) Perma(lang string) string {
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
	case "video":
		return fmt.Sprintf(
			"/video/%v/%v",
			e.Date().Format("06-01"),
			fmt.Sprintf("%v-%v", e.Date().Format("02"), name),
		)
	}

	return fmt.Sprintf("%v/%v", e.parent.Path(lang), name)
}

// Code generated by go generate; DO NOT EDIT.

package tree

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

func (e *Tree) Type() string {
	return "tree"
}

func (e *Tree) Parent() entry.Entry {
	return e.parent
}

func (e *Tree) File() *file.File {
	return e.file
}

func (e *Tree) Id() int64 {
	return e.date.Unix()
}

func (e *Tree) Timestamp() string {
	return e.date.Format(helper.Timestamp)
}

func (e *Tree) Hash() string {
	return helper.ToB16(e.date)
}

func (e *Tree) HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func (e *Tree) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Tree) Date() time.Time {
	return e.date
}

func (e *Tree) Info() info.Info {
	return e.info
}

func (e *Tree) Slug(lang string) string {
	if slug := e.info.Slug(lang); slug != "" {
		return slug
	}
	return helper.Normalize(e.info.Title(lang))
}

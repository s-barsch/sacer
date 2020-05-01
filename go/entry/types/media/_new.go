// +build ignore

{{define "main" -}}
// Code generated by go generate; DO NOT EDIT.

package media 

import (
	"fmt"
	"stferal/go/entry"
	"stferal/go/entry/types/media/text"
)

func New(path string, parent interface{}) (*entry.Entry, error) {
	obj, err := NewEntryObject(path)
	if err != nil {
		return nil, err
	}
	return &Entry{
		Parent: parent,
		Object: obj,
	}, nil
}

func NewMediaObject(path string) (interface{}, error) {
	switch helper.FileType(path) {
{{- range $, $name := .Media}}
	case "{{$name}}":
		return {{$name}}.New(path)
{{- end}}
	}
	return nil, fmt.Errorf("invalid entry type: %v", path)
}
{{end}}

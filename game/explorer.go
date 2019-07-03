package game

import (
	"github.com/pkg/errors"
	"strings"
)

type NoSuchEntry struct{}

type Explorer struct {
	location string
	root     *folderEntry
}

type folderEntry struct {
	directories map[string]*folderEntry
	files       map[string]*FileEntry
}

type FileEntry struct {
	name       string
	size       uint
	folderName string
}

func (e *Explorer) traverse(path string, create bool) (*FileEntry, error) {
	var ok bool
	folder := e.root
	split := strings.Split(path, "/")
	for i, seg := range split {
		seg = strings.ToLower(seg)
		if i+1 == len(split) {
			return folder.files[seg], nil
		}
		if folder, ok = folder.directories[seg]; !ok {
			if create {
				folder.directories[seg] = &folderEntry{make(map[string]*folderEntry), make(map[string]*FileEntry)}
			} else {
				return nil, NoSuchEntry{}
			}
		}
	}
}

func (NoSuchEntry) Error() string {
	return "No such entry in the file tree"
}

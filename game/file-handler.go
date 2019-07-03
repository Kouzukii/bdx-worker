package game

import (
	"path"
	"strings"
)

type FileType int

const (
	DDS   FileType = 1
	PAC   FileType = 2
	Text  FileType = 3
	Other FileType = 4
)

func (e *Explorer) ExportFile(entry *FileEntry, exportTextures bool) (FileType, string) {
	ft := Other
	ext := path.Ext(entry.name)
	switch ext {
	case ".pac":
		ft = PAC
	case ".dds":
		ft = DDS
	case ".txt":
	case ".xml":
		ft = Text
	}

	filepath := path.Join(ExportDir(), entry.folderName, entry.name)
	upscaledPath := filepath[:-len(ext)] + ".upscaled" + ext

	return ft, filepath

}

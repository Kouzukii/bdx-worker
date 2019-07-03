package game

import (
	"github.com/Kouzukii/bdx-worker/util"
	"github.com/lxn/win"

	"os"
	"path"
)

type BlackDesertInstallationNotFound struct{}
type ExportDirectoryNotFound struct{}

var (
	location  string
	exportDir string
)

func Location() string {
	if location != "" {
		return location
	}

	if loc, err := util.ReadRegistryValue("SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\{C1F96C92-7B8C-485F-A9CD-37A0708A2A60}\\InstallLocation"); err == nil {
		if f, err := os.Stat(loc); err == nil && f.IsDir() {
			location = loc
			return location
		}
	}

	buf := util.NewUTF16(512)
	win.SHGetSpecialFolderPath(0, buf.Ptr(), win.CSIDL_COMMON_STARTMENU, false)
	if loc, err := util.ResolveShortcutTarget(path.Join(buf.String(), "Black Desert Online.lnk")); err == nil {
		if _, err := os.Stat(loc); err == nil {
			location = path.Dir(loc)
			return location
		}
	}

	buf = util.NewUTF16(512)
	win.SHGetSpecialFolderPath(0, buf.Ptr(), win.CSIDL_COMMON_DESKTOPDIRECTORY, false)
	if loc, err := util.ResolveShortcutTarget(path.Join(buf.String(), "Black Desert Online.lnk")); err == nil {
		if _, err := os.Stat(loc); err == nil {
			location = path.Dir(loc)
			return location
		}
	}

	buf = util.NewUTF16(512)
	win.SHGetSpecialFolderPath(0, buf.Ptr(), win.CSIDL_PROGRAM_FILESX86, false)
	loc := path.Join(buf.String(), "Black Desert Online")
	if f, err := os.Stat(loc); err == nil && f.IsDir() {
		location = loc
		return location
	}

	panic(BlackDesertInstallationNotFound{})
}

func PazLocation() string {
	return path.Join(Location(), "Paz")
}

func ExportDir() string {
	if exportDir != "" {
		return exportDir
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(ExportDirectoryNotFound{})
	}
	exportDir = path.Join(dir, "export")
	return exportDir
}

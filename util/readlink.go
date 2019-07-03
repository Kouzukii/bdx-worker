package util

// #include "shortcut.h"
import "C"

import (
	"github.com/lxn/win"
	"github.com/pkg/errors"
	"syscall"
)

func ResolveShortcutTarget(lnk string) (string, error) {
	buf := NewUTF16(512)
	if result := C.ResolveIt(0, syscall.BytePtrFromString(lnk), buf.Ptr(), 512); win.SUCCEEDED(result) {
		return buf.String(), nil
	} else {
		return "", errors.Errorf("Win32 Error %x", result)
	}
}

package util

import (
	"github.com/lxn/win"
	"github.com/pkg/errors"
	"syscall"
)

func ReadRegistryValue(key string) (string, error) {
	regkey, err := syscall.UTF16PtrFromString(key)
	if err == nil {
		buf := NewUTF16(512)
		size := buf.Sizeof()
		if result := win.HRESULT(win.RegQueryValueEx(win.HKEY_LOCAL_MACHINE, regkey, nil, nil, buf.BytePtr(), &size)); win.SUCCEEDED(result) {
			return buf.String(), nil
		} else {
			return "", errors.Errorf("Win32 error %x", result)
		}
	}
	return "", err
}

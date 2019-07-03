package util

import (
	"syscall"
	"unsafe"
)

type UTF16 struct {
	buf []uint16
}

func (b *UTF16) Ptr() *uint16 {
	return &b.buf[0]
}

func (b *UTF16) BytePtr() *byte {
	return (*byte)(unsafe.Pointer(&b.buf[0]))
}

func (b *UTF16) String() string {
	return syscall.UTF16ToString(b.buf)
}

func (b *UTF16) Sizeof() uint32 {
	return uint32(len(b.buf) * 2)
}

func NewUTF16(size uint32) *UTF16 {
	return &UTF16{make([]uint16, size)}
}

package utils

import (
	"os"
	"syscall"
)

//file struct
type File struct {
	fd   int    //file descriptor number
	name string //file name at Oprate time
}

/**
public operate method
*/
// open file
func Open(name string, mode int, perm uint32) (file *File, err error) {
	r, e := syscall.Open(name, mode, perm)
	if e != nil {
		return nil, e
	}
	return &File{r, name}, e
}

// private new file method
func newFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

/**
file's method
*/
func (file *File) Close() error {
	if file == nil {
		return os.ErrNotExist
	}
	e := syscall.Close(file.fd)
	file.fd = -1
	if e != nil {
		return e
	}
	return nil
}

func (file *File) Read(b []byte) (int, error) {
	if file == nil {
		return -1, os.ErrNotExist
	}
	r, e := syscall.Read(file.fd, b)
	return r, e
}

func (file *File) Write(b []byte) (int, error) {
	if file == nil {
		return -1, os.ErrNotExist
	}
	r, e := syscall.Write(file.fd, b)
	return r, e
}

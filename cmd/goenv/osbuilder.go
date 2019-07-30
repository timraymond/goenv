package main

import (
	"io"
	"os"
)

// OSBuilder is an implementation of the Builder interface that creates files
// and directories using functions available in the os package
type OSBuilder struct {
	DirMode  os.FileMode
	FileMode os.FileMode
}

// MkdirAll creates directories in the filesystem using the DirMode for
// permissions
func (o *OSBuilder) MkdirAll(path string) error {
	return os.MkdirAll(path, o.DirMode)
}

// Chdir changes the current working directory
func (o *OSBuilder) Chdir(path string) error {
	return os.Chdir(path)
}

// OpenFile opens a new file for writing and creates it if it doens't exist. It uses the FileMode configured in the builder struct for files that it creates.
func (o *OSBuilder) OpenFile(path string) (io.WriteCloser, error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE, o.FileMode)
}

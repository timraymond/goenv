package mocks

import "io"

type Builder struct {
	MkdirAllF func(string) error
	ChdirF    func(string) error
	OpenFileF func(string) (io.WriteCloser, error)
}

func (b *Builder) MkdirAll(path string) error {
	return b.MkdirAllF(path)
}

func (b *Builder) Chdir(path string) error {
	return b.ChdirF(path)
}

func (b *Builder) OpenFile(path string) (io.WriteCloser, error) {
	return b.OpenFileF(path)
}

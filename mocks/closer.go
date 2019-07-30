package mocks

import "io"

type WriteCloser struct {
	io.Writer
}

func (wc *WriteCloser) Close() error {
	return nil
}

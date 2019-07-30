package internal

import (
	"encoding"
	"io"
	"path"
)

type Command struct {
	Builder interface {
		MkdirAll(string) error
		Chdir(string) error
		OpenFile(string) (io.WriteCloser, error)
	}
}

// buildGopath constructs a GOPATH
func (c *Command) BuildGoPath(fullpath string) error {
	if err := c.Builder.MkdirAll(path.Join("src", fullpath)); err != nil {
		return err
	}

	if err := c.Builder.MkdirAll("bin"); err != nil {
		return err
	}

	if err := c.Builder.MkdirAll("pkg"); err != nil {
		return err
	}
	return nil
}

func (c *Command) BuildProjectPath(path string) error {
	// create the project directory
	if err := c.Builder.MkdirAll(path); err != nil {
		return err
	}

	if err := c.Builder.Chdir(path); err != nil {
		return err
	}
	return nil
}

func (c *Command) WriteConfig(path string, cfg encoding.TextMarshaler) error {
	envrc, err := c.Builder.OpenFile(path)
	if err != nil {
		return err
	}
	defer envrc.Close()

	conf, err := cfg.MarshalText()
	if err != nil {
		return err
	}

	envrc.Write(conf)
	return nil
}

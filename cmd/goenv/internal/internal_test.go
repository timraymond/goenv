package internal

import (
	"bytes"
	"io"
	"testing"

	"github.com/timraymond/goenv/direnv"
	"github.com/timraymond/goenv/mocks"
)

func Test_Command_BuildGoPath(t *testing.T) {
	got := []string{}
	exp := map[string]bool{
		"src/foo": true,
		"bin":     true,
		"pkg":     true,
	}

	c := &Command{
		Builder: &mocks.Builder{
			MkdirAllF: func(path string) error {
				got = append(got, path)
				return nil
			},
		},
	}

	err := c.BuildGoPath("foo")
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}

	if len(got) != len(exp) {
		t.Error("Length mismatch: got:", len(got), "exp:", len(exp))
	}

	for _, path := range got {
		if !exp[path] {
			t.Error("Path not expected:", path)
		}
	}
}

func Test_Command_CreateProjectPath(t *testing.T) {
	var got string
	var chdirPath string

	c := &Command{
		Builder: &mocks.Builder{
			MkdirAllF: func(path string) error {
				got = path
				return nil
			},
			ChdirF: func(path string) error {
				chdirPath = path
				return nil
			},
		},
	}

	exp := "foo"

	err := c.BuildProjectPath(exp)
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}

	if got != exp {
		t.Error("Unexpected path: got:", got, "exp:", exp)
	}

	if chdirPath != exp {
		t.Error("Unexpected chdir path: got:", chdirPath, "exp:", exp)
	}
}

func Test_Command_WriteConfig(t *testing.T) {
	var got string
	buf := bytes.NewBufferString("")

	c := &Command{
		Builder: &mocks.Builder{
			OpenFileF: func(path string) (io.WriteCloser, error) {
				got = path
				return &mocks.WriteCloser{buf}, nil
			},
		},
	}

	exp := "blah"
	err := c.WriteConfig(exp, &direnv.Config{})
	if err != nil {
		t.Fatal("Unexpected error: err:", err)
	}

	if exp != got {
		t.Error("Filenames do not match. exp:", exp, "got:", got)
	}
}

package internal

import (
	"bytes"
	"io"
	"path"
	"testing"

	"github.com/timraymond/goenv/mocks"
)

func Test_Root_Generate(t *testing.T) {
	createdDirs := map[string]bool{}
	visitedDirs := map[string]bool{}
	writtenFiles := map[string]*mocks.WriteCloser{}

	projName := "github.com/timraymond/foo"
	expCreatedDirs := []string{
		path.Base(projName),
		path.Join("src", projName),
		"bin",
		"pkg",
	}
	expVisitedDirs := []string{
		path.Base(projName),
	}

	expWrittenFiles := []string{
		".envrc",
	}

	c := &Command{
		Builder: &mocks.Builder{
			MkdirAllF: func(path string) error {
				createdDirs[path] = true
				return nil
			},
			ChdirF: func(path string) error {
				visitedDirs[path] = true
				return nil
			},
			OpenFileF: func(path string) (io.WriteCloser, error) {
				writtenFiles[path] = &mocks.WriteCloser{bytes.NewBufferString("")}
				return writtenFiles[path], nil
			},
		},
	}

	err := Run([]string{"generate", projName}, c)
	if err != nil {
		t.Fatal("Unexpected error: err:", err)
	}

	for _, exp := range expCreatedDirs {
		if !createdDirs[exp] {
			t.Error("Expected dir to be created but wasn't:", exp)
		}
	}

	for _, exp := range expVisitedDirs {
		if !visitedDirs[exp] {
			t.Error("Expected dir to be visited but wasn't:", exp)
		}
	}

	for _, exp := range expWrittenFiles {
		if _, ok := writtenFiles[exp]; !ok {
			t.Error("Expected file to be written but wasn't:", exp)
		}
	}

	if t.Failed() {
		for dir, _ := range createdDirs {
			t.Log("Created:", dir)
		}
	}
}

package internal

import (
	"os"
	"path"

	"github.com/timraymond/goenv/direnv"
)

func Run(args []string) error {
	name := path.Base(args[1])
	fullpath := "src/" + args[1]

	c := &Command{}

	c.BuildProjectPath(name)
	c.BuildGoPath(fullpath)

	// write the direnv config
	cfg := &direnv.Config{
		GoPath: true,
		Paths: []string{
			"/usr/local/goroots/go1.12/bin",
			"/usr/local/go-global/1.12/bin",
		},
		Envs: map[string]string{
			"GO111MODULE": "on",
		},
	}

	envrc, err := os.OpenFile(".envrc", os.O_RDWR|os.O_CREATE, 0644)
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

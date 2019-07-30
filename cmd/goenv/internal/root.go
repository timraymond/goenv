package internal

import (
	"path"

	"github.com/timraymond/goenv/direnv"
)

func Run(args []string, c *Command) error {
	name := path.Base(args[1])

	c.BuildProjectPath(name)
	c.BuildGoPath(args[1])

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

	c.WriteConfig(".envrc", cfg)
	return nil
}

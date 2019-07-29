package direnv

import (
	"bytes"
	"fmt"
)

// Config is a collection of information necessary for
// generating an envrc file
type Config struct {
	GoPath bool              // controls whether a GOPATH will be used
	Paths  []string          // path additions
	Envs   map[string]string // additional envs to set
}

// MarshalText writes the config as an .envrc file
func (c *Config) MarshalText() ([]byte, error) {
	out := bytes.NewBufferString("")

	for _, path := range c.Paths {
		out.WriteString("PATH_add " + path + "\n")
	}

	for k, v := range c.Envs {
		out.WriteString(fmt.Sprintf("export %s=%s\n", k, v))
	}

	if c.GoPath {
		out.WriteString("layout go\n")
	}

	out.WriteString("\n")
	return out.Bytes(), nil
}

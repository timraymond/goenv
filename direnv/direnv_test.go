package direnv_test

import (
	"testing"

	"github.com/timraymond/goenv/direnv"
)

func Test_Generate(t *testing.T) {
	genTests := []struct {
		name      string
		config    *direnv.Config
		shouldErr bool
		exp       string
	}{
		{
			"empty",
			&direnv.Config{},
			false,
			"\n",
		},
		{
			"gopath",
			&direnv.Config{
				GoPath: true,
			},
			false,
			"layout go\n\n",
		},
		{
			"go binary path",
			&direnv.Config{
				Paths: []string{
					"/usr/local/goroots/go1.12/bin",
				},
			},
			false,
			"PATH_add /usr/local/goroots/go1.12/bin\n\n",
		},
		{
			"envs",
			&direnv.Config{
				Envs: map[string]string{
					"GO111MODULE": "on",
				},
			},
			false,
			"export GO111MODULE=on\n\n",
		},
		{
			"all together now",
			&direnv.Config{
				Envs: map[string]string{
					"GO111MODULE": "on",
				},
				Paths: []string{
					"/usr/local/goroots/go1.12/bin",
				},
				GoPath: true,
			},
			false,
			"PATH_add /usr/local/goroots/go1.12/bin\nexport GO111MODULE=on\nlayout go\n\n",
		},
	}

	for _, test := range genTests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.config.MarshalText()
			if err != nil && !test.shouldErr {
				t.Fatal("Unexpected err:", err)
			}

			if err == nil && test.shouldErr {
				t.Fatal("Expected error but received none")
			}

			if string(got) != test.exp {
				t.Fatal("Config differs:\nGot:\n\n", string(got), "\n\nexp:\n", test.exp)
			}
		})
	}
}

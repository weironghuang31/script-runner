package scripts

import (
	"io"
	"os"
	"path/filepath"

	"github.com/vatine/env"
	"gopkg.in/yaml.v3"
)

type Spec struct {
	Dir     string
	Envs    map[string]string
	Scripts map[string]string
	Stderr  io.Writer
	Stdout  io.Writer
}

func ParseSpec(dir, filename string) (*Spec, error) {
	var spec = new(Spec)
	var path string

	if dir == "" {
		path = filename
	} else {
		path = filepath.Join(dir, filename)
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, spec)

	if err != nil {
		return nil, err
	}

	// expand envs
	for key, value := range spec.Envs {
		value, err = env.Expand(value)

		if err != nil {
			return nil, err
		}

		spec.Envs[key] = value
	}

	spec.Stderr = os.Stderr
	spec.Stdout = os.Stdout
	spec.Dir = dir

	return spec, nil
}

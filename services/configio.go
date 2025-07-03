package services

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"

	"github.com/your-handle/pipeai/domain"
)

const (
	dirName  = ".pipeai"
	fileName = "config.yaml"
)

func path() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, dirName, fileName), nil
}

func Load() (*domain.Config, error) {
	cfg := &domain.Config{}
	p, err := path()
	if err != nil {
		return cfg, err
	}
	f, err := os.OpenFile(p, os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	_ = yaml.NewDecoder(f).Decode(cfg) // ignore EOF on first run
	return cfg, nil
}

func Save(cfg *domain.Config) error {
	p, err := path()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0700); err != nil {
		return err
	}
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()
	e := yaml.NewEncoder(f)
	e.SetIndent(2)
	return e.Encode(cfg)
}

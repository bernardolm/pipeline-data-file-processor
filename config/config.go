package config

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// https://yaml.to-go.online/

var (
	configs          = make(map[string]Config)
	regexOnlyLetters = regexp.MustCompile(`[^a-z]`)
)

type Element struct {
	Name   string                 `yaml:"name"`
	Params map[string]interface{} `yaml:"params"`
	Type   string                 `yaml:"type"`
}

type Source Element

type Hook Element

type Output Element

type Config struct {
	Description string `yaml:"description"`
	PluginPath  string
	Hooks       []Hook   `yaml:"hooks"`
	Name        string   `yaml:"name"`
	Outputs     []Output `yaml:"outputs"`
	Sources     []Source `yaml:"sources"`
}

func pluginFilepath(path string) string {
	return fmt.Sprintf("pipes/%s/%s.so", path, regexOnlyLetters.ReplaceAllString(path, ""))
}

func loadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	if _, err := file.Read(buffer); err != nil {
		return err
	}

	var config Config

	if err := yaml.Unmarshal(buffer, &config); err != nil {
		return err
	}

	path = filepath.Base(filepath.Dir(path))
	config.PluginPath = pluginFilepath(path)
	configs[path] = config

	return nil
}

func Load() map[string]Config {
	if err := filepath.Walk("./pipes", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".yaml" {
			return nil
		}

		if matched, err := filepath.Match("config.yaml", filepath.Base(path)); err != nil {
			return err
		} else if matched {
			if err := loadFromFile(path); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		log.Panic(err)
	}

	return configs
}

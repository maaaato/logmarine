package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

const ConfigPath = "logmarine.toml"

// ConfigToml Struct
type ConfToml struct {
	TailFile      string        `toml:"TailFile"`
	PositionFile  string        `toml:"PositionFile"`
	SearchStart   string        `toml:"SearchStart"`
	SearchEnd     string        `toml:"SearchEnd"`
	TagName       string        `toml:"TagName"`
	TagStartValue string        `toml:"TagStartValue"`
	TagEndValue   string        `toml:"TagEndValue"`
	Delay         time.Duration `toml:"Delay"`
}

// New Read config.
func LoadConfig(p string) (*ConfToml, error) {
	c := new(ConfToml)
	_, err := toml.DecodeFile(p, c)
	if err != nil {
		fmt.Println(err)
	}

	// check error.
	if c.TailFile == "" {
		return c, errors.New("please set tailfile")
	}

	if c.PositionFile == "" {
		return c, errors.New("please set position file")
	}

	if c.SearchStart == "" {
		return c, errors.New("please set search start")
	}

	if c.SearchEnd == "" {
		return c, errors.New("please set search end")
	}

	return c, nil
}

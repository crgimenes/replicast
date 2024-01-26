package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"replicast/constants"
)

type Config struct {
	Debug    bool
	Listen   string
	InitFile string
	Path     string
}

var CFG = &Config{}

func Load() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	homedir := home + "/.config/replicast"
	fullpath, err := filepath.Abs(homedir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(fullpath, constants.DefaultDirMode)
	if err != nil {
		return err
	}

	CFG.Path = fullpath

	log.Println("CFG.Path:", CFG.Path)

	// Parse environment variables
	CFG.Listen = os.Getenv("REPLICAST_LISTEN")
	if CFG.Listen == "" {
		CFG.Listen = "0.0.0.0:2200"
	}
	CFG.Path = os.Getenv("COMPTERM_PATH")
	if CFG.Path == "" {
		CFG.Path = fullpath
	}
	CFG.InitFile = os.Getenv("COMPTERM_INIT_FILE")
	if CFG.InitFile == "" {
		CFG.InitFile = "init.lua"
	}

	// Parse command line flags
	flag.StringVar(&CFG.Listen, "listen", CFG.Listen, "")
	flag.StringVar(&CFG.Path, "path", CFG.Path, "")
	flag.BoolVar(&CFG.Debug, "debug", CFG.Debug, "")
	flag.StringVar(&CFG.InitFile, "init", CFG.InitFile, "")

	p := func(msg string) {
		_, _ = os.Stderr.WriteString(msg)
	}

	flag.Usage = func() { // TODO: Add more info
		p("\n")
	}

	flag.Parse()

	return err
}

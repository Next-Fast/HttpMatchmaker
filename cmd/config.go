package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/pelletier/go-toml/v2"
)

func init_config() {
	dir, err := os.Getwd();
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}

	config := path.Join(dir, "config.toml")
	_, err = os.Stat(config)
	if os.IsNotExist(err) || err != nil {
		fmt.Println("Error: %v", err)
		return
	}
}

type Config struct {
}

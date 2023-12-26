package main

import "os"

func callbackExit(cfg *Config, params ...string) error {
	os.Exit(0)
	return nil
}
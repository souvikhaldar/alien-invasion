package config

import "testing"

func TestLoadConfig(t *testing.T) {
	c := LoadConfig("../../../config.json")
	if c == nil {
		t.Fatal("Could not load config")
	}
}

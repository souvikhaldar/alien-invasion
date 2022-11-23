package config

import "testing"

func TestLoadConfig(t *testing.T) {
	c := LoadConfig("testfiles/config.json")
	if c == nil {
		t.Fatal("Could not load config")
	}
}

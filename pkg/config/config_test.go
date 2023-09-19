package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	f, _ := os.ReadFile("testdata/config_example.yml")
	t.Logf("file content: %s\n", string(f))
	config, err := LoadConfig("testdata/config_example.yml")
	if err != nil {
		t.Fatalf("load config error: %s\n", err)
	}
	server := config.Server
	if server.Host != "localhost" || server.Port != 8080 {
		t.Fatalf("not expected data")
	}
}

package configs

import (
	"testing"
)

func TestMakeConfigs(t *testing.T) {
	cfg := MakeConfigs()
	if cfg.Host == "" {
		t.Fatalf("configs must have non zero value for Host")
	}
	if cfg.Port == "" {
		t.Fatal("configs must have non zero value for Port")
	}
}

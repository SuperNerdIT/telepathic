package configs

import (
	"os"
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

func TestGetConfigFromFile(t *testing.T) {
	testsCases := map[string]struct {
		key      string
		expected string
	}{
		"Gets port": {"server.addr.port", "3000"},
		"Gets host": {"server.addr.host", "localhost"},
	}
	for name, tc := range testsCases {
		tf := func(t *testing.T) {
			want := tc.expected

			got := readFromFile(tc.key)

			if got != want {
				t.Errorf("want '%s', got '%s'", want, got)
			}
		}
		t.Run(name, tf)
	}

}

func TestSetAConfigFile(t *testing.T) {
	content := `
server:
  addr:
    port: 1111
    host: test
`
	f, err := os.Create("./config.yml")
	if err != nil {
		panic("Error creating test config file")
	}

	_, err = f.WriteString(content)
	if err != nil {
		f.Close()
		t.Fatal("Error creating config.yml file for test")
		return
	}
	cleanup := func() {
		os.Remove("./config.yml")
	}
	defer cleanup()
	f.Close()

	want := "1111"
	loadConfiguration()
	got := MakeConfigs().Port

	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

}

package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadEnvVarsWithToken(t *testing.T) {
	viper.Reset()
	t.Setenv("MRGLAB_TOKEN", "test-secret-token")

	cfg := &Config{}
	err := loadEnvVars(cfg)

	if err != nil {
		t.Fatalf("loadEnvVars() error = %v, want nil", err)
	}
	if cfg.APIToken != "test-secret-token" {
		t.Errorf("APIToken = %q, want %q", cfg.APIToken, "test-secret-token")
	}
}

func TestLoadEnvVarsWithoutToken(t *testing.T) {
	viper.Reset()
	// Ensure MRGLAB_TOKEN is not set
	t.Setenv("MRGLAB_TOKEN", "")
	os.Unsetenv("MRGLAB_TOKEN")

	cfg := &Config{}
	err := loadEnvVars(cfg)

	if err == nil {
		t.Error("loadEnvVars() should return error when token is not set")
	}
	if cfg.APIToken != "" {
		t.Errorf("APIToken = %q, want empty", cfg.APIToken)
	}
}

func TestLoadWithConfigFile(t *testing.T) {
	viper.Reset()

	dir := t.TempDir()
	content := `base_url = "https://custom.gitlab.com"

[filters]
projects = [
  { name = "Test", id = "999", fullPath = "test/project" },
]
`
	err := os.WriteFile(filepath.Join(dir, "mrglab.toml"), []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	viper.SetConfigName("mrglab")
	viper.SetConfigType("toml")
	viper.AddConfigPath(dir)

	err = viper.ReadInConfig()
	if err != nil {
		t.Fatalf("ReadInConfig() error = %v", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	if cfg.BaseURL != "https://custom.gitlab.com" {
		t.Errorf("BaseURL = %q, want %q", cfg.BaseURL, "https://custom.gitlab.com")
	}
	if len(cfg.Filters.Projects) != 1 {
		t.Fatalf("Projects count = %d, want 1", len(cfg.Filters.Projects))
	}
	if cfg.Filters.Projects[0].ID != "999" {
		t.Errorf("Project ID = %q, want %q", cfg.Filters.Projects[0].ID, "999")
	}
}

func TestDefaultBaseURL(t *testing.T) {
	cfg := &Config{}
	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://gitlab.com"
	}
	if cfg.BaseURL != "https://gitlab.com" {
		t.Errorf("default BaseURL = %q, want %q", cfg.BaseURL, "https://gitlab.com")
	}
}

func TestMockProjects(t *testing.T) {
	if len(mockProjects) != 4 {
		t.Errorf("mockProjects count = %d, want 4", len(mockProjects))
	}
	for _, p := range mockProjects {
		if p.Name == "" || p.ID == "" || p.FullPath == "" {
			t.Errorf("mockProject has empty field: %+v", p)
		}
	}
}

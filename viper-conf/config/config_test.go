package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Setup: Create a temporary config file
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
PORT: "8080"
NAME: "TestApp"
`
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("failed to create temp config file: %v", err)
	}

	// Change working directory to the temp directory
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	defer os.Chdir(originalWd)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("failed to change working directory: %v", err)
	}

	// Test LoadConfig
	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Assertions
	if config.Port != 8080 {
		t.Errorf("expected Port to be 8080, got %d", config.Port)
	}
	if config.Name != "TestApp" {
		t.Errorf("expected Name to be 'TestApp', got '%s'", config.Name)
	}
}

func TestConfig_GetPort(t *testing.T) {
	config := &Config{Port: 8080}
	if config.GetPort() != 8080 {
		t.Errorf("expected GetPort to return 8080, got %d", config.GetPort())
	}
}

func TestConfig_GetName(t *testing.T) {
	config := &Config{Name: "TestApp"}
	if config.GetName() != "TestApp" {
		t.Errorf("expected GetName to return 'TestApp', got '%s'", config.GetName())
	}
}

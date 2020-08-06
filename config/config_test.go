package config_test

import (
	"testing"

	"github.com/Perwira-Media/go-boilerplate-backend/config"
)

func TestGetServiceConfigOK(t *testing.T) {
	yamlFileLocation := "_fixtures/test-config.yaml"
	config, err := config.GetConfig(yamlFileLocation)
	if err != nil {
		t.Fatalf("It should be OK: %v", err)
	}
	if config == nil {
		t.Fatalf("It should be not nil: %v", err)
	}
}

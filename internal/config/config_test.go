package config_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yvv4git/jobs-tg-collector/internal/config"
	"github.com/yvv4git/jobs-tg-collector/internal/infrastructure/logger"
	"github.com/yvv4git/jobs-tg-collector/internal/utils"
)

func TestConfigLoad(t *testing.T) {
	const configPath = "config.toml"

	// Change directory to project root
	require.NoError(t, utils.ChangeDirToProjectRoot("../../"), "Should change directory without errors")

	// Setup logger & config
	log := logger.SetupDefaultLogger()
	cfg := config.NewConfig(log)

	// Load config
	assert.NoError(t, cfg.Load(configPath), "Should load config without errors")
	fmt.Println(cfg)
}

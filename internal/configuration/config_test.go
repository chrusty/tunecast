package configuration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	// Override one of the defaults:
	err := os.Setenv("GRPC_DISCOVERY_DNS_DOMAIN", "cruft.com")
	assert.NoError(t, err, "Couldn't set an env var")

	// Load config from env-vars:
	newConfig, err := Load()
	assert.NoError(t, err, "Error while loading config")

	// Check that default options are being picked up:
	assert.Equal(t, newConfig.GRPC.DiscoveryDefaultPort, 3000, "Default config value wasn't loaded")

	// Check that we can override the defaults:
	assert.Equal(t, newConfig.GRPC.DiscoveryDNSDomain, "cruft.com", "Default config value wasn't overridden")
}

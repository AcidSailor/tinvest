package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithAppName(t *testing.T) {
	cfg := connConfig{}
	WithAppName("myapp")(&cfg)

	assert.Equal(t, "myapp", cfg.appName)
}

func TestWithAppName_EmptyClearsAppName(t *testing.T) {
	cfg := connConfig{appName: "preset"}
	WithAppName("")(&cfg)

	assert.Equal(t, "", cfg.appName)
}

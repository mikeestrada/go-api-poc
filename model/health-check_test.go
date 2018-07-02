package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestHealthCheck(t *testing.T) {
	healthCheck := HealthCheck{}
	actual := healthCheck.status()

	expected := HealthCheck{Alive: true}
	assert.Equal(t, expected.Alive, actual.Alive)
}

func TestHealthCheckFalse(t *testing.T) {
	healthCheck := HealthCheck{}
	actual := healthCheck.status()

	expected := HealthCheck{Alive: false}
	assert.NotEqual(t, expected.Alive, actual.Alive)
}

func TestHealthCheckJson(t *testing.T) {
	healthCheck := HealthCheck{}
	actual, _ := json.Marshal(healthCheck.status())

	expected, _ := json.Marshal(HealthCheck{Alive: true})

	assert.JSONEq(t, string(expected), string(actual))
}
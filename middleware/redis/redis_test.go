package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRedis(t *testing.T) {
	err := InitRedis()
	assert.NoError(t, err)
}

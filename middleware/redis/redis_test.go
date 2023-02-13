package redis

import (
	"SimpleDouyin/middleware/DBUtils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRedis(t *testing.T) {
	err := DBUtils.InitRedis()
	assert.NoError(t, err)
}

package dateutils_test

import (
	"testing"
	"time"

	"github.com/Goldziher/go-utils/dateutils"
	"github.com/stretchr/testify/assert"
)

func TestFloor(t *testing.T) {
	now := time.Now()
	assert.Equal(t, "00:00:00", dateutils.Floor(now).Format("15:04:05"))
}

func TestCeil(t *testing.T) {
	now := time.Now()
	assert.Equal(t, "23:59:59", dateutils.Ceil(now).Format("15:04:05"))
}

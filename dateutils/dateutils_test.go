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

func TestBeforeOrEqual(t *testing.T) {
	milestone, _ := time.Parse("2006-01-02", "2023-01-01")

	dBefore, _ := time.Parse("2006-01-02", "2022-12-31")
	dEqual, _ := time.Parse("2006-01-02", "2023-01-01")
	dAfter, _ := time.Parse("2006-01-02", "2023-01-31")

	assert.Equal(t, true, dateutils.BeforeOrEqual(milestone, dBefore))
	assert.Equal(t, true, dateutils.BeforeOrEqual(milestone, dEqual))
	assert.Equal(t, false, dateutils.BeforeOrEqual(milestone, dAfter))
}

func TestAfterOrEqual(t *testing.T) {
	milestone, _ := time.Parse("2006-01-02", "2023-01-01")

	dBefore, _ := time.Parse("2006-01-02", "2022-12-31")
	dEqual, _ := time.Parse("2006-01-02", "2023-01-01")
	dAfter, _ := time.Parse("2006-01-02", "2023-01-31")

	assert.Equal(t, false, dateutils.AfterOrEqual(milestone, dBefore))
	assert.Equal(t, true, dateutils.AfterOrEqual(milestone, dEqual))
	assert.Equal(t, true, dateutils.AfterOrEqual(milestone, dAfter))
}

func TestOverlap(t *testing.T) {
	s1, _ := time.Parse("2006-01-02", "2022-12-28")
	e1, _ := time.Parse("2006-01-02", "2022-12-31")

	s2, _ := time.Parse("2006-01-02", "2022-12-30")
	e2, _ := time.Parse("2006-01-02", "2023-01-01")

	s3, _ := time.Parse("2006-01-02", "2023-01-02")
	e3, _ := time.Parse("2006-01-02", "2023-01-04")

	assert.Equal(t, true, dateutils.Overlap(s1, e1, s2, e2))
	assert.Equal(t, false, dateutils.Overlap(s1, e1, s3, e3))

	s4, _ := time.Parse("2006-01-02", "2023-07-13")
	e4, _ := time.Parse("2006-01-02", "2023-07-14")

	s5, _ := time.Parse("2006-01-02", "2023-07-10")
	e5, _ := time.Parse("2006-01-02", "2023-07-17")

	assert.Equal(t, true, dateutils.Overlap(s4, e4, s5, e5))
	assert.Equal(t, true, dateutils.Overlap(s5, e5, s4, e4))
}

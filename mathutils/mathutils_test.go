package mathutils_test

import (
	"testing"

	"github.com/Goldziher/go-utils/mathutils"
	"github.com/stretchr/testify/assert"
)

func TestClamp(t *testing.T) {
	assert.Equal(t, 5, mathutils.Clamp(5, 0, 10))
	assert.Equal(t, 0, mathutils.Clamp(-5, 0, 10))
	assert.Equal(t, 10, mathutils.Clamp(15, 0, 10))
	assert.Equal(t, 2.5, mathutils.Clamp(2.5, 0.0, 5.0))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 5, mathutils.Abs(5))
	assert.Equal(t, 5, mathutils.Abs(-5))
	assert.Equal(t, 0, mathutils.Abs(0))
	assert.Equal(t, 3.5, mathutils.Abs(-3.5))
	assert.Equal(t, uint(5), mathutils.Abs(uint(5)))
}

func TestInRange(t *testing.T) {
	assert.True(t, mathutils.InRange(5, 0, 10))
	assert.True(t, mathutils.InRange(0, 0, 10))
	assert.True(t, mathutils.InRange(10, 0, 10))
	assert.False(t, mathutils.InRange(-1, 0, 10))
	assert.False(t, mathutils.InRange(11, 0, 10))
}

func TestAverage(t *testing.T) {
	assert.Equal(t, 3.0, mathutils.Average([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0.0, mathutils.Average([]int{}))
	assert.Equal(t, 5.0, mathutils.Average([]float64{2.5, 7.5}))
	assert.Equal(t, -2.5, mathutils.Average([]int{-2, -3}))
}

func TestGcd(t *testing.T) {
	assert.Equal(t, 6, mathutils.Gcd(48, 18))
	assert.Equal(t, 1, mathutils.Gcd(17, 13))
	assert.Equal(t, 5, mathutils.Gcd(10, 5))
	assert.Equal(t, 12, mathutils.Gcd(-12, 24))
	assert.Equal(t, 7, mathutils.Gcd(7, 0))
}

func TestLcm(t *testing.T) {
	assert.Equal(t, 12, mathutils.Lcm(4, 6))
	assert.Equal(t, 221, mathutils.Lcm(17, 13))
	assert.Equal(t, 10, mathutils.Lcm(10, 5))
	assert.Equal(t, 0, mathutils.Lcm(0, 5))
	assert.Equal(t, 0, mathutils.Lcm(5, 0))
}

func TestIsPrime(t *testing.T) {
	assert.True(t, mathutils.IsPrime(2))
	assert.True(t, mathutils.IsPrime(3))
	assert.True(t, mathutils.IsPrime(5))
	assert.True(t, mathutils.IsPrime(7))
	assert.True(t, mathutils.IsPrime(11))
	assert.True(t, mathutils.IsPrime(13))
	assert.True(t, mathutils.IsPrime(17))
	assert.True(t, mathutils.IsPrime(19))

	assert.False(t, mathutils.IsPrime(0))
	assert.False(t, mathutils.IsPrime(1))
	assert.False(t, mathutils.IsPrime(4))
	assert.False(t, mathutils.IsPrime(6))
	assert.False(t, mathutils.IsPrime(8))
	assert.False(t, mathutils.IsPrime(9))
	assert.False(t, mathutils.IsPrime(10))
	assert.False(t, mathutils.IsPrime(15))
	assert.False(t, mathutils.IsPrime(21))
}

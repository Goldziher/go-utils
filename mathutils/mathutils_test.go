package mathutils_test

import (
	"testing"

	"github.com/Goldziher/go-utils/mathutils"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, mathutils.Min(1, 2))
	assert.Equal(t, 1, mathutils.Min(2, 1))
	assert.Equal(t, -5, mathutils.Min(-5, -2))
	assert.Equal(t, 1.5, mathutils.Min(1.5, 2.5))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 2, mathutils.Max(1, 2))
	assert.Equal(t, 2, mathutils.Max(2, 1))
	assert.Equal(t, -2, mathutils.Max(-5, -2))
	assert.Equal(t, 2.5, mathutils.Max(1.5, 2.5))
}

func TestMinN(t *testing.T) {
	assert.Equal(t, 1, mathutils.MinN([]int{3, 1, 4, 1, 5, 9, 2, 6}))
	assert.Equal(t, -10, mathutils.MinN([]int{5, -10, 3, 8}))
	assert.Equal(t, 1.5, mathutils.MinN([]float64{3.5, 1.5, 2.5}))
	assert.Equal(t, 0, mathutils.MinN([]int{}))
}

func TestMaxN(t *testing.T) {
	assert.Equal(t, 9, mathutils.MaxN([]int{3, 1, 4, 1, 5, 9, 2, 6}))
	assert.Equal(t, 8, mathutils.MaxN([]int{5, -10, 3, 8}))
	assert.Equal(t, 3.5, mathutils.MaxN([]float64{3.5, 1.5, 2.5}))
	assert.Equal(t, 0, mathutils.MaxN([]int{}))
}

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

func TestSum(t *testing.T) {
	assert.Equal(t, 15, mathutils.Sum([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, mathutils.Sum([]int{}))
	assert.Equal(t, -5, mathutils.Sum([]int{-2, -3}))
	assert.Equal(t, 7.5, mathutils.Sum([]float64{2.5, 5.0}))
}

func TestProduct(t *testing.T) {
	assert.Equal(t, 120, mathutils.Product([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 1, mathutils.Product([]int{}))
	assert.Equal(t, 6, mathutils.Product([]int{-2, -3}))
	assert.Equal(t, 12.5, mathutils.Product([]float64{2.5, 5.0}))
}

func TestAverage(t *testing.T) {
	assert.Equal(t, 3.0, mathutils.Average([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0.0, mathutils.Average([]int{}))
	assert.Equal(t, 5.0, mathutils.Average([]float64{2.5, 7.5}))
	assert.Equal(t, -2.5, mathutils.Average([]int{-2, -3}))
}

func TestRound(t *testing.T) {
	assert.Equal(t, 3.0, mathutils.Round(3.4))
	assert.Equal(t, 4.0, mathutils.Round(3.5))
	assert.Equal(t, -3.0, mathutils.Round(-3.4))
	assert.Equal(t, -4.0, mathutils.Round(-3.5))
}

func TestRoundToDecimal(t *testing.T) {
	assert.Equal(t, 3.14, mathutils.RoundToDecimal(3.14159, 2))
	assert.Equal(t, 3.142, mathutils.RoundToDecimal(3.14159, 3))
	assert.Equal(t, 3.0, mathutils.RoundToDecimal(3.14159, 0))
	assert.Equal(t, 10.0, mathutils.RoundToDecimal(9.999, 0))
}

func TestFloor(t *testing.T) {
	assert.Equal(t, 3.0, mathutils.Floor(3.9))
	assert.Equal(t, 3.0, mathutils.Floor(3.1))
	assert.Equal(t, -4.0, mathutils.Floor(-3.1))
}

func TestCeil(t *testing.T) {
	assert.Equal(t, 4.0, mathutils.Ceil(3.1))
	assert.Equal(t, 4.0, mathutils.Ceil(3.9))
	assert.Equal(t, -3.0, mathutils.Ceil(-3.1))
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

func TestPow(t *testing.T) {
	assert.Equal(t, 8, mathutils.Pow(2, 3))
	assert.Equal(t, 1, mathutils.Pow(5, 0))
	assert.Equal(t, 5, mathutils.Pow(5, 1))
	assert.Equal(t, 1000, mathutils.Pow(10, 3))
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

func TestIsEven(t *testing.T) {
	assert.True(t, mathutils.IsEven(0))
	assert.True(t, mathutils.IsEven(2))
	assert.True(t, mathutils.IsEven(4))
	assert.True(t, mathutils.IsEven(-2))

	assert.False(t, mathutils.IsEven(1))
	assert.False(t, mathutils.IsEven(3))
	assert.False(t, mathutils.IsEven(-1))
}

func TestIsOdd(t *testing.T) {
	assert.True(t, mathutils.IsOdd(1))
	assert.True(t, mathutils.IsOdd(3))
	assert.True(t, mathutils.IsOdd(-1))

	assert.False(t, mathutils.IsOdd(0))
	assert.False(t, mathutils.IsOdd(2))
	assert.False(t, mathutils.IsOdd(4))
	assert.False(t, mathutils.IsOdd(-2))
}

func TestSign(t *testing.T) {
	assert.Equal(t, 1, mathutils.Sign(5))
	assert.Equal(t, 1, mathutils.Sign(0.1))
	assert.Equal(t, -1, mathutils.Sign(-5))
	assert.Equal(t, -1, mathutils.Sign(-0.1))
	assert.Equal(t, 0, mathutils.Sign(0))
	assert.Equal(t, 0, mathutils.Sign(0.0))
}

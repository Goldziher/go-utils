package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTo(t *testing.T) {
	t.Run("should return pointer to int", func(t *testing.T) {
		value := 42
		ptr := To(value)
		assert.NotNil(t, ptr)
		assert.Equal(t, value, *ptr)
	})

	t.Run("should return pointer to string", func(t *testing.T) {
		value := "hello"
		ptr := To(value)
		assert.NotNil(t, ptr)
		assert.Equal(t, value, *ptr)
	})

	t.Run("should return pointer to struct", func(t *testing.T) {
		type TestStruct struct {
			Name string
			Age  int
		}
		value := TestStruct{Name: "Alice", Age: 30}
		ptr := To(value)
		assert.NotNil(t, ptr)
		assert.Equal(t, value, *ptr)
	})

	t.Run("should return pointer to zero value", func(t *testing.T) {
		var value int
		ptr := To(value)
		assert.NotNil(t, ptr)
		assert.Equal(t, 0, *ptr)
	})

	t.Run("should allow creating pointer to literal", func(t *testing.T) {
		ptr := To(100)
		assert.NotNil(t, ptr)
		assert.Equal(t, 100, *ptr)
	})
}

func TestDeref(t *testing.T) {
	t.Run("should return value when pointer is not nil", func(t *testing.T) {
		value := 42
		ptr := &value
		result := Deref(ptr, 0)
		assert.Equal(t, 42, result)
	})

	t.Run("should return default when pointer is nil", func(t *testing.T) {
		var ptr *int
		result := Deref(ptr, 99)
		assert.Equal(t, 99, result)
	})

	t.Run("should work with strings", func(t *testing.T) {
		value := "hello"
		ptr := &value
		result := Deref(ptr, "default")
		assert.Equal(t, "hello", result)

		var nilPtr *string
		result = Deref(nilPtr, "default")
		assert.Equal(t, "default", result)
	})

	t.Run("should work with structs", func(t *testing.T) {
		type TestStruct struct {
			Name string
		}
		value := TestStruct{Name: "Alice"}
		ptr := &value
		defaultValue := TestStruct{Name: "Default"}

		result := Deref(ptr, defaultValue)
		assert.Equal(t, "Alice", result.Name)

		var nilPtr *TestStruct
		result = Deref(nilPtr, defaultValue)
		assert.Equal(t, "Default", result.Name)
	})

	t.Run("should work with zero values", func(t *testing.T) {
		zero := 0
		ptr := &zero
		result := Deref(ptr, 42)
		assert.Equal(t, 0, result)
	})
}

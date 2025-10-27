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

func TestEqual(t *testing.T) {
	t.Run("should return true when both are nil", func(t *testing.T) {
		var a, b *int
		assert.True(t, Equal(a, b))
	})

	t.Run("should return false when one is nil", func(t *testing.T) {
		value := 42
		ptr := &value
		var nilPtr *int
		assert.False(t, Equal(ptr, nilPtr))
		assert.False(t, Equal(nilPtr, ptr))
	})

	t.Run("should return true when both point to equal values", func(t *testing.T) {
		a := 42
		b := 42
		assert.True(t, Equal(&a, &b))
	})

	t.Run("should return false when values differ", func(t *testing.T) {
		a := 42
		b := 43
		assert.False(t, Equal(&a, &b))
	})

	t.Run("should work with strings", func(t *testing.T) {
		a := "hello"
		b := "hello"
		c := "world"
		assert.True(t, Equal(&a, &b))
		assert.False(t, Equal(&a, &c))
	})
}

func TestCoalesce(t *testing.T) {
	t.Run("should return first non-nil pointer", func(t *testing.T) {
		var a, b *int
		c := 42
		d := 99
		result := Coalesce(a, b, &c, &d)
		assert.NotNil(t, result)
		assert.Equal(t, 42, *result)
	})

	t.Run("should return nil when all are nil", func(t *testing.T) {
		var a, b, c *int
		result := Coalesce(a, b, c)
		assert.Nil(t, result)
	})

	t.Run("should return first when first is non-nil", func(t *testing.T) {
		a := 1
		b := 2
		result := Coalesce(&a, &b)
		assert.NotNil(t, result)
		assert.Equal(t, 1, *result)
	})

	t.Run("should work with no arguments", func(t *testing.T) {
		result := Coalesce[int]()
		assert.Nil(t, result)
	})
}

func TestIsNil(t *testing.T) {
	t.Run("should return true for nil pointer", func(t *testing.T) {
		var ptr *int
		assert.True(t, IsNil(ptr))
	})

	t.Run("should return false for non-nil pointer", func(t *testing.T) {
		value := 42
		assert.False(t, IsNil(&value))
	})
}

func TestToSlice(t *testing.T) {
	t.Run("should return empty slice for nil", func(t *testing.T) {
		var ptr *int
		result := ToSlice(ptr)
		assert.Empty(t, result)
		assert.NotNil(t, result)
	})

	t.Run("should return slice with single element", func(t *testing.T) {
		value := 42
		result := ToSlice(&value)
		assert.Equal(t, []int{42}, result)
	})

	t.Run("should work with strings", func(t *testing.T) {
		value := "hello"
		result := ToSlice(&value)
		assert.Equal(t, []string{"hello"}, result)
	})
}

func TestDerefSlice(t *testing.T) {
	t.Run("should dereference all non-nil pointers", func(t *testing.T) {
		a := 1
		b := 2
		c := 3
		ptrs := []*int{&a, &b, &c}
		result := DerefSlice(ptrs, 0)
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("should use default for nil pointers", func(t *testing.T) {
		a := 1
		var b *int
		c := 3
		ptrs := []*int{&a, b, &c}
		result := DerefSlice(ptrs, 99)
		assert.Equal(t, []int{1, 99, 3}, result)
	})

	t.Run("should handle empty slice", func(t *testing.T) {
		var ptrs []*int
		result := DerefSlice(ptrs, 0)
		assert.Empty(t, result)
		assert.NotNil(t, result)
	})

	t.Run("should work with strings", func(t *testing.T) {
		a := "hello"
		var b *string
		c := "world"
		ptrs := []*string{&a, b, &c}
		result := DerefSlice(ptrs, "default")
		assert.Equal(t, []string{"hello", "default", "world"}, result)
	})
}

func TestNonNilSlice(t *testing.T) {
	t.Run("should filter out nil pointers", func(t *testing.T) {
		a := 1
		var b *int
		c := 3
		var d *int
		e := 5
		ptrs := []*int{&a, b, &c, d, &e}
		result := NonNilSlice(ptrs)
		assert.Len(t, result, 3)
		assert.Equal(t, 1, *result[0])
		assert.Equal(t, 3, *result[1])
		assert.Equal(t, 5, *result[2])
	})

	t.Run("should return empty slice when all are nil", func(t *testing.T) {
		var a, b, c *int
		ptrs := []*int{a, b, c}
		result := NonNilSlice(ptrs)
		assert.Empty(t, result)
		assert.NotNil(t, result)
	})

	t.Run("should return all when none are nil", func(t *testing.T) {
		a := 1
		b := 2
		c := 3
		ptrs := []*int{&a, &b, &c}
		result := NonNilSlice(ptrs)
		assert.Len(t, result, 3)
	})

	t.Run("should handle empty slice", func(t *testing.T) {
		var ptrs []*int
		result := NonNilSlice(ptrs)
		assert.Empty(t, result)
		assert.NotNil(t, result)
	})
}

func TestToPointerSlice(t *testing.T) {
	t.Run("should convert all values to pointers", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}
		result := ToPointerSlice(values)
		assert.Len(t, result, 5)
		for i, ptr := range result {
			assert.NotNil(t, ptr)
			assert.Equal(t, values[i], *ptr)
		}
	})

	t.Run("should handle empty slice", func(t *testing.T) {
		var values []int
		result := ToPointerSlice(values)
		assert.Empty(t, result)
		assert.NotNil(t, result)
	})

	t.Run("should work with strings", func(t *testing.T) {
		values := []string{"hello", "world"}
		result := ToPointerSlice(values)
		assert.Len(t, result, 2)
		assert.Equal(t, "hello", *result[0])
		assert.Equal(t, "world", *result[1])
	})

	t.Run("should work with zero values", func(t *testing.T) {
		values := []int{0, 0, 0}
		result := ToPointerSlice(values)
		assert.Len(t, result, 3)
		for _, ptr := range result {
			assert.NotNil(t, ptr)
			assert.Equal(t, 0, *ptr)
		}
	})
}

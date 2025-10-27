package exc

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustResult(t *testing.T) {
	t.Run("should return value when error is nil", func(t *testing.T) {
		result := MustResult(42, nil)
		assert.Equal(t, 42, result)
	})

	t.Run("should panic when error is not nil", func(t *testing.T) {
		assert.Panics(t, func() {
			MustResult(0, errors.New("test error"))
		})
	})

	t.Run("should panic with custom message", func(t *testing.T) {
		assert.PanicsWithError(t, "custom message: test error", func() {
			MustResult(0, errors.New("test error"), "custom message")
		})
	})

	t.Run("should panic with multiple messages including empty", func(t *testing.T) {
		assert.PanicsWithError(t, "first message third message: test error", func() {
			MustResult(0, errors.New("test error"), "first message", "", "third message")
		})
	})

	t.Run("should work with different types", func(t *testing.T) {
		result := MustResult("hello", nil)
		assert.Equal(t, "hello", result)

		type testStruct struct {
			Name string
		}
		s := testStruct{Name: "test"}
		result2 := MustResult(s, nil)
		assert.Equal(t, "test", result2.Name)
	})
}

func TestMust(t *testing.T) {
	t.Run("should not panic when error is nil", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Must(nil)
		})
	})

	t.Run("should panic when error is not nil", func(t *testing.T) {
		assert.Panics(t, func() {
			Must(errors.New("test error"))
		})
	})

	t.Run("should panic with custom message", func(t *testing.T) {
		assert.PanicsWithError(t, "operation failed: test error", func() {
			Must(errors.New("test error"), "operation failed")
		})
	})
}

func TestTry(t *testing.T) {
	t.Run("should return error from function", func(t *testing.T) {
		err := Try(func() error {
			return errors.New("test error")
		})
		assert.Error(t, err)
		assert.Equal(t, "test error", err.Error())
	})

	t.Run("should return nil when function succeeds", func(t *testing.T) {
		err := Try(func() error {
			return nil
		})
		assert.NoError(t, err)
	})

	t.Run("should catch panic and convert to error", func(t *testing.T) {
		err := Try(func() error {
			panic(errors.New("panic error"))
		})
		assert.Error(t, err)
		assert.Equal(t, "panic error", err.Error())
	})

	t.Run("should catch non-error panic", func(t *testing.T) {
		err := Try(func() error {
			panic("string panic")
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "string panic")
	})
}

func TestCatch(t *testing.T) {
	t.Run("should return error from function", func(t *testing.T) {
		err := Catch(func() error {
			return errors.New("test error")
		})
		assert.Error(t, err)
		assert.Equal(t, "test error", err.Error())
	})

	t.Run("should return nil when function succeeds", func(t *testing.T) {
		err := Catch(func() error {
			return nil
		})
		assert.NoError(t, err)
	})

	t.Run("should catch panic and convert to error", func(t *testing.T) {
		err := Catch(func() error {
			panic(errors.New("panic error"))
		})
		assert.Error(t, err)
		assert.Equal(t, "panic error", err.Error())
	})

	t.Run("should catch non-error panic", func(t *testing.T) {
		err := Catch(func() error {
			panic("string panic")
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "panic: string panic")
	})
}

func TestIgnoreErr(t *testing.T) {
	t.Run("should ignore error from function", func(t *testing.T) {
		called := false
		assert.NotPanics(t, func() {
			IgnoreErr(func() error {
				called = true
				return errors.New("ignored error")
			})
		})
		assert.True(t, called)
	})

	t.Run("should not panic on nil error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			IgnoreErr(func() error {
				return nil
			})
		})
	})
}

func TestFirstErr(t *testing.T) {
	t.Run("should return first non-nil error", func(t *testing.T) {
		err1 := errors.New("error 1")
		err2 := errors.New("error 2")
		result := FirstErr(nil, nil, err1, err2)
		assert.Equal(t, err1, result)
	})

	t.Run("should return nil when all errors are nil", func(t *testing.T) {
		result := FirstErr(nil, nil, nil)
		assert.NoError(t, result)
	})

	t.Run("should return nil when no errors provided", func(t *testing.T) {
		result := FirstErr()
		assert.NoError(t, result)
	})

	t.Run("should return error when only one provided", func(t *testing.T) {
		err := errors.New("single error")
		result := FirstErr(err)
		assert.Equal(t, err, result)
	})
}

func TestReturnAnyErr(t *testing.T) {
	t.Run("should return first error in values", func(t *testing.T) {
		err := errors.New("test error")
		result := ReturnAnyErr("string", 42, err, "other")
		assert.Equal(t, err, result)
	})

	t.Run("should return nil when no errors", func(t *testing.T) {
		result := ReturnAnyErr("string", 42, "other")
		assert.NoError(t, result)
	})

	t.Run("should return nil when empty", func(t *testing.T) {
		result := ReturnAnyErr()
		assert.NoError(t, result)
	})

	t.Run("should skip nil errors", func(t *testing.T) {
		var nilErr error
		err := errors.New("real error")
		result := ReturnAnyErr(nilErr, "test", err)
		assert.Equal(t, err, result)
	})
}

func TestReturnNotNil(t *testing.T) {
	t.Run("should return value when not nil", func(t *testing.T) {
		value := 42
		ptr := &value
		result := ReturnNotNil(ptr)
		assert.Equal(t, ptr, result)
		assert.Equal(t, 42, *result)
	})

	t.Run("should panic when value is nil", func(t *testing.T) {
		var ptr *int
		assert.Panics(t, func() {
			ReturnNotNil(ptr)
		})
	})

	t.Run("should panic with custom message", func(t *testing.T) {
		var ptr *string
		assert.Panics(t, func() {
			ReturnNotNil(ptr, "value must not be nil")
		})
	})
}

func TestMustNotNil(t *testing.T) {
	t.Run("should return dereferenced value when not nil", func(t *testing.T) {
		value := 42
		ptr := &value
		result := MustNotNil(ptr)
		assert.Equal(t, 42, result)
	})

	t.Run("should panic when value is nil", func(t *testing.T) {
		var ptr *int
		assert.Panics(t, func() {
			MustNotNil(ptr)
		})
	})

	t.Run("should panic with custom message", func(t *testing.T) {
		var ptr *string
		assert.Panics(t, func() {
			MustNotNil(ptr, "required field missing")
		})
	})

	t.Run("should work with structs", func(t *testing.T) {
		type TestStruct struct {
			Name string
		}
		value := TestStruct{Name: "test"}
		ptr := &value
		result := MustNotNil(ptr)
		assert.Equal(t, "test", result.Name)
	})
}

func TestAllErr(t *testing.T) {
	t.Run("should return joined errors", func(t *testing.T) {
		err1 := errors.New("error 1")
		err2 := errors.New("error 2")
		err3 := errors.New("error 3")
		result := AllErr(err1, err2, err3)
		assert.Error(t, result)
		assert.Contains(t, result.Error(), "error 1")
		assert.Contains(t, result.Error(), "error 2")
		assert.Contains(t, result.Error(), "error 3")
	})

	t.Run("should return nil when all errors are nil", func(t *testing.T) {
		result := AllErr(nil, nil, nil)
		assert.NoError(t, result)
	})

	t.Run("should skip nil errors", func(t *testing.T) {
		err1 := errors.New("error 1")
		err2 := errors.New("error 2")
		result := AllErr(nil, err1, nil, err2, nil)
		assert.Error(t, result)
		assert.Contains(t, result.Error(), "error 1")
		assert.Contains(t, result.Error(), "error 2")
	})

	t.Run("should return error when only one provided", func(t *testing.T) {
		err := errors.New("single error")
		result := AllErr(err)
		assert.Error(t, result)
		assert.Contains(t, result.Error(), "single error")
	})
}

package structutils_test

import (
	"reflect"
	"testing"

	"github.com/Goldziher/go-utils/structutils"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	First  string
	Second int
	Third  bool `struct:"third"`
}

func TestForEach(t *testing.T) {
	instance := TestStruct{
		"moishe",
		22,
		true,
	}
	structutils.ForEach(instance, func(key string, value any, tag reflect.StructTag) {
		switch key {
		case "First":
			assert.Equal(t, "moishe", value.(string))
			assert.Zero(t, tag)
		case "Second":
			assert.Equal(t, 22, value.(int))
			assert.Zero(t, tag)
		case "Third":
			assert.Equal(t, true, value.(bool))
			assert.Equal(t, "third", tag.Get("struct"))
		}
	})
}

func TestToMap(t *testing.T) {
	instance := TestStruct{
		"moishe",
		22,
		true,
	}
	assert.Equal(t, map[string]any{
		"First":  "moishe",
		"Second": 22,
		"Third":  true,
	}, structutils.ToMap(instance))

	assert.Equal(t, map[string]any{
		"First":  "moishe",
		"Second": 22,
		"third":  true,
	}, structutils.ToMap(instance, "struct"))
}

func TestFields(t *testing.T) {
	instance := TestStruct{
		"moishe",
		22,
		true,
	}
	fields := structutils.Fields(instance)
	assert.Len(t, fields, 3)
	assert.Contains(t, fields, "First")
	assert.Contains(t, fields, "Second")
	assert.Contains(t, fields, "Third")
}

func TestValues(t *testing.T) {
	instance := TestStruct{
		"moishe",
		22,
		true,
	}
	values := structutils.Values(instance)
	assert.Len(t, values, 3)
	assert.Contains(t, values, "moishe")
	assert.Contains(t, values, 22)
	assert.Contains(t, values, true)
}

func TestHasField(t *testing.T) {
	instance := TestStruct{
		"moishe",
		22,
		true,
	}
	assert.True(t, structutils.HasField(instance, "First"))
	assert.True(t, structutils.HasField(instance, "Second"))
	assert.True(t, structutils.HasField(instance, "Third"))
	assert.False(t, structutils.HasField(instance, "Nonexistent"))
}

func TestGetField(t *testing.T) {
	instance := TestStruct{
		"moishe",
		22,
		true,
	}

	value, found := structutils.GetField(instance, "First")
	assert.True(t, found)
	assert.Equal(t, "moishe", value)

	value, found = structutils.GetField(instance, "Second")
	assert.True(t, found)
	assert.Equal(t, 22, value)

	value, found = structutils.GetField(instance, "Third")
	assert.True(t, found)
	assert.Equal(t, true, value)

	value, found = structutils.GetField(instance, "Nonexistent")
	assert.False(t, found)
	assert.Nil(t, value)
}

func TestFieldNames(t *testing.T) {
	instance := TestStruct{
		"moishe",
		22,
		true,
	}

	// Without tags
	names := structutils.FieldNames(instance)
	assert.Len(t, names, 3)
	assert.Contains(t, names, "First")
	assert.Contains(t, names, "Second")
	assert.Contains(t, names, "Third")

	// With tags
	names = structutils.FieldNames(instance, "struct")
	assert.Len(t, names, 3)
	assert.Contains(t, names, "First")
	assert.Contains(t, names, "Second")
	assert.Contains(t, names, "third") // Uses tag value
}

func TestFieldNamesWithOmit(t *testing.T) {
	type OmitStruct struct {
		Keep   string
		Omit   string `json:"-"`
		Rename string `json:"renamed"`
	}

	instance := OmitStruct{
		Keep:   "value",
		Omit:   "omitted",
		Rename: "renamed_value",
	}

	names := structutils.FieldNames(instance, "json")
	assert.Len(t, names, 2)
	assert.Contains(t, names, "Keep")
	assert.Contains(t, names, "renamed")
	assert.NotContains(t, names, "Omit")
}

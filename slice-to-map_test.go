package slicetomap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToMap_ReturnsExpectedStringData(t *testing.T) {
	t.Parallel()
	// Arrange
	input := []string{"a", "b", "c"}
	compare := func(v string) string {
		return v
	}

	// Act
	result := ToMap(input, compare)

	// Assert
	expected := map[string]string{"a": "a", "b": "b", "c": "c"}

	assert.Equal(t, expected, result)
}

func TestToMap_ReturnsExpectedStructData(t *testing.T) {
	t.Parallel()
	// Arrange
	type TestStruct struct {
		Name   string
		Number int
	}

	input := []*TestStruct{
		{Name: "a", Number: 1},
		{Name: "b", Number: 2},
		{Name: "c", Number: 3},
	}
	compare := func(v *TestStruct) string {
		return v.Name
	}

	// Act
	result := ToMap(input, compare)

	// Assert
	expected := map[string]*TestStruct{
		"a": input[0],
		"b": input[1],
		"c": input[2],
	}

	assert.Equal(t, expected, result)
}

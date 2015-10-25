package jdog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopLevelMap(t *testing.T) {
	m := map[string]interface{}{
		"foo": "bar",
	}
	v, err := Get(m, "foo")
	assert.NoError(t, err)
	assert.Equal(t, "bar", v)
}

func TestNestedMap(t *testing.T) {
	m := map[string]interface{}{
		"nest": map[string]interface{}{
			"level2": "hello",
		},
	}
	v, err := Get(m, "nest.level2")
	assert.NoError(t, err)
	assert.Equal(t, "hello", v)
}

func TestTopLevelSliceIndex(t *testing.T) {
	a := []interface{}{"a", "b", 3}
	v, err := Get(a, "[0]")
	assert.NoError(t, err)
	assert.Equal(t, "a", v)
}

func TestNestedSliceIndex(t *testing.T) {
	m := map[string]interface{}{
		"foo": []interface{}{"a", "b", 3},
	}
	v, err := Get(m, "foo[0]")
	assert.NoError(t, err)
	assert.Equal(t, "a", v)
}

func TestNestedCrazy(t *testing.T) {
	m := map[string]interface{}{
		"foo": "bar",
		"baz": map[string]interface{}{
			"qux": map[string]interface{}{
				"wowe": []interface{}{
					map[string]interface{}{
						"hello": "dog",
					},
				},
			},
		},
	}

	v, err := Get(m, "foo")
	assert.NoError(t, err)
	assert.Equal(t, "bar", v)

	v, err = Get(m, "baz.qux.wowe[0].hello")
	assert.NoError(t, err)
	assert.Equal(t, "dog", v)
}

func TestArrPart(t *testing.T) {
	i, s := arrPart("[0]")
	assert.Equal(t, 0, i)
	assert.Empty(t, s)

	i, _ = arrPart("[[2]]")
	assert.Equal(t, -1, i)

	i, s = arrPart("[10].foo")
	assert.Equal(t, 10, i)
	assert.Equal(t, "foo", s)
}

func TestMapPart(t *testing.T) {
	v, q := mapPart("foo.bar")
	assert.Equal(t, "foo", v)
	assert.Equal(t, "bar", q)

	v, q = mapPart("foo[0]")
	assert.Equal(t, "foo", v)
	assert.Equal(t, "[0]", q)
}

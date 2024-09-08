package dig

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringKey(t *testing.T) {
	var jsonBlob = []byte(`{"foo": {"bar": {"baz": 1}}}`)
	var v interface{}
	if err := json.Unmarshal(jsonBlob, &v); err != nil {
		t.Fatal(err)
	}

	// nested successful lookup
	value, err := Dig(v, "foo", "bar", "baz")
	assert.Equal(t, float64(1), value, "foo.bar.baz should be 1")
	assert.Nil(t, err)

	// semi-nested successful lookup
	value, err = Dig(v, "foo", "bar")
	assert.Equal(t, map[string]interface{}{"baz": float64(1)}, value)
	assert.Nil(t, err)

	// lookup without keys fails
	value, err = Dig(v)
	assert.Nil(t, value)
	assert.Equal(t, "no key given", err.Error())

	// nested failed lookup, qux does not exist
	value, err = Dig(v, "foo", "qux", "quux")
	assert.Nil(t, value)
	assert.Equal(t, "key qux not found in map[bar:map[baz:1]]", err.Error())

	// nested failed lookup, unsupported format
	value, err = Dig(v, "foo", []int{1})
	assert.Nil(t, value)
	assert.Equal(t, "unsupported key type: [1]", err.Error())
}

func TestIntKey(t *testing.T) {
	var jsonBlob = []byte(`{"foo": [10, 11, 12]}`)
	var v interface{}
	if err := json.Unmarshal(jsonBlob, &v); err != nil {
		t.Fatal(err)
	}

	success, err := Dig(v, "foo", 1)
	assert.Equal(t, float64(11), success, "foo.bar.baz should be 1")
	assert.Nil(t, err)

	failure, err := Dig(v, "foo", 1, 0)
	assert.Nil(t, failure)
	assert.NotNil(t, err)

	failure2, err := Dig(v, "foo", "bar")
	assert.Nil(t, failure2)
	assert.NotNil(t, err)

	failure3, err := Dig(v, "foo", -1)
	assert.Nil(t, failure3)
	assert.NotNil(t, err)

	failure4, err := Dig(v, "foo", 3)
	assert.Nil(t, failure4)
	assert.NotNil(t, err)
}

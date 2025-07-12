package time_based_key_value_store_981

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKV(t *testing.T) {
	kv := Constructor()

	// original flow
	kv.Set("foo", "bar", 1)
	kv.Set("foo", "baz", 4)

	assert.Equal(t, "", kv.Get("foo", 0))
	assert.Equal(t, "bar", kv.Get("foo", 1))
	assert.Equal(t, "bar", kv.Get("foo", 2))
	assert.Equal(t, "bar", kv.Get("foo", 3))
	assert.Equal(t, "baz", kv.Get("foo", 4))
	assert.Equal(t, "baz", kv.Get("foo", 5))
	assert.Equal(t, "baz", kv.Get("foo", 6))

	// overwrite same timestamp
	kv.Set("foo", "qux", 4)
	assert.Equal(t, "qux", kv.Get("foo", 4))
	assert.Equal(t, "qux", kv.Get("foo", 999))

	// another key
	kv.Set("other", "val", 10)
	assert.Equal(t, "", kv.Get("other", 9))
	assert.Equal(t, "val", kv.Get("other", 10))
	assert.Equal(t, "val", kv.Get("other", 11))
	assert.Equal(t, "", kv.Get("missing", 100))

	// empty store
	empty := Constructor()
	assert.Equal(t, "", empty.Get("anything", 1))
}

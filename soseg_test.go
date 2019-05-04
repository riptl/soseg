package soseg

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestTree(t *testing.T) {
	var tree Tree
	tree.Put(0, 1)
	tree.Put(1, 3)
	tree.Put(2, 4)
	tree.Put(3, 1)
	tree.Put(4, 2)

	assert.Equal(t, tree.Total(), 11, "Wrong total amount")
	assert.Equal(t, tree.Size(), 5, "Wrong number of nodes")

	{
		val, offset, ok := tree.Get(2)
		assert.Equal(t, ok, true, "Not found but inserted")
		assert.Equal(t, val, 4, "Got wrong value")
		assert.Equal(t, offset, 4, "Got wrong offset")
	}

	{
		key, ok := tree.Find(6)
		assert.Equal(t, ok, true, "Not found but inserted")
		assert.Equal(t, key, 2, "Found wrong key")
	}

	{
		ok := tree.Remove(2)
		assert.Equal(t, ok, true, "Could not remove but was inserted")
	}

	{
		_, _, ok := tree.Get(2)
		assert.Equal(t, ok, false, "Found but was removed")
	}

	assert.Equal(t, tree.Total(), 7, "Wrong total amount")

	tree.Remove(0)
	tree.Remove(1)
	tree.Remove(3)

	assert.Equal(t, tree.Total(), 2, "Wrong total amount")

	tree.Remove(4)

	assert.Equal(t, tree.Total(), 0, "Tree isn't empty")
	assert.Equal(t, tree.Size(), 0, "Tree isn't empty")
}

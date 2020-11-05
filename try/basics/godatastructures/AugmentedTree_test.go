package godatastructures_test

import (
	"github.com/Workiva/go-datastructures/augmentedtree"
	"github.com/Workiva/go-datastructures/bitarray"
	"github.com/Workiva/go-datastructures/btree/immutable"
	"github.com/Workiva/go-datastructures/list"
	"gotest.tools/assert"
	"testing"
)

func TestAugmentedTree(t *testing.T) {
	data := augmentedtree.New(1000)
	t.Log(data)
	assert.Equal(t, 1 > 2, false)
	return
}

func TestBitarray(t *testing.T) {
	data := bitarray.NewBitArray(256, true)
	t.Log(data.ToNums())
	t.Log(data)
	assert.Equal(t, 1 > 2, false)
	return
}

func TestBtree(t *testing.T) {
	data := btree.New(btree.Config{
		NodeWidth:  1,
		Persister:  nil,
		Comparator: nil,
	})
	t.Log(data)
	t.Log(data.ID())
	t.Log(data.Len())
	a := btree.Tr{
		UUID:      nil,
		Count:     1,
		Root:      nil,
		NodeWidth: 1,
	}
	t.Log(a.Root)
	assert.Equal(t, 1 > 2, false)
	return
}

func TestList(t *testing.T) {
	data := list.Empty
	t.Log(data.IsEmpty())
	data = data.Add(1)
	data = data.Add(12)
	data = data.Add(13)
	t.Log(data.Length())
	t.Log("=========")
	for i := uint(0); i < data.Length(); i++ {
		t.Log(data.Get(i))
	}
	data, err := data.Insert("你好啊", data.Length())
	if err != nil {
		t.Log(err)
	}
	t.Log("=========")
	for i := uint(0); i < data.Length(); i++ {
		t.Log(data.Get(i))
	}
}

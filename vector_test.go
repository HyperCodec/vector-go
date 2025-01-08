package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	v := EmptyVector[int](5)

	assert.True(t, v.Push(1))
	assert.False(t, v.Push(2))
	assert.False(t, v.Push(3))
	assert.False(t, v.PushBack(4))

	assert.Equal(t, []int{4, 1, 2, 3, 0}, v.data)
}

func TestIndexFuncs(t *testing.T) {
	v := VectorFromSlice([]int{1, 2, 3, 4, 5}, 5)

	// set
	assert.Nil(t, v.Set(0, 3))
	assert.Nil(t, v.Set(2, 4))

	assert.EqualError(t, v.Set(-1, 5), "index out of bounds")
	assert.EqualError(t, v.Set(5, 10), "index out of bounds")

	assert.Equal(t, []int{3, 2, 4, 4, 5}, v.data)

	// get
	val, err := v.Get(0)

	assert.Nil(t, err)
	assert.Equal(t, 3, *val)

	val, err = v.Get(2)

	assert.Nil(t, err)
	assert.Equal(t, 4, *val)

	_, err = v.Get(-1)
	assert.EqualError(t, err, "index out of bounds")
}

func TestInsert(t *testing.T) {
	v := VectorFromSlice([]int{1, 2, 3}, 5)

	alloc, err := v.Insert(1, 4)

	assert.Nil(t, err)
	assert.True(t, alloc)

	alloc, err = v.Insert(2, 5)

	assert.Nil(t, err)
	assert.False(t, alloc)

	_, err = v.Insert(-1, 1)

	assert.EqualError(t, err, "index out of bounds")
}

func TestMutateData(t *testing.T) {
	v := VectorFromSlice([]int{1, 2, 3}, 5)

	slice := v.Data()

	slice[0] = 2

	assert.Equal(t, []int{2, 2, 3}, v.data)
}
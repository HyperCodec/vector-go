package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushRemove(t *testing.T) {
	v, _ := EmptyVector[int](5)

	assert.True(t, v.PushBack(1))
	assert.False(t, v.PushBack(2))
	assert.False(t, v.PushBack(3))
	assert.False(t, v.PushFront(4))

	val, err := v.Remove(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, *val)

	assert.Equal(t, []int{4, 2, 3, 0}, v.data)
}

func TestIndexFuncs(t *testing.T) {
	v, _ := VectorFromSlice([]int{1, 2, 3, 4, 5}, 5)

	// set
	assert.Nil(t, v.Set(0, 3))
	assert.Nil(t, v.Set(2, 4))

	assert.EqualError(t, v.Set(-1, 5), OutOfBounds)
	assert.EqualError(t, v.Set(5, 10), OutOfBounds)

	assert.Equal(t, []int{3, 2, 4, 4, 5}, v.data)

	// get
	val, err := v.Get(0)

	assert.Nil(t, err)
	assert.Equal(t, 3, *val)

	val, err = v.Get(2)

	assert.Nil(t, err)
	assert.Equal(t, 4, *val)

	_, err = v.Get(-1)
	assert.EqualError(t, err, OutOfBounds)
}

func TestInsert(t *testing.T) {
	v, _ := VectorFromSlice([]int{1, 2, 3}, 5)

	alloc, err := v.Insert(1, 4)

	assert.Nil(t, err)
	assert.True(t, alloc)

	alloc, err = v.Insert(2, 5)

	assert.Nil(t, err)
	assert.False(t, alloc)

	_, err = v.Insert(-1, 1)

	assert.EqualError(t, err, OutOfBounds)
}

func TestExtraneousMutate(t *testing.T) {
	v, _ := VectorFromSlice([]int{1, 2, 3}, 5)

	slice := v.Data()

	slice[0] = 2

	assert.Equal(t, []int{2, 2, 3}, v.data)

	// TODO mutate pointers from Get.
}
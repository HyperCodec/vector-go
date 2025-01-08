package vector

import (
	"errors"
)

type Vector[T any] struct {
	data        []T
	len         int
	capacity    int
	AllocAmount int
}

func VectorFromSlice[T any](slice []T, allocAmount int) *Vector[T] {
	size := len(slice)
	return &Vector[T]{data: slice, len: size, capacity: size, AllocAmount: allocAmount}
}

func EmptyVector[T any](allocAmount int) *Vector[T] {
	return &Vector[T]{data: []T{}, len: 0, capacity: 0, AllocAmount: allocAmount}
}

func EmptyVectorWithCapacity[T any](capacity, allocAmount int) *Vector[T] {
	return &Vector[T]{data: make([]T, capacity), len: 0, capacity: capacity, AllocAmount: allocAmount}
}

func (v *Vector[T]) Len() int {
	return v.len
}

func (v *Vector[T]) Capacity() int {
	return v.capacity
}

func (v *Vector[T]) AddCapacity(amount int) {
	v.capacity += amount

	newSlice := make([]T, v.capacity)
	copy(newSlice, v.data)

	v.data = newSlice
}

func (v *Vector[T]) Push(val T) bool {
	allocated := v.len == v.capacity

	if allocated {
		v.AddCapacity(v.AllocAmount)
	}

	v.data[v.len] = val
	v.len++

	return allocated
}

func (v *Vector[T]) PushBack(val T) bool {
	allocated, err := v.Insert(0, val)
	if err != nil {
		panic("unreachable")
	}

	return allocated
}

func (v *Vector[T]) Insert(index int, val T) (bool, error) {
	if index < 0 || index > v.len {
		return false, errors.New("index out of bounds")
	}

	allocated := v.len == v.capacity

	if allocated {
		v.AddCapacity(v.AllocAmount)
	}

	v.len++
	newData := make([]T, v.capacity)

	i := 0
	for j := 0; j < v.len; j++ {
		if j == index {
			newData[j] = val
			continue
		}
		newData[j] = v.data[i]
		i++
	}

	v.data = newData

	return allocated, nil
}

func (v *Vector[T]) Get(index int) (*T, error) {
	if index < 0 || index >= v.len {
		return nil, errors.New("index out of bounds")
	}

	return &v.data[index], nil
}

func (v *Vector[T]) GetUnchecked(index int) *T {
	return &v.data[index]
}

func (v *Vector[T]) Set(index int, val T) error {
	if index < 0 || index >= v.len {
		return errors.New("index out of bounds")
	}

	v.data[index] = val
	return nil
}

func (v *Vector[T]) SetUnchecked(index int, val T) {
	v.data[index] = val
}

func (v *Vector[T]) Copy(dst []T) int {
	return copy(dst, v.data[:v.len])
}

func (v *Vector[T]) Data() []T {
	return v.data[:v.len]
}
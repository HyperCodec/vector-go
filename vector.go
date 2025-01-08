package vector

import (
	"errors"
	"slices"
)

type Vector[T any] struct {
	data     []T
	size     int
	capacity int
	AllocAmount int
}

func VectorFromSlice[T any](slice []T, allocAmount int) *Vector[T] {
	size := len(slice)
	return &Vector[T]{data: slice, size: size, capacity: size, AllocAmount: allocAmount}
}

func EmptyVector[T any](allocAmount int) *Vector[T] {
	return &Vector[T]{data: []T{}, size: 0, capacity: 0, AllocAmount: allocAmount}
}

func EmptyVectorWithCapacity[T any](capacity, allocAmount int) *Vector[T] {
	return &Vector[T]{data: make([]T, capacity), size: 0, capacity: capacity, AllocAmount: allocAmount}
}

func (v *Vector[T]) Size() int {
	return v.size
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
	allocated := v.size == v.capacity
	
	if allocated {
		v.AddCapacity(v.AllocAmount)
	}

	v.data[v.size] = val
	v.size++

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
	if index < 0 || index > v.size {
		return false, errors.New("index out of bounds")
	}
	
	allocated := v.size == v.capacity

	if allocated {
		v.AddCapacity(v.AllocAmount)
	}

	v.data = slices.Insert(v.data, index, val)
	v.size++

	return allocated, nil
}

func (v *Vector[T]) Get(index int) (*T, error) {
	if index < 0 || index >= v.size {
		return nil, errors.New("index out of bounds")
	}

	return &v.data[index], nil
}

func (v *Vector[T]) GetUnchecked(index int) *T {
	return &v.data[index]
}

func (v *Vector[T]) Set(index int, val T) error {
	if index < 0 || index >= v.size {
		return errors.New("index out of bounds")
	}

	v.data[index] = val
	return nil
}

func (v *Vector[T]) SetUnchecked(index int, val T) {
	v.data[index] = val
}

func (v *Vector[T]) Copy(dst []T) int {
	return copy(dst, v.data[:v.size])
}
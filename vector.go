/*
A package that adds a variable-length Vector datatype, similar to the Rust implementation.
*/
package vector

import (
	"errors"
	"slices"
)

const (
	OutOfBounds        = "index out of bounds"
	InvalidAllocAmount = "invalid `allocAmount`"
	CannotAddAmount    = "cannot add this amount"
)

/*
A variable-length collection datatype that allows for simple/efficient pushing and insertion.
*/
type Vector[T any] struct {
	data        []T
	len         int
	capacity    int
	allocAmount int
}

/*
Create a Vector from a slice with capacity len(slice).

Returns an error if allocAmount <= 0.
*/
func FromSlice[T any](slice []T, allocAmount int) *Vector[T]{
	if allocAmount <= 0 {
		panic(InvalidAllocAmount)
	}

	size := len(slice)
	return &Vector[T]{data: slice, len: size, capacity: size, allocAmount: allocAmount}
}

/*
Create an empty Vector with a capacity of 0.

Returns an error if allocAmount <= 0.
*/
func Empty[T any](allocAmount int) *Vector[T] {
	if allocAmount <= 0 {
		panic(InvalidAllocAmount)
	}

	return &Vector[T]{data: []T{}, len: 0, capacity: 0, allocAmount: allocAmount}
}

/*
Create an empty Vector with a specified capacity.

Returns an error if allocAmount <= 0.
*/
func EmptyWithCapacity[T any](capacity, allocAmount int) *Vector[T] {
	if allocAmount <= 0 {
		panic(InvalidAllocAmount)
	}

	return &Vector[T]{data: make([]T, capacity), len: 0, capacity: capacity, allocAmount: allocAmount}
}

/*
Get the length of the vector. Runs in O(1) time.
*/
func (v *Vector[T]) Len() int {
	return v.len
}

/*
Get the current capacity of the vector.
*/
func (v *Vector[T]) Capacity() int {
	return v.capacity
}

/*
Get the current allocation amount.
*/
func (v *Vector[T]) AllocAmount() int {
	return v.allocAmount
}

/*
Set the allocation amount to newVal.

Returns an error if newVal <= 0.
*/
func (v *Vector[T]) SetAllocAmount(newVal int) error {
	if newVal <= 0 {
		return errors.New(InvalidAllocAmount)
	}

	v.allocAmount = newVal

	return nil
}

/*
Add new capacity to the vector. Takes O(newCapacity) time to copy the vector's elements to a larger allocation.

Returns an error if amount <= 0.
*/
func (v *Vector[T]) AddCapacity(amount int) error {
	if amount <= 0 {
		return errors.New(CannotAddAmount)
	}

	v.capacity += amount

	newSlice := make([]T, v.capacity)
	copy(newSlice, v.data)

	v.data = newSlice

	return nil
}

/*
Appends an item to the end of the Vector. Runs in O(1) time if there is no allocation. Otherwise takes O(newCapacity) time
to copy values to a bigger allocation.

Returns whether an allocation has occurred.
*/
func (v *Vector[T]) PushBack(val T) bool {
	allocated := v.len == v.capacity

	if allocated {
		_ = v.AddCapacity(v.allocAmount)
	}

	v.data[v.len] = val
	v.len++

	return allocated
}

/*
Inserts an element at index 0. Takes the same amount of time as insertion.

Returns whether an allocation has occurred.
*/
func (v *Vector[T]) PushFront(val T) bool {
	allocated, err := v.Insert(0, val)
	if err != nil {
		panic("unreachable")
	}

	return allocated
}

/*
Inserts an element at the index. Takes O(capacity) time without an allocation, or O(newCapacity) with an allocation.

Returns whether an allocation has occurred. Otherwise it returns an error if the index is out of bounds.
*/
func (v *Vector[T]) Insert(index int, val T) (bool, error) {
	if err := v.boundsCheck(index); err != nil {
		return false, err
	}

	allocated := v.len == v.capacity

	if allocated {
		_ = v.AddCapacity(v.allocAmount)
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

/*
Gets a pointer to the value at a specified index.

Returns the pointer to the value. Otherwise it returns an error if the index is out of bounds.
*/
func (v *Vector[T]) Get(index int) (*T, error) {
	if err := v.boundsCheck(index); err != nil {
		return nil, err
	}

	return &v.data[index], nil
}

/*
Gets a pointer to the value at the specified index without checking that the index is in bounds (panics if out of bounds).
*/
func (v *Vector[T]) GetUnchecked(index int) *T {
	return &v.data[index]
}

/*
Sets the value at the specified index.

Returns an error if the index is out of bounds.
*/
func (v *Vector[T]) Set(index int, val T) error {
	if err := v.boundsCheck(index); err != nil {
		return err
	}

	v.data[index] = val
	return nil
}

/*
Sets the value at the specified index without checking whether the index is in bounds.
*/
func (v *Vector[T]) SetUnchecked(index int, val T) {
	v.data[index] = val
}

/*
Copies the data of the Vector to another slice.

Returns the amount of elements written.
*/
func (v *Vector[T]) Copy(dst []T) int {
	return copy(dst, v.data[:v.len])
}

/*
Get a slice of the data that is within the correct range. Mutating this value will mutate the original Vector.

Warning: Do not use this value after modifying the vector elsewhere (especially if the capacity changes) as it likely will not point to the correct data anymore.
*/
func (v *Vector[T]) Data() []T {
	return v.data[:v.len]
}

/*
Removes the value at the specified index.

Returns the removed value. Returns an error if the index is out of bounds.
*/
func (v *Vector[T]) Remove(index int) (*T, error) {
	if err := v.boundsCheck(index); err != nil {
		return nil, err
	}

	val := v.data[index]

	v.data = slices.Delete(v.data, index, index+1)
	v.capacity--
	v.len--

	return &val, nil
}

/*
Removes the value at the specified index and returns it without checking that the index is in bounds (panics if out of bounds).
*/
func (v *Vector[T]) RemoveUnchecked(index int) *T {
	val := v.data[index]

	v.data = slices.Delete(v.data, index, index+1)
	v.capacity--
	v.len--

	return &val
}

func (v *Vector[T]) boundsCheck(index int) error {
	if !v.IsInBounds(index) {
		return errors.New("index out of bounds")
	}

	return nil
}

/*
Returns true if the index is valid. Returns false if using it would return an index out of bounds error.
*/
func (v *Vector[T]) IsInBounds(index int) bool {
	return index >= 0 && index < v.len
}

/*
Returns the index of the first instance of val in v. If that value does not exist, it returns -1. Requires the type in the vector to be comparable.
*/
func Find[T comparable](v *Vector[T], val T) int {
	for i := range v.len {
		if v.data[i] == val {
			return i
		}
	}

	return -1
}

/*
Returns whether or not the value exists in the vector. Requires the type in the vector to be comparable.
*/
func Contains[T comparable](v *Vector[T], val T) bool {
	for i := range v.len {
		if v.data[i] == val {
			return true
		}
	}

	return false
}

// TODO maybe RemoveMultiple.

/*
A package that adds a variable-length Vector datatype, similar to the Rust implementation.

TODO make sure go actually parses Markdown in docs. Also run doctests or something to make sure this code actually works.
### Example
```go
import (
	"fmt"

	"github.com/HyperCodec/vector-go"
)

func main() {
	// create a vector with a capacity of 3 and an allocation amount of 5.
	v := vector.EmptyVectorWithCapacity(3, 5)
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	
	fmt.Println(v.Data())
}
```
*/
package vector

import (
	"errors"
)

/*
A variable-length collection datatype that allows for simple/efficient pushing and insertion.

`AllocAmount`: The amount of entries to allocate when the vector runs out of room.
*/
type Vector[T any] struct {
	data        []T
	len         int
	capacity    int
	AllocAmount int
}

/*
Create a `Vector` from a slice with capacity `len(slice)`.

Returns an error if `allocAmount <= 0`.
*/
func VectorFromSlice[T any](slice []T, allocAmount int) (*Vector[T], error) {
	if allocAmount <= 0 {
		return nil, errors.New("invalid `allocAmount`")
	}

	size := len(slice)
	return &Vector[T]{data: slice, len: size, capacity: size, AllocAmount: allocAmount}, nil
}

/*
Create an empty `Vector` with a `capacity` of 0.

Returns an error if `allocAmount <= 0`.
*/
func EmptyVector[T any](allocAmount int) (*Vector[T], error) {
	if allocAmount <= 0 {
		return nil, errors.New("invalid `allocAmount`")
	}

	return &Vector[T]{data: []T{}, len: 0, capacity: 0, AllocAmount: allocAmount}, nil
}

/*
Create an empty `Vector` with a specified `capacity`.

Returns an error if `allocAmount <= 0`.
*/
func EmptyVectorWithCapacity[T any](capacity, allocAmount int) (*Vector[T], error) {
	if allocAmount <= 0 {
		return nil, errors.New("invalid `allocAmount`")
	}

	return &Vector[T]{data: make([]T, capacity), len: 0, capacity: capacity, AllocAmount: allocAmount}, nil
}

/*
Get the length of the vector. Runs in `O(1)` time.
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
Add new capacity to the vector. Takes `O(newCapacity)` time to copy the vector's elements to a larger allocation.

Returns an error if `amount <= 0`.
*/
func (v *Vector[T]) AddCapacity(amount int) error {
	if amount <= 0 {
		return errors.New("cannot add this amount")
	}
	
	v.capacity += amount

	newSlice := make([]T, v.capacity)
	copy(newSlice, v.data)

	v.data = newSlice

	return nil
}

/*
Appends an item to the end of the `Vector`. Runs in `O(1)` time if there is no allocation. Otherwise takes `O(newCapacity)` time 
to copy values to a bigger allocation.

Returns whether an allocation has occurred.
*/
func (v *Vector[T]) PushBack(val T) bool {
	allocated := v.len == v.capacity

	if allocated {
		v.AddCapacity(v.AllocAmount)
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
Inserts an element at the index. Takes `O(capacity)` time without an allocation, or `O(newCapacity)` with an allocation.

Returns whether an allocation has occurred. Otherwise it returns an error if the index is out of bounds.
*/
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

/*
Gets a pointer to the value at a specified index.

Returns the pointer to the value. Otherwise it returns an error if the index is out of bounds.
*/
func (v *Vector[T]) Get(index int) (*T, error) {
	if index < 0 || index >= v.len {
		return nil, errors.New("index out of bounds")
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
	if index < 0 || index >= v.len {
		return errors.New("index out of bounds")
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
Copies the data of the `Vector` to another slice.

Returns the amount of elements written.
*/
func (v *Vector[T]) Copy(dst []T) int {
	return copy(dst, v.data[:v.len])
}

/*
Get a slice of the data that is within the correct range. Mutating this value will mutate the original `Vector`.

Warning: Do not use this value after modifying the vector elsewhere (especially if an allocation occurs) as it likely will not point to the correct data anymore.
*/
func (v *Vector[T]) Data() []T {
	return v.data[:v.len]
}

// TODO `Remove` method.
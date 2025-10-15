package arrays

import (
	"errors"
	"fmt"
)

// When and where is a static used
// 1) Storing and accessing sequentila data
// 2) Temporarily storing objects
// 3) Used by IO routines as buffers
// 4) Lookup tables and inverse lookup tables
// 5) Can be used to return multiple values from a function
// 6) Used in dynamic programing to cache answers to subproblems

// Static array
type Array struct {
	items [10]int
}

// Dynamic array
type DynamicArray struct {
	data     []int
	length   int
	capacity int
}

func (d *DynamicArray) NewDynamicArray(initialCapacicty int) *DynamicArray {
	return &DynamicArray{
		data:     make([]int, initialCapacicty),
		length:   0,
		capacity: initialCapacicty,
	}
}

func (d *DynamicArray) Append(value int) {
	if d.length == d.capacity {
		d.resize()
	}
	d.data[d.length] = value
	d.length++
}

func (d *DynamicArray) resize() {
	newCapacity := d.capacity
	if newCapacity == 0 {
		newCapacity = 1
	} else {
		newCapacity *= 2
	}

	newData := make([]int, newCapacity)

	for i := 0; i < d.length; i++ {
		newData[i] = d.data[i]
	}
	d.data = newData
	d.capacity = newCapacity
}

func (d *DynamicArray) Get(index int) (int, error) {
	if index < 0 || index >= d.length {
		return 0, errors.New("index out of bounds")
	}
	return d.data[index], nil
}

func (d *DynamicArray) Cap() int {
	return d.capacity
}
func (d *DynamicArray) Len() int {
	return d.length
}

func (d *DynamicArray) Print() {
	fmt.Print("[")
	for i := 0; i < d.length; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(d.data[i])
	}
	fmt.Println("]")
}

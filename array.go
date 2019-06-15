package __

import (
	"reflect"
)

type Array []interface{}

func NewArray(i interface{}) Array {
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		vSlice := reflect.ValueOf(i)
		var slice []interface{}

		for j := 0; j < vSlice.Len(); j++ {
			slice = append(slice, vSlice.Index(j).Interface())
		}
		return slice
	default:
		panic("type must be slice or array")
	}
}

func NewEmptyArray() Array {
	return NewArray([]interface{}{})
}

func ToInterfaceArray(array Array) []interface{} {
	return []interface{}(array)
}

func Len(array Array) int {
	return len(array)
}

// This func mutates array
func Push(array *Array, element interface{}) {
	*array = append(*array, element)
}

// This func mutates array
func Pop(array *Array) interface{} {
	interfaceArray := array.InterfaceArray()
	poppedItem := interfaceArray[len(interfaceArray)-1]
	*array = interfaceArray[0 : len(interfaceArray)-1]
	return poppedItem
}

// This func mutates array
func Enqueue(array *Array, element interface{}) {
	*array = append(*array, element)
}

// This func mutates array
func Dequeue(array *Array) interface{} {
	interfaceArray := array.InterfaceArray()
	dequeuedItem := interfaceArray[0]
	*array = interfaceArray[1:]
	return dequeuedItem
}

func Nth(array *Array, index int) interface{} {
	return array.InterfaceArray()[index]
}

func Map(array *Array, mapFunc func(item interface{}, index int) interface{}) Array {
	resultArr := NewEmptyArray()

	for i, value := range array.Iterable() {
		resultArr.Push(mapFunc(value, i))
	}
	return resultArr
}

func Filter(array *Array, filterFunc func(item interface{}, index int) bool) Array {
	resultArr := NewEmptyArray()

	for i, value := range array.Iterable() {
		if filterFunc(value, i) {
			resultArr.Push(value)
		}
	}
	return resultArr
}

func Find(array *Array, findFunc func(item interface{}, index int) bool) interface{} {
	for i, value := range array.Iterable() {
		if findFunc(value, i) {
			return value
		}
	}
	return nil
}

func FindIndex(array *Array, findFunc func(item interface{}, index int) bool) int {
	for i, value := range array.Iterable() {
		if findFunc(value, i) {
			return i
		}
	}
	return -1
}

func Includes(array *Array, targetValue interface{}) bool {
	for _, value := range array.Iterable() {
		if value == targetValue {
			return true
		}
	}
	return false
}

func Chunk(array *Array, chunkSize int) Array {
	if chunkSize == 0 {
		panic("chunk size must be greater than 0")
	}

	resultArr := NewEmptyArray()
	var chunkArr Array
	for i, value := range array.Iterable() {
		mod := i % chunkSize

		if mod == 0 {
			chunkArr = NewEmptyArray()
		}

		chunkArr.Push(value)

		if mod == chunkSize - 1 || i == array.Len() - 1 {
			resultArr.Push(chunkArr)
		}
	}
	return resultArr
}

// This func mutates array
func Concat(array *Array, elements ...interface{}) {
	for _, element := range elements {
		Push(array, element)
	}
}

// This func mutates array
func Fill(array *Array, item interface{}, startIndex, endIndex int) {
	if startIndex > endIndex ||
		startIndex < 0 ||
		endIndex < 0 ||
		startIndex >= len(*array) ||
		endIndex >= len(*array) {
		panic("invalid start/end index")
	}

	resultArr := Map(array, func(iterItem interface{}, index int) interface{} {
		if startIndex <= index && index <= endIndex {
			return item
		}

		return iterItem
	})

	*array = resultArr
}

func HasIndex(array *Array, index int) bool {
	if index < len(*array) {
		return true
	}
	return false
}

func GetAt(array *Array, index int) interface{} {
	if HasIndex(array, index) {
		arr := array.InterfaceArray()
		return arr[index]
	}

	panic("index out of range")
}

func Flatten(array *Array) Array {
	resultArr := NewEmptyArray()
	for _, value := range array.Iterable() {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Array, reflect.Slice:
			vSlice := reflect.ValueOf(value)
			for j := 0; j < vSlice.Len(); j++ {
				resultArr.Push(vSlice.Index(j).Interface())
			}
		default:
			resultArr.Push(value)
		}
	}
	return resultArr
}

func Union(arrays ...*Array) Array {
	if len(arrays) < 2 {
		panic("at least 2 arrays required")
	}

	resultArr := NewEmptyArray()

	for _, array := range arrays {
		for _, value := range array.Iterable() {
			if Includes(array, value) {
				resultArr.Push(value)
			}
		}
	}

	return resultArr
}

func Intersection(arrays ...*Array) Array {
	if len(arrays) < 2 {
		panic("at least 2 arrays required")
	}

	resultArr := NewEmptyArray()
	intersectionCount := map[interface{}]int{}

	for _, array := range arrays {
		for _, value := range array.Iterable() {
			_, ok := intersectionCount[value]
			if !ok {
				intersectionCount[value] = 0
			}

			intersectionCount[value]++

			if intersectionCount[value] == 
		}
	}

	return resultArr
}
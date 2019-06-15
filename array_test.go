package __

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLen(t *testing.T) {
	a := []int{1,2,3}
	arrA := NewArray(a)

	assert.Equal(t, 3, Len(arrA))
}

func TestPush(t *testing.T) {
	a := []int{1,2,3}
	arrA := NewArray(a)
	Push(&arrA, 4)
	Push(&arrA, "I'm not a string")
	Push(&arrA, true)

	assert.Len(t, arrA, 6)
}

func TestPop(t *testing.T) {
	a := []int{1,2,3}
	arrA := NewArray(a)
	item := Pop(&arrA)

	assert.Len(t, arrA, 2)
	assert.Equal(t, 3, item.(int))
}

func TestEnqueue(t *testing.T) {
	a := []int{1,2,3}
	arrA := NewArray(a)
	Enqueue(&arrA, 4)
	Enqueue(&arrA, "I'm not a string")
	Enqueue(&arrA, true)

	assert.Len(t, arrA, 6)
}

func TestDequeue(t *testing.T) {
	a := []int{1,2,3}
	arrA := NewArray(a)
	item := Dequeue(&arrA)

	assert.Len(t, arrA, 2)
	assert.Equal(t, 1, item.(int))
}

func TestNth(t *testing.T) {
	a := []int{1,2,3}
	arrA := NewArray(a)
	item := Nth(&arrA, 1)

	assert.Equal(t, 2, item)
}

func TestMap(t *testing.T) {
	a := []int{1,2,3}
	arrA := NewArray(a)
	arrA = Map(&arrA, func(item interface{}, index int) interface{} {
		return item.(int) * 2
	})

	assert.Len(t, arrA, 3)
	assert.Equal(t, 2, arrA.InterfaceArray()[0].(int))
	assert.Equal(t, 4, arrA.InterfaceArray()[1].(int))
	assert.Equal(t, 6, arrA.InterfaceArray()[2].(int))
}

func TestFilter(t *testing.T) {
	a := []int{1,2,3,4,5,6,7,8}
	arrA := NewArray(a)
	arrA = Filter(&arrA, func(item interface{}, index int) bool {
		return item.(int) % 2 == 0
	})

	assert.Len(t, arrA, 4)
	assert.Equal(t, 2, arrA.InterfaceArray()[0].(int))
	assert.Equal(t, 4, arrA.InterfaceArray()[1].(int))
	assert.Equal(t, 6, arrA.InterfaceArray()[2].(int))
	assert.Equal(t, 8, arrA.InterfaceArray()[3].(int))
}

func TestFind(t *testing.T) {
	a := []int{1,2,3,4,5,6,7,8}
	arrA := NewArray(a)
	item := Find(&arrA, func(item interface{}, index int) bool {
		return item.(int) == 5
	})

	assert.Equal(t, 5, item)
}

func TestChunk(t *testing.T) {
	a := []int{1,2,3,4,5,6,7,8}
	arrA := NewArray(a)
	arrA = Chunk(arrA, 3)
	chunk1 := arrA.InterfaceArray()[0].(Array)
	chunk2 := arrA.InterfaceArray()[1].(Array)
	chunk3 := arrA.InterfaceArray()[2].(Array)

	assert.Len(t, arrA, 3)
	assert.Len(t, chunk1, 3)
	assert.Len(t, chunk2, 3)
	assert.Len(t, chunk3, 2)
	assert.Equal(t, 1, chunk1.InterfaceArray()[0])
	assert.Equal(t, 2, chunk1.InterfaceArray()[1])
	assert.Equal(t, 3, chunk1.InterfaceArray()[2])
}

func TestConcat(t *testing.T) {
	a := []int{1,2,3}
	arrA := NewArray(a)
	Concat(&arrA, "x", 4, true, nil)

	assert.Len(t, arrA, 7)
	assert.Equal(t, "x", arrA.Nth(3))
	assert.Equal(t, 4, arrA.Nth(4))
	assert.Equal(t, true, arrA.Nth(5))
	assert.Equal(t, nil, arrA.Nth(6))
}

func TestFill(t *testing.T) {
	t.Run("Okay", func(t *testing.T) {
		a := []int{1,2,3,4,5,6,7,8}
		arrA := NewArray(a)
		Fill(&arrA, "Hello", 3,6)

		assert.Len(t, arrA, 8)
		assert.Equal(t, 3, arrA.Nth(2))
		assert.Equal(t, "Hello", arrA.Nth(3))
		assert.Equal(t, "Hello", arrA.Nth(4))
		assert.Equal(t, "Hello", arrA.Nth(5))
		assert.Equal(t, "Hello", arrA.Nth(6))
		assert.Equal(t, 8, arrA.Nth(7))
	})

	t.Run("Panic", func(t *testing.T) {
		a := []int{1,2,3,4,5,6,7,8}
		arrA := NewArray(a)

		assert.Panics(t, func() {
			Fill(&arrA, "Hello", 6,3)
		})
		assert.Panics(t, func() {
			Fill(&arrA, "Hello", 0,10)
		})
	})
}

func TestFlatten(t *testing.T) {
	arr := NewEmptyArray()
	arr.Push(1)
	arr.Push("A")
	arr.Push(NewArray([]string{"hello", "world"}))
	arr.Push(NewArray([]int{5, 6, 7}))

	arr = Flatten(&arr)
	assert.Len(t, arr, 7)
	assert.Equal(t, 1, arr.GetAt(0))
	assert.Equal(t, "A", arr.GetAt(1))
	assert.Equal(t, "hello", arr.GetAt(2))
	assert.Equal(t, "world", arr.GetAt(3))
	assert.Equal(t, 5, arr.GetAt(4))
	assert.Equal(t, 6, arr.GetAt(5))
	assert.Equal(t, 7, arr.GetAt(6))
}
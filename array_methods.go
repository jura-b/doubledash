package __

func (array *Array) Iterable() []interface{} {
	return ToInterfaceArray(*array)
}

func (array *Array) InterfaceArray() []interface{} {
	return ToInterfaceArray(*array)
}

func (array *Array) Len() int {
	return Len(*array)
}

func (array *Array) Push(element interface{}) {
	Push(array, element)
}

func (array *Array) Pop() interface{} {
	return Pop(array)
}

func (array *Array) Enqueue(element interface{}) {
	Enqueue(array, element)
}

func (array *Array) Dequeue() interface{} {
	return Dequeue(array)
}

func (array *Array) Nth(index int) interface{} {
	return Nth(array, index)
}

func (array *Array) Map(mapFunc func(item interface{}, index int) interface{}) Array {
	return Map(array, mapFunc)
}

func (array *Array) Filter(filterFunc func(item interface{}, index int) bool) Array {
	return Filter(array, filterFunc)
}

func (array *Array) Chunk(chunkSize int) Array {
	return Chunk(array, chunkSize)
}

func (array *Array) Concat(elements ...interface{}) {
	Concat(array, elements)
}

func (array *Array) Fill(item interface{}, startIndex, endIndex int) {
	Fill(array, item, startIndex, endIndex)
}

func (array *Array) HasIndex(index int) bool {
	return HasIndex(array, index)
}

func (array *Array) GetAt(index int) interface{} {
	return GetAt(array, index)
}

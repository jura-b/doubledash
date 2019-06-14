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



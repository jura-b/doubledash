package __

import (
	"strings"
)

// Indexing not support
func Get(object map[string]interface{}, path string) (interface{}, bool) {
	if path == "" {
		return nil, false
	}

	fieldStack := NewArray(strings.Split(path, "."))
	return get(object, &fieldStack)
}

func get(object map[string]interface{}, fieldStack *Array) (interface{}, bool) {
	if fieldStack.Len() == 0 {
		return object, true
	}
	fieldName := Dequeue(fieldStack).(string)

	if value, ok := object[fieldName]; ok {
		if fieldStack.Len() == 0 {
			return value, true
		}

		if castedValue, ok := value.(map[string]interface{}); ok {
			return get(castedValue, fieldStack)
		}
	}

	return nil, false
}

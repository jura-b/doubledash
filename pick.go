package __

func Pick(object map[string]interface{}, fields ...string) map[string]interface{} {
	resultObject := map[string]interface{}{}
	for _, field := range fields {
		resultObject[field] = object[field]
	}
	return resultObject
}
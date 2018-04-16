package accumulate

import (
	"reflect"
)

func Accumulate(collection interface{}, f interface{}) []interface{} {
	slice, _ := convertToSlice(collection)
	accumulator, _ := convertArg(f, reflect.Func)

	result := make([]interface{}, len(slice))
	for i, elem := range slice {
		result[i] = invokeFn(accumulator, elem)
	}

	return result
}

// convertToSlice creates a generic slice of elements from the given collection
func convertToSlice(c interface{}) (slice []interface{}, ok bool) {
	collection, ok := convertArg(c, reflect.Slice)

	if !ok {
		return
	}

	n := collection.Len()
	slice = make([]interface{}, n)

	for i := 0; i < n; i++ {
		slice[i] = collection.Index(i).Interface()
	}

	return slice, true
}

// invokeFn takes a func of a single argument and returns the result of applying that func on the given arg.
func invokeFn(fn reflect.Value, arg interface{}) interface{} {
	rarg := []reflect.Value{reflect.ValueOf(arg)}
	vals := fn.Call(rarg) // slice of one element
	return vals[0]
}

// convertArg ensures that the value of arg is of the given kind, to ensure the collection
// given to convertToSlice is a slice.
func convertArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)

	if val.Kind() == kind {
		ok = true
	}
	return
}

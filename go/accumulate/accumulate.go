package accumulate

import (
	"reflect"
)

func Accumulate(collection interface{}, f interface{}) []interface{} {
	slice := convertArg(collection, reflect.Slice)
	accumulator := convertArg(f, reflect.Func)
	n := slice.Len()
	result := make([]interface{}, slice.Len())
	for i := 0; i < n; i++ {
		result[i] = invoke(accumulator, slice.Index(i))
	}

	return result
}

// invoke takes a func of a single argument and returns the result of applying that func on the given arg.
func invoke(fn reflect.Value, arg reflect.Value) reflect.Value {
	rarg := []reflect.Value{arg}
	return fn.Call(rarg)[0] //slice of one element
}

// convertArg ensures that the value of arg is of the given kind
func convertArg(arg interface{}, kind reflect.Kind) (val reflect.Value) {
	val = reflect.ValueOf(arg)

	if val.Kind() != kind {
		panic("incorrect type reflected")
	}

	return
}

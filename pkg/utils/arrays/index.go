package arrays

import (
	"errors"
	"reflect"
)

var mySlice = []interface{}{}

func ArrayInsideArray(valores interface{}, array interface{}) (interface{}, error) {

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		// If it's a slice
		return printSliceArray(reflect.ValueOf(valores), reflect.ValueOf(array))

	case reflect.Map:
		// If it's a map
		return printMapVArray(reflect.ValueOf(valores), reflect.ValueOf(array))

	case reflect.Struct:
		// If it's a struct
		return printStructArray(reflect.ValueOf(valores), reflect.ValueOf(array))
	case reflect.Array:
		// If it's a array
		return printSliceArray(reflect.ValueOf(valores), reflect.ValueOf(array))

	default:
		return "", errors.New("Nenhuma valor encontrado")
	}
}

func printSliceArray(compare reflect.Value, list reflect.Value) (interface{}, error) {

	for i := 0; i < list.Len(); i++ {
		for x := 0; x < compare.Len(); x++ {
			elem1 := list.Index(i).Interface()
			elem2 := compare.Index(x).Interface()

			if reflect.DeepEqual(elem1, elem2) {
				// Elements are equal, do something here if needed
				mySlice = append(mySlice, elem2)
			}
		}
	}

	length := len(mySlice)

	if length == 0 {
		return mySlice, errors.New("nenhum item encontrado")
	}

	return Return(len(mySlice), mySlice)

}

func printMapVArray(compare reflect.Value, list reflect.Value) (interface{}, error) {

	mapKeys := list.MapKeys()
	mapKeysC := compare.MapKeys()
	for _, mapKey := range mapKeys {
		for _, mapKeyc := range mapKeysC {
			if reflect.DeepEqual(mapKey.Interface(), mapKeyc.Interface()) {
				mySlice = append(mySlice, list.MapIndex(mapKey).Interface())
			}
		}
	}

	return Return(len(mySlice), mySlice)
}

func printStructArray(compare reflect.Value, list reflect.Value) (interface{}, error) {

	for i := 0; i < list.NumField(); i++ {
		field := list.Field(i)
		for x := 0; x < compare.NumField(); x++ {
			field2 := compare.Field(x)
			if reflect.DeepEqual(field.Interface(), field2.Interface()) {
				mySlice = append(mySlice, field.Interface())
			}

		}

	}

	return Return(len(mySlice), mySlice)

}

func Return(length int, mySlice interface{}) (interface{}, error) {
	if length == 0 {
		return mySlice, errors.New("nenhum item encontrado")
	}

	return mySlice, nil
}

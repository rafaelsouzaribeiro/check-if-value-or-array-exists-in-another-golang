package array

import (
	"errors"
	"fmt"
	"reflect"
)

type MapSliceConstraint[K comparable, V interface{}] interface {
	map[K]V | []V
}

/*
count := GetTamanho[[]string, string, string]([]string{"rafael", "Paulo"}) -> Slice
count := GetTamanho[map[string]string, string, string](map[string]string{"rafael": "1", "Paulo": "2"}) -> map
count := GetTamanho[map[int]string, int, string](map[int]string{1: "1", 2: "2"})
*/

func GetTamanho[T MapSliceConstraint[K, V], K comparable, V interface{}](collection T) int {
	return len(collection)
}

/*

Array has value

*/

func ArrayHasValue(value interface{}, data interface{}) (string, error) {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		// If it's a slice
		return printSliceValue(value, reflect.ValueOf(data))

	case reflect.Map:
		// If it's a map
		return printMapValue(value, reflect.ValueOf(data))

	case reflect.Struct:
		// If it's a struct
		return printStructValue(value, reflect.ValueOf(data))
	case reflect.Array:

		// if it's a array
		return printSliceValue(value, reflect.ValueOf(data))

	default:
		return "", errors.New("Nenhuma valor encontrado")
	}

}

/*

Array key Exists

*/

func ArrayKeyExists(key interface{}, data interface{}) (string, error) {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		// If it's a slice
		return printSlice(key, reflect.ValueOf(data))

	case reflect.Map:
		// If it's a map
		return printMap(key, reflect.ValueOf(data))

	case reflect.Struct:
		return printStruct(key, reflect.ValueOf(data))
	case reflect.Array:
		// If it's a array
		return printSlice(key, reflect.ValueOf(data))
	default:
		return "", errors.New("Nenhuma chave encontrada")
	}
}

func printSlice(key interface{}, list reflect.Value) (string, error) {
	for i := 0; i < list.Len(); i++ {
		if i == key {
			return fmt.Sprintf("%v", key), nil
		}
	}

	return "", errors.New("nenhuma chave encontrada")
}

func printMap(key interface{}, list reflect.Value) (string, error) {
	mapKeys := list.MapKeys()
	for _, mapKey := range mapKeys {
		if reflect.DeepEqual(mapKey.Interface(), key) {
			return fmt.Sprintf("%v", key), nil
		}
	}

	return "", errors.New("nenhuma chave encontrada")
}

func printStruct(key interface{}, list reflect.Value) (string, error) {
	for i := 0; i < list.NumField(); i++ {
		field := list.Field(i)
		if reflect.DeepEqual(field.Interface(), key) {
			return fmt.Sprintf("%v", key), nil
		}
	}

	return "", errors.New("nenhuma chave encontrada")
}

func printSliceValue(value interface{}, list reflect.Value) (string, error) {
	for i := 0; i < list.Len(); i++ {
		if reflect.DeepEqual(list.Index(i).Interface(), value) {
			return fmt.Sprintf("%v", list.Index(i).Interface()), nil
		}
	}

	return "", errors.New("Nenhuma valor encontrado")
}

func printMapValue(value interface{}, list reflect.Value) (string, error) {
	mapKeys := list.MapKeys()
	for _, key := range mapKeys {
		if reflect.DeepEqual(list.MapIndex(key).Interface(), value) {
			return fmt.Sprintf("%v", list.MapIndex(key).Interface()), nil
		}
	}

	return "", errors.New("Nenhuma valor encontrado")
}

func printStructValue(value interface{}, list reflect.Value) (string, error) {
	for i := 0; i < list.NumField(); i++ {
		field := list.Field(i)
		if reflect.DeepEqual(field.Interface(), value) {
			return fmt.Sprintf("%v", field.Interface()), nil
		}
	}
	return "", errors.New("Nenhuma valor encontrado")
}

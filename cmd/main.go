package main

import (
	"github.com/rafaelsouzaribeiro/lab-utils/pkg/utils/array"
)

func main() {
	arr := []string{"1.1", "1.2"}
	//maps := map[interface{}]interface{}{"1": 1, 1.5: 1.6, "4": 3}
	// count := GetTamanho[map[int]string, int, string](map[int]string{1: "1", 2: "2"})
	// println(count)
	//var k interface{}
	//k = 1
	//result := Result{Value: "valor", Type: "TIPO", TypeFull: "TIPOFULL"}
	sucess, err := array.ArrayHasValue("1.1", arr)

	if err != nil {
		panic(err)
	} else {
		println(sucess)
	}

	// result := GetValueType(map[string]int{"rafa": 1, "bbb": 2})
	// fmt.Printf("Type: %s,Type Inteiro: %s, Value: %v\n", result.Type, result.TypeFull, result.Value)
}

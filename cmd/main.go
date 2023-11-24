package main

import (
	"github.com/rafaelsouzaribeiro/lab-utils/pkg/utils/array"
)

func main() {
	//maps := map[interface{}]interface{}{"1": 1, 1.5: 1.6, "4": 3}
	count := array.GetTamanho[map[int]string, int, string](map[int]string{1: "1", 2: "2"})
	println(count)

	arr := []string{"1.1", "1.2"}
	sucess, err := array.ArrayHasValue("1.3", arr)

	if err != nil {
		panic(err)
	}

	println(sucess)
	
}

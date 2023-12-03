package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/lab-utils/pkg/utils/array"

	"github.com/rafaelsouzaribeiro/lab-utils/pkg/utils/arrays"
)

type Arrays struct {
	Nome  string
	Idade int
}

func main() {

	//maps := map[interface{}]interface{}{"1": 1, 1.5: 1.6, "4": 3}
	// count := array.GetTamanho[map[int]string, int, string](map[int]string{1: "1", 2: "2"})
	// println(count)

	// arr := []string{"1.1", "1.2"}
	// sucess, err := array.ArrayHasValue("1.2", arr)

	// if err != nil {
	// 	panic(err)
	// }

	// println(sucess)
	// array := map[int]string{1: "1", 3: "2", 5: "3"}
	// valores := map[int]string{5: "2"}
	client := Arrays{
		Nome:  "Rafael",
		Idade: 37,
	}

	clientC := Arrays{
		Nome:  "Paulo",
		Idade: 37,
	}

	res, er := arrays.ArrayInsideArray(clientC, client)

	if er != nil {
		panic(er)
	}

	fmt.Println(res)

	var g [1]int
	g[0] = 1
	var t [3]int
	t[0] = g[0]
	t[1] = 2
	resp, err := arrays.ArrayInsideArray(t, g)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

	var r [2]string
	r[0] = "1"
	r[1] = "2"

	sucess, err := array.ArrayKeyExists(0, r)

	if err != nil {
		panic(err)
	}

	fmt.Println(sucess)

	var r2 [2]string
	r2[0] = "1"
	r2[1] = "2"

	sucess2, err2 := array.ArrayHasValue("1", r2)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println(sucess2)

}

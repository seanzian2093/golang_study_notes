package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}


func main() {
	// Marshal returns JSON encoding of provided Go value
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(33)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(33.33)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	// slcB is a list of integers corresponding to the characters.
	fmt.Println(slcB)
	// convert to string
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "peach": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(mapB)
	fmt.Println(string(mapB))

	res1D := &response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(res1B)
	fmt.Println(string(res1B))

	// use tag on struct field to customrize encoded JSON key names
	res2D := &response2{Page: 2, Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(res2B)
	// keys are page, fruits, intead of Page, Fruits
	fmt.Println(string(res2B))


	// use Unmarshal to decode JSON to Go value
	// a raw string cast to byte
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// a map type to string-to-any type to hold decoded raw value
	var dat map[string]interface{}
	// always handle error
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// convert raw value to appropriate type
	num := dat["num"].(float64)
	fmt.Println(num)
	// nested data requires more conversion
	// cast to a slice of interface
	strs := dat["strs"].([]interface{})
	fmt.Println(strs)
	// []interface{} can not be cast to string directly
	// fmt.Println(string(strs))
	// get 1st element and cast to string
	strs1 := strs[0].(string)
	fmt.Println(strs1)

	// Decode JSON to a specific type
	str := `{"page":2,"fruits":["apple","peach","pear"]}`
	// a empty value of target type
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page     int
	Cuisines []string
}

type response2 struct {
	Page     int      `json:"page"`
	Cuisines []string `json:"cuisines"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(3)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(20.12)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("dog")
	fmt.Println(string(strB))

	slcD := []string{"italian", "mexican", "thai", "romanian", "french"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"italian": 1, "mexican": 2}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{Page: 1, Cuisines: []string{"italian", "mexican", "thai", "romanian", "french"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{Page: 1, Cuisines: []string{"italian", "mexican", "thai", "romanian", "french"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]any

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]any)
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page":1,"cuisines":["italian","mexican","thai","romanian","french"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Cuisines[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"italian": 1, "mexican": 2}
	enc.Encode(d)

	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}

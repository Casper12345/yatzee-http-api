package main

import (
	"encoding/json"
	"fmt"
)

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type response1 struct {
	IsOld bool `json:"old"`
	Name string `json:"name"`
}

func main() {

	a, _ := json.Marshal(true)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}

	str1 := `{"name": "jerry", "old": true}`

	res2 := response1{}

	json.Unmarshal([]byte(str1), &res2)

	json.Unmarshal([]byte(str), &res)

	fmt.Println(res)
	fmt.Println(res2.Name)
	fmt.Println(res2.IsOld)
	fmt.Println(res.Fruits[0])
	fmt.Println(res.Page + 2)

	fmt.Println(string(a))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)

	fmt.Println(string(mapB))

}
package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	var data Data
	data = Data{A: 1, B: 1}
	fmt.Print(data)
	res, _ := json.Marshal(&data)
	fmt.Print(string(res))
}

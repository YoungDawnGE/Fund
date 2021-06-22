package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := "[[\"20180808\",\"1.0000\"],[\"20180809\",\"0.9999\"]]"

	arrayObj := make([][]string, 2)

	err := json.Unmarshal([]byte(str), &arrayObj)
	if err != nil {
		panic(err)
	}

	fmt.Println(arrayObj)

}

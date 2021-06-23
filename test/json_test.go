package main

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func Test01(t *testing.T) {
	//str := "[[\"20180808\",\"1.0000\"],[\"20180809\",\"0.9999\"]]"
	//
	//arrayObj := make([][]string, 2)
	//
	//err := json.Unmarshal([]byte(str), &arrayObj)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(arrayObj)

	matchString, err := regexp.MatchString("\\b\\d{6}\\b", "000230")
	if err != nil {
		fmt.Println(err)
		fmt.Println(matchString)
	}
	fmt.Println(matchString)

}

func Test02(t *testing.T) {
	v, _ := strconv.ParseInt("21", 10, 32)
	fmt.Println(v)
}

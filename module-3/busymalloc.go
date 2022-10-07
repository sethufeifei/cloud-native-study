package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	var myMap = make(map[int]string)
	index := 0
	for {
		index = index + 1
		time.Sleep(time.Second * 1)

		marshal, _ := json.Marshal(myMap)
		value := string(marshal)

		myMap[index] = value
		fmt.Println(myMap)
	}
}

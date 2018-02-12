package main

import (
	"fmt"

	"github.com/shanshanliwei/gogogo/input"
)

func main() {
	dataAsStr, err := input.GetFileStringByPath("./file")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dataAsStr)

}

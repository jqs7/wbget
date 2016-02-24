package main

import (
	"fmt"

	"github.com/jqs7/wbget"
)

func main() {
	wb, _ := getwb.Get("1910069117")
	fmt.Println(wb.Name)
	for _, v := range wb.Posts {
		fmt.Printf("%+v\n", v)
	}
}

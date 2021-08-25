package main

import (
	"fmt"

	"github.com/meanhide/nomadcoder/ch2.bankAndDict/myDict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}

package main

import (
	"fmt"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func main() {
	str := "风和味"
	pattern := "风 味"
	fmt.Print(fuzzy.Match(pattern, str))
}

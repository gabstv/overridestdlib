package main

import (
	"github.com/gabstv/overridestdlib/internal"
)

func main() {
	x := internal.Explode("123456", 6)
	for _, v := range x {
		println(v)
	}
}

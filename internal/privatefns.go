package internal

import (
	_ "strings"
	_ "unsafe"
)

// The instruction below tells the go compiler (and linker) that this function signature
// is actually the private function "strings.explode()"

//go:linkname strings_explode strings.explode

func strings_explode(s string, n int) []string // strings/strings.go

func Explode(s string, n int) []string {
	return strings_explode(s, n)
}

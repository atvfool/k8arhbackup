package option

import "strings"

type Option struct {
	Key   string
	Value string
}

func New(input string) *Option {
	kvp := strings.Split(input, "=")
	o := &Option{Key: kvp[0], Value: kvp[1]}
	return o
}

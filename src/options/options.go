package options

import (
	"k8arh/backup/option"
	"strings"
)

type Options struct {
	rawOptions string
	options    []option.Option
}

func NewFromString(input string) *Options {
	opts := &Options{rawOptions: input}
	lines := strings.Split(input, "\r\n")
	for _, v := range lines {
		newOpt := option.New(v)
		opts.options = append(opts.options, *newOpt)
	}
	return opts
}

func NewFromByte(input byte) *Options {
	opts := NewFromString(string(input))
	return opts
}

func (o Options) GetOptions() []option.Option {
	return o.options
}

func (o Options) GetValueByKey(key string) []string {
	opts := []string{}

	for _, v := range o.options {
		if v.Key == key {
			opts = append(opts, strings.TrimSpace(v.Value))
		}
	}

	return opts
}

func (o Options) GetKeys() []string {
	keys := []string{}

	for _, v := range o.options {
		keys = append(keys, v.Key)
	}

	return keys
}

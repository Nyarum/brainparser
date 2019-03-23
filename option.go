package brainparser

import "encoding/binary"

type subOption func(*AdditionalOption)

type AdditionalOption struct {
	Order binary.ByteOrder
	Len   int
	Skip  int
}

type Option struct {
	Field      string
	Value      interface{}
	SubOptions []subOption
	*AdditionalOption
}

func SetOrder(v binary.ByteOrder) func(opt *AdditionalOption) {
	return func(opt *AdditionalOption) {
		opt.Order = v
	}
}

func SetLen(v int) func(opt *AdditionalOption) {
	return func(opt *AdditionalOption) {
		opt.Len = v
	}
}

func SetSkip(v int) func(opt *AdditionalOption) {
	return func(opt *AdditionalOption) {
		opt.Skip = v
	}
}

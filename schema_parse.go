package brainparser

import (
	"encoding/binary"
)

/*
	var sch = brainparser.Schema{
		{"ID", t.ID, [brainparser.SetOrder(binary.LittleEndian)], nil},
		{"ID2", t.ID2, nil, nil},
		{"ID3", t.ID3, nil, nil},
		{"ID4", t.ID4, nil, nil},
	}
*/

// Schema generic field that cover all of types with options
type Schema []*Option

type schemaParse struct {
}

func newSchemaParse() schemaParse {
	return schemaParse{}
}

func (s *schemaParse) Configure(schema Schema) Schema {
	for _, opt := range schema {
		if opt.AdditionalOption == nil {
			opt.AdditionalOption = &AdditionalOption{}
		}

		for _, subOpt := range opt.SubOptions {
			subOpt(opt.AdditionalOption)
		}

		if opt.Order == nil {
			opt.Order = binary.LittleEndian
		}
	}

	return schema
}

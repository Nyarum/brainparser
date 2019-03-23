package brainparser

import (
	"encoding/binary"
)

/*
	var sch = brainparser.Schema{
		"id": {
			"value": uint16(1),
		},
		"name": {
			"value": string("default"),
			"order": binary.LittleEndian,
			"len": 15,
		},
		"attrs": {
			"value": []uint16{1,2},
			"len": 2,
			"skip": 2,
		}
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

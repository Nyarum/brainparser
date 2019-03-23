package brainparser

import "bytes"

type BrainParser struct {
	schParse schemaParse
}

func NewBrainParser() BrainParser {
	return BrainParser{}
}

func (b BrainParser) Parse(buf []byte, pkt Packet) error {
	sch := pkt.Schema()

	resParse, err := newBytesParse(b.schParse.Configure(sch)).
		Decode(bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	pkt.SetFields(resParse)

	return nil
}

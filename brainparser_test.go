package brainparser_test

import (
	"testing"

	"github.com/Nyarum/brainparser"
)

type TestPacket struct {
	ID  uint8
	ID2 uint16
	ID3 uint32
	ID4 uint64
}

func (t *TestPacket) SetFields(v map[string]interface{}) {
	t.ID = v["ID"].(uint8)
	t.ID2 = v["ID2"].(uint16)
	t.ID3 = v["ID3"].(uint32)
	t.ID4 = v["ID4"].(uint64)
}

func (t *TestPacket) Default() {
	t.ID = 1
	t.ID2 = 3
	t.ID3 = 4
	t.ID4 = 6
}

func (t TestPacket) Schema() brainparser.Schema {
	return brainparser.Schema{
		{"ID", t.ID, nil, nil},
		{"ID2", t.ID2, nil, nil},
		{"ID3", t.ID3, nil, nil},
		{"ID4", t.ID4, nil, nil},
	}
}

func TestBrainParser(t *testing.T) {
	bp := brainparser.NewBrainParser()
	tp := &TestPacket{}

	err := bp.Parse([]byte{0x05, 0x2B, 0x00, 0x21, 0x00, 0x00, 0x00, 0x2C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, tp)
	if err != nil {
		t.Error(err)
	}

	t.Log(tp)
}

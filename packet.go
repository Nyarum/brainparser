package brainparser

type Packet interface {
	SetFields(map[string]interface{})
	Default()
	Schema() Schema
}

package brainparser

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type bytesParse struct {
	sch Schema
}

var (
	ErrUnefficientBytes = errors.New("Can't parse value from allocated data")
)

func newBytesParse(sch Schema) bytesParse {
	return bytesParse{
		sch: sch,
	}
}

func (b bytesParse) getUint8(data []byte, opt *Option) uint8 {
	return uint8(data[0])
}

func (b bytesParse) getUint16(data []byte, opt *Option) uint16 {
	if opt.Order == binary.LittleEndian {
		return binary.LittleEndian.Uint16(data)
	}

	return binary.BigEndian.Uint16(data)
}

func (b bytesParse) getUint32(data []byte, opt *Option) uint32 {
	if opt.Order == binary.LittleEndian {
		return binary.LittleEndian.Uint32(data)
	}

	return binary.BigEndian.Uint32(data)
}

func (b bytesParse) getUint64(data []byte, opt *Option) uint64 {
	if opt.Order == binary.LittleEndian {
		return binary.LittleEndian.Uint64(data)
	}

	return binary.BigEndian.Uint64(data)
}

func (b bytesParse) Decode(buf *bytes.Buffer) (map[string]interface{}, error) {
	res := make(map[string]interface{}, len(b.sch))

	for _, opt := range b.sch {
		switch opt.Value.(type) {
		case uint8:
			if buf.Len() < 1 {
				return nil, ErrUnefficientBytes
			}

			res[opt.Field] = b.getUint8(buf.Next(1), opt)
		case uint16:
			if buf.Len() < 2 {
				return nil, ErrUnefficientBytes
			}

			res[opt.Field] = b.getUint16(buf.Next(2), opt)
		case uint32:
			if buf.Len() < 4 {
				return nil, ErrUnefficientBytes
			}

			res[opt.Field] = b.getUint32(buf.Next(4), opt)
		case uint64:
			if buf.Len() < 8 {
				return nil, ErrUnefficientBytes
			}

			res[opt.Field] = b.getUint64(buf.Next(8), opt)
		}
	}

	return res, nil
}

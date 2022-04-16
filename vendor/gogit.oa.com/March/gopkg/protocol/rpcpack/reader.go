package rpcpack

import (
	"bytes"
	"encoding/binary"
)

type Reader struct {
	cmd  uint8
	buff *bytes.Buffer
}

func NewReader(cmd uint8, buff []byte) *Reader {
	return &Reader{
		cmd:  cmd,
		buff: bytes.NewBuffer(buff),
	}
}

func (r *Reader) GetCmd() uint8 {
	return r.cmd
}

func (r *Reader) Byte() uint8 {
	b, _ := r.buff.ReadByte()
	return b
}

func (r *Reader) Short() int16 {
	var v int16
	binary.Read(r.buff, binary.BigEndian, &v)
	return v
}

func (r *Reader) Int() int32 {
	var v int32
	binary.Read(r.buff, binary.BigEndian, &v)
	return v
}

func (r *Reader) Int64() int64 {
	var v int64
	binary.Read(r.buff, binary.BigEndian, &v)
	return v
}

func (r *Reader) String() string {

	var l uint32
	binary.Read(r.buff, binary.LittleEndian, &l)
	if l <= 0 {
		return ""
	}
	data := make([]byte, l)
	r.buff.Read(data)
	return string(data[:l-1])
}

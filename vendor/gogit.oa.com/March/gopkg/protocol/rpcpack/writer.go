package rpcpack

import (
	"bytes"
	"encoding/binary"
)

type Writer struct {
	cmd  uint8
	buff *bytes.Buffer
}

func NewWriter(cmd uint8) *Writer {
	return &Writer{
		cmd:  cmd,
		buff: new(bytes.Buffer),
	}
}

func (w *Writer) header(l uint32) {
	w.buff.Write([]byte("RPC"))
	binary.Write(w.buff, binary.BigEndian, w.cmd)
	binary.Write(w.buff, binary.BigEndian, packetVer)
	binary.Write(w.buff, binary.BigEndian, packetAck)
	binary.Write(w.buff, binary.BigEndian, l)
}

func (w *Writer) End() {
	buf := make([]byte, w.buff.Len())
	copy(buf, w.buff.Bytes())
	w.buff.Reset()
	w.header(uint32(len(buf)))
	if len(buf) > 0 {
		w.buff.Write(buf)
	}
}

func (w *Writer) Byte(v byte) {
	binary.Write(w.buff, binary.BigEndian, v)
}

func (w *Writer) Short(v int16) {
	binary.Write(w.buff, binary.BigEndian, v)
}

func (w *Writer) Int(v int32) {
	binary.Write(w.buff, binary.BigEndian, v)
}

func (w *Writer) Int64(v int64) {
	binary.Write(w.buff, binary.BigEndian, v)
}

func (w *Writer) String(v string) {
	binary.Write(w.buff, binary.LittleEndian, uint32(len(v))+1)
	w.buff.WriteString(v)
	w.buff.WriteByte(0)
}

func (w *Writer) GetBuffer() []byte {
	return w.buff.Bytes()
}

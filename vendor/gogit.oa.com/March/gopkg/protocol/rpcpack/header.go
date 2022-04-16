package rpcpack

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	HeaderSize   int    = 13
	packetVer    uint8  = 1
	packetAck    uint32 = 1
	maxBufferLen uint32 = 2 << 12
)

var (
	ErrIdentify   = errors.New("invalid identify")
	ErrVer        = errors.New("invalid ver")
	ErrAck        = errors.New("invalid ack")
	ErrCmd        = errors.New("invalid cmd")
	ErrBufferSize = errors.New("invalid buffer size")
)

type Header struct {
	Identify [3]byte //3
	Cmd      uint8   //1
	Ver      uint8   //1
	Ack      uint32  //4
	Size     uint32  //4
}

func (h *Header) GetSize() uint32 {
	return h.Size
}

func (h *Header) GetCmd() uint8 {
	return h.Cmd
}

func NewHeader(buf []byte) (*Header, error) {
	if len(buf) != HeaderSize {
		return nil, ErrBufferSize
	}
	h := new(Header)
	if err := binary.Read(bytes.NewBuffer(buf), binary.BigEndian, h); err != nil {
		return nil, err
	}
	if err := h.Check(); err != nil {
		return nil, err
	}
	return h, nil
}

func (h *Header) Check() error {
	if h.Identify[0] != 'R' && h.Identify[1] != 'P' && h.Identify[2] != 'C' {
		return fmt.Errorf("header %+v error %w", h, ErrIdentify)
	}
	if h.Ver != packetVer {
		return fmt.Errorf("header %+v error %w", h, ErrVer)
	}
	if h.Ack != packetAck {
		return fmt.Errorf("header %+v error %w", h, ErrAck)
	}
	if h.Cmd <= 0 || h.Cmd >= 255 {
		return fmt.Errorf("header %+v error %w", h, ErrCmd)
	}
	if h.Size > maxBufferLen {
		return fmt.Errorf("buffer len is too long %d", h.Size)
	}
	return nil
}

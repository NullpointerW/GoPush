package protocol

import (
	"bytes"
	"encoding/binary"
	"gopush/utils"
)

var heartLen = 2

func Pack(msg string) (b []byte) {
	return utils.Scb(msg + string(EndFlag))
}

func PackByteStream(contextLen uint16, b []byte) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, heartLen+len(b)))
	binary.Write(buf, binary.BigEndian, contextLen)
	buf.Write(b)
	return buf.Bytes()
}

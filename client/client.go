package client

import (
	"fmt"
	"gogit.oa.com/March/gopkg/protocol/rpcpack"
	"gogit.oa.com/March/gopkg/util"
	"io"
	"net"
)

const (
	ADDR = "192.168.97.55:6766"           //服务地址
	API  = "mtt.getAvailableUserTool"     //方法
	ARGS = `{"mid":260689,"ddcard":1676}` //参数

)

func rpcRequest() []byte {
	w := rpcpack.NewWriter(0x1)
	w.String(API)
	w.String(ARGS)
	w.End()
	return w.GetBuffer()
}

func Run() {
	reader := SendAndRecv(rpcRequest())
	// cmd := reader.GetCmd()
	code := reader.Byte()
	json := reader.String()
	// fmt.Println("cmd:", cmd)
	fmt.Println("code:", code)
	fmt.Println("json:", json)
}

func SendAndRecv(data []byte) *rpcpack.Reader {

	//请求
	conn := Conn()
	_, err := conn.Write(data)
	util.MustNil(err)

	//读头部
	hb := make([]byte, rpcpack.HeaderSize)
	_, err = io.ReadFull(conn, hb)
	util.MustNil(err)
	header, err := rpcpack.NewHeader(hb)
	util.MustNil(err)

	//读主体
	body := make([]byte, header.GetSize())
	n, err := io.ReadFull(conn, body)
	util.MustNil(err)

	return rpcpack.NewReader(0x10, body[:n])
}

func Conn() net.Conn {
	conn, err := net.Dial("tcp", ADDR)
	util.MustNil(err)
	return conn
}

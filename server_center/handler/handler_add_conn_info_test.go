package handler

import (
	"bytes"
	"io"
	"testing"

	"github.com/gochenzl/chess/pb/center"
	"github.com/gochenzl/chess/server_center/conn_info"
	"github.com/gochenzl/chess/util/rpc"
	"google.golang.org/protobuf/proto"
)

func TestHandleAddConnInfo(t *testing.T) {
	conn_info.InitTest()

	var clients []io.ReadWriter
	clients = append(clients, &bytes.Buffer{})
	clients = append(clients, &bytes.Buffer{})

	for i := 0; i < len(clients); i++ {
		addClient(clients[i])
	}

	req := &center.AddConnInfoReq{Info: &center.ConnInfo{Userid: 10000, Gateid: 1, Connid: 1}}
	client := &bytes.Buffer{}
	addClient(client)
	HandleAddConnInfo(client, req)

	pb, err := rpc.DecodePb(client)
	if err != nil {
		t.Errorf("decode resp:%s", err.Error())
		return
	}
	if proto.MessageName(pb) != "center.AddConnInfoResp" {
		t.Errorf("invalid response:%s", proto.MessageName(pb))
	}

	for i := 0; i < len(clients); i++ {
		pb, err := rpc.DecodePb(clients[i])
		if err != nil {
			t.Errorf("decode resp:%s", err.Error())
			return
		}

		if proto.MessageName(pb) != "center.NewConnInfoNotify" {
			t.Errorf("invalid response:%s", proto.MessageName(pb))
		}
	}

	if !conn_info.Exist(center.ConnInfo{Userid: 10000, Gateid: 1, Connid: 1}) {
		t.Errorf("add conn info fail")
	}

	req = &center.AddConnInfoReq{Info: &center.ConnInfo{Userid: 20000, Gateid: 1, Connid: 1}}
	HandleAddConnInfo(client, req)

	for i := 0; i < len(clients); i++ {
		pb, err := rpc.DecodePb(clients[i])
		if err != nil {
			t.Errorf("decode resp:%s", err.Error())
			return
		}

		if proto.MessageName(pb) != "center.DelConnInfoNotify" {
			t.Errorf("invalid response:%s", proto.MessageName(pb))
		}
	}

	for i := 0; i < len(clients); i++ {
		pb, err := rpc.DecodePb(clients[i])
		if err != nil {
			t.Errorf("decode resp:%s", err.Error())
			return
		}

		if proto.MessageName(pb) != "center.NewConnInfoNotify" {
			t.Errorf("invalid response:%s", proto.MessageName(pb))
		}
	}

	if !conn_info.Exist(center.ConnInfo{Userid: 20000, Gateid: 1, Connid: 1}) {
		t.Errorf("add conn info fail")
	}
}

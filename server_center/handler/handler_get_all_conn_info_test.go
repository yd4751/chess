package handler

import (
	"bytes"

	"github.com/gochenzl/chess/pb/center"
	"github.com/gochenzl/chess/server_center/conn_info"
	"github.com/gochenzl/chess/util/rpc"

	"testing"
)

func TestHandleGetAllConnInfo(t *testing.T) {
	conn_info.InitTest()
	var connInfos []center.ConnInfo
	connInfos = append(connInfos, center.ConnInfo{Userid: 10000, Gateid: 1, Connid: 1})
	connInfos = append(connInfos, center.ConnInfo{Userid: 20000, Gateid: 1, Connid: 2})
	connInfos = append(connInfos, center.ConnInfo{Userid: 30000, Gateid: 2, Connid: 1})
	connInfos = append(connInfos, center.ConnInfo{Userid: 40000, Gateid: 2, Connid: 2})

	for i := 0; i < len(connInfos); i++ {
		conn_info.Add(connInfos[i])
	}

	req := &center.GetAllConnInfoReq{}
	client := &bytes.Buffer{}
	HandleGetAllConnInfo(client, req)

	pb, err := rpc.DecodePb(client)
	if err != nil {
		t.Errorf("decode resp:%s", err.Error())
		return
	}
	resp := pb.(*center.GetAllConnInfoResp)
	if resp == nil {
		t.Errorf("invalid resp")
	}

	if len(resp.Infos) != len(connInfos) {
		t.Errorf("resp error %d", len(resp.Infos))
	}

	for i := 0; i < len(resp.Infos); i++ {
		var find bool
		for j := 0; j < len(connInfos); j++ {
			if resp.Infos[i].String() == connInfos[j].String() {
				find = true
				break
			}
		}

		if !find {
			t.Errorf("resp error")
		}
	}
}

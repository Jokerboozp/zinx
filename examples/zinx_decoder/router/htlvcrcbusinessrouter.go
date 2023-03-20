package router

import (
	"github.com/aceld/zinx/zdecoder"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

type HtlvCrcBusinessRouter struct {
	znet.BaseRouter
}

func (this *HtlvCrcBusinessRouter) Handle(request ziface.IRequest) {
	zlog.Ins().DebugF("Call HtlvCrcBusinessRouter Handle %d %+v\n", request.GetMessage().GetMsgID(), request.GetMessage().GetData())
	msgID := request.GetMessage().GetMsgID()
	if msgID == 0x10 {
		_response := request.GetResponse()
		if _response != nil {
			switch _response.(type) {
			case zdecoder.HtlvCrcData:
				tlvData := _response.(zdecoder.HtlvCrcData)
				zlog.Ins().DebugF("do msgid=0x10 data business %+v\n", tlvData)
			}
		}
	}
}

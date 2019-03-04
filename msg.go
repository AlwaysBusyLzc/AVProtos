package msg

import (
	"server/messages"
	"server/network/protobuf"
)

var Processor = protobuf.NewProcessor()

const (
	Gate2LogicConnectSuccessSyn = 0
	Logic2GateRegisterReq = 1
	Gate2LogicRegisterRes = 2
	Gate2LogicStartAsrProcReq = 3
	Logic2GateAsrProcRes = 4
	Gate2LogicAudioData = 5
	Gate2LogicOpenUserReq = 6
	Logic2GateOpenUserRes = 7
	CloseUserReq = 8
	Logic2GateLoadStatus = 9

	// 客户端协议  从1001开始 前面预留给服务器协议
	Dev2GateStartAsrProcReq = 1001
	Gate2DevAsrProcRes = 1002
	Dev2GateAudioData = 1003
	Dev2GateEndAsrReq = 1004

)

func init() {
	// 注册消息
	Processor.Register(&messages.Gate2LogicConnectSuccessSyn{}, Gate2LogicConnectSuccessSyn)
	Processor.Register(&messages.Logic2GateRegisterReq{}, Logic2GateRegisterReq)
	Processor.Register(&messages.Gate2LogicRegisterRes{}, Gate2LogicRegisterRes)
	Processor.Register(&messages.Gate2LogicStartAsrProcReq{}, Gate2LogicStartAsrProcReq)
	Processor.Register(&messages.Logic2GateAsrProcRes{}, Logic2GateAsrProcRes)
	Processor.Register(&messages.Gate2LogicAudioData{}, Gate2LogicAudioData)
	Processor.Register(&messages.Gate2LogicOpenUserReq{}, Gate2LogicOpenUserReq)
	Processor.Register(&messages.Logic2GateOpenUserRes{}, Logic2GateOpenUserRes)
	Processor.Register(&messages.CloseUserReq{}, CloseUserReq)
	Processor.Register(&messages.Logic2GateLoadStatus{}, Logic2GateLoadStatus)

	Processor.Register(&messages.Dev2GateStartAsrProcReq{}, Dev2GateStartAsrProcReq)
	Processor.Register(&messages.Gate2DevAsrProcRes{}, Gate2DevAsrProcRes)
	Processor.Register(&messages.Dev2GateAudioData{}, Dev2GateAudioData)
	Processor.Register(&messages.Dev2GateEndAsrReq{}, Dev2GateEndAsrReq)
}

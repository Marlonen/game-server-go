package internal

import (
	"github.com/name5566/leaf/gate"
)

var agents = make(map[gate.Agent]struct{})

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	// skeleton.RegisterChanRPC("WriteAll", rpcWriteMsgToAll)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	agents[a] = struct{}{}
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	delete(agents, a)
}

// func rpcWriteMsgToAll(args []interface{}) {
// 	for a := range agents {
// 		a.WriteMsg(args[0])
// 	}
// }

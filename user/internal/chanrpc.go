package internal

import (
	"server/gamedata"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("LoginSuccess", rpcLoginSuccess)
}

func rpcNewAgent(args []interface{}) {
	// a := args[0].(gate.Agent)
}

func rpcCloseAgent(args []interface{}) {
	gamedata.RemoveAgent(args[0].(gate.Agent))
}

func rpcLoginSuccess(args []interface{}) {
	session := args[0].(string)
	agent := args[1].(gate.Agent)

	gamedata.SetAgentSession(agent, session)
}

// func rpcWriteMsgToAll(args []interface{}) {
// 	for a := range agents {
// 		a.WriteMsg(args[0])
// 	}
// }

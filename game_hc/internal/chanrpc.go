package internal

import "github.com/name5566/leaf/gate"

var players = make(map[gate.Agent]struct{})

func init() {
	ChanRPC.Register("UserJoin", rpcUserJoin)
	ChanRPC.Register("UserLeave", rpcUserLeave)
}

func rpcUserJoin(args []interface{}) {
	a := args[0].(gate.Agent)
	players[a] = struct{}{}
	sendSceneMsg(a)
}

func rpcUserLeave(args []interface{}) {
	a := args[0].(gate.Agent)
	delete(players, a)
}

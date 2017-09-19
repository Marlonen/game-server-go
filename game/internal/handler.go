package internal

import (
	"reflect"
	"server/game_ba"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

var gameName = make(map[int]string)

var agentGames = make(map[gate.Agent]int)

func init() {
	handler(&msg.C_GameJoin{}, handleGameJoin)

	gameName[GameBA] = "baccarat"
	gameName[GameHC] = "hundred_cow"
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleGameJoin(args []interface{}) {
	m := args[0].(*msg.C_GameJoin)
	a := args[1].(gate.Agent)
	// 查找键值是否存在
	if id, ok := agentGames[a]; ok {
		if id == m.GameID {
			game_ba.ChanRPC.Go("UserJoin", a)
		} else {
			ret := new(msg.S_JoinResult)
			ret.Result = 1
			ret.ErrorMsg = "you are inside of the game " + gameName[GameBA]
			a.WriteMsg(&ret)
		}
	} else {
		switch m.GameID {
		case GameBA:
			agentGames[a] = m.GameID
			game_ba.ChanRPC.Go("UserJoin", a)
		case GameHC:
		}

	}

}

package internal

import (
	"reflect"
	"server/msg"
)

func init() {
	handleMsg(&msg.C_GameLogin{}, handleGameLogin)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleGameLogin(args []interface{}) {
	// m := args[0].(*msg.C_GameJoin)
	// conf.ConnProxy
	// UserData

	// a := args[1].(gate.Agent)
	// // 查找键值是否存在
	// if id, ok := agentGames[a]; ok {
	// 	if id == m.GameID {
	// 		game_ba.ChanRPC.Go("UserJoin", a)
	// 	} else {
	// 		a.WriteMsg(&msg.S_JoinResult{
	// 			Result:   1,
	// 			ErrorMsg: "you are inside of the game " + gameName[GameBA],
	// 		})
	// 	}
	// } else {
	// 	switch m.GameID {
	// 	case GameBA:
	// 		agentGames[a] = m.GameID
	// 		game_ba.ChanRPC.Go("UserJoin", a)
	// 	case GameHC:
	// 	}

	// }

}

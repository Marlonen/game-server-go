package internal

import (
	"reflect"
	"server/msg"
	"server/user"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.C_GameLogin{}, handleGameLogin)
}

func handleGameLogin(args []interface{}) {
	m := args[0].(*msg.C_GameLogin)

	session := m.UserName

	user.ChanRPC.Go("LoginSuccess", session, args[1])
}

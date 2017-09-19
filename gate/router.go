package gate

import (
	"server/game"
	"server/msg"
)

func init() {
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.C_GameJoin{}, game.ChanRPC)
}

package gate

import (
	"server/msg"
	"server/user"
)

func init() {
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.C_GameLogin{}, user.ChanRPC)
}

package msg

func registerMsgCommon() {
	Processor.Register(&C_GameJoin{})
	Processor.Register(&S_JoinResult{})
}

type C_GameJoin struct {
	GameID int
	RoomID int
}

type S_JoinResult struct {
	Result   int
	ErrorMsg string
}

package msg

func registerMsgCommon() {
	Processor.Register(&C_GameJoin{})
	Processor.Register(&C_GameLogin{})
	Processor.Register(&S_HandleResult{})
}

type C_GameJoin struct {
	GameID int
	RoomID int
}

type S_HandleResult struct {
	Result   int
	ErrorMsg string
}

type C_GameLogin struct {
	UserName string
	PassWord string
}

type M_UserData struct {
	GameID int
}

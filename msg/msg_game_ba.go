package msg

func registerMsgGameBA() {
	Processor.Register(&S_BA_SCENE_FREE{})
	Processor.Register(&S_BA_SCENE_CHIP{})
	Processor.Register(&S_BA_SCENE_OPEN{})
	Processor.Register(&S_BA_GAME_FREE{})
	Processor.Register(&S_BA_GAME_CHIP{})
	Processor.Register(&S_BA_GAME_OPEN{})
}

type S_BA_SCENE_FREE struct {
	TimeLeave int
}

type S_BA_SCENE_CHIP struct {
	TimeLeave int
}

type S_BA_SCENE_OPEN struct {
	TimeLeave int
}

type S_BA_GAME_FREE struct {
	TimeLeave int
}

type S_BA_GAME_CHIP struct {
	TimeLeave int
}

type S_BA_GAME_OPEN struct {
	TimeLeave int
}

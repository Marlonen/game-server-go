package internal

import (
	"reflect"
	"server/msg"
	"time"

	"github.com/name5566/leaf/gate"
)

//
const (
	StatusFree = iota
	StatusChip
	StatusOpen
)

//
const (
	TimeFree = 5
	TimeChip = 7
	TimeOpen = 9
)

var (
	currStatus int
	currTime   int
)

func init() {
}

func broadcastTable(message interface{}) {
	for a := range players {
		a.WriteMsg(message)
	}
}

func handleMsg(m interface{}, h interface{}) {
	ChanRPC.Register(reflect.TypeOf(m), h)
}

func startGame() {
	onGameFree()
	timeCountdown()
}

func timeCountdown() {
	currTime--
	skeleton.AfterFunc(1*time.Second, timeCountdown)
}

func onGameFree() {
	currStatus = StatusFree
	currTime = TimeFree

	broadcastTable(&msg.S_BA_GAME_FREE{
		TimeLeave: currTime,
	})
	skeleton.AfterFunc(5*time.Second, onGameChip)
}

func onGameChip() {
	currStatus = StatusChip
	currTime = TimeChip

	broadcastTable(&msg.S_BA_GAME_CHIP{
		TimeLeave: currTime,
	})
	skeleton.AfterFunc(5*time.Second, onGameOpen)
}

func onGameOpen() {
	currStatus = StatusOpen
	currTime = TimeOpen

	broadcastTable(&msg.S_BA_GAME_OPEN{
		TimeLeave: currTime,
	})
	skeleton.AfterFunc(5*time.Second, onGameFree)
}

func sendSceneMsg(a gate.Agent) {
	switch currStatus {
	case StatusFree:
		a.WriteMsg(&msg.S_BA_SCENE_FREE{
			TimeLeave: currTime,
		})
	case StatusChip:
		a.WriteMsg(&msg.S_BA_SCENE_CHIP{
			TimeLeave: currTime,
		})
	case StatusOpen:
		a.WriteMsg(&msg.S_BA_SCENE_OPEN{
			TimeLeave: currTime,
		})
	}
}

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
	// handler(&msg.GameJoin{}, handleGameJoin)
}

func broadcastTable(message interface{}) {
	for a := range players {
		a.WriteMsg(message)
	}
}

func handler(m interface{}, h interface{}) {
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

	var m msg.S_BA_GAME_FREE
	m.TimeLeave = currTime
	broadcastTable(&m)
	skeleton.AfterFunc(5*time.Second, onGameChip)
}

func onGameChip() {
	currStatus = StatusChip
	currTime = TimeChip

	var m msg.S_BA_GAME_CHIP
	m.TimeLeave = currTime
	broadcastTable(&m)
	skeleton.AfterFunc(5*time.Second, onGameOpen)
}

func onGameOpen() {
	currStatus = StatusOpen
	currTime = TimeOpen

	var m msg.S_BA_GAME_OPEN
	m.TimeLeave = currTime
	broadcastTable(&m)
	skeleton.AfterFunc(5*time.Second, onGameFree)
}

func sendSceneMsg(a gate.Agent) {
	switch currStatus {
	case StatusFree:
		var m msg.S_BA_SCENE_FREE
		m.TimeLeave = currTime
		a.WriteMsg(&m)
	case StatusChip:
		var m msg.S_BA_SCENE_CHIP
		m.TimeLeave = currTime
		a.WriteMsg(&m)
	case StatusOpen:
		var m msg.S_BA_SCENE_OPEN
		m.TimeLeave = currTime
		a.WriteMsg(&m)
	}
}

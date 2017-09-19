package internal

import (
	"reflect"
	"server/msg"
	"time"

	"github.com/name5566/leaf/gate"
)

const (
	STATUS_FREE = iota
	STATUS_CHIP
	STATUS_OPEN
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

func onGameFree() {
	currStatus = STATUS_FREE

	var m msg.S_BA_GAME_FREE
	m.TimeLeave = 5
	broadcastTable(&m)
	skeleton.AfterFunc(5*time.Second, onGameChip)
}

func onGameChip() {
	currStatus = STATUS_CHIP

	var m msg.S_BA_GAME_CHIP
	m.TimeLeave = 6
	broadcastTable(&m)
	skeleton.AfterFunc(5*time.Second, onGameOpen)
}

func onGameOpen() {
	currStatus = STATUS_OPEN

	var m msg.S_BA_GAME_OPEN
	m.TimeLeave = 7
	broadcastTable(&m)
	skeleton.AfterFunc(5*time.Second, onGameFree)
}

func sendSceneMsg(a gate.Agent) {
	switch currStatus {
	case STATUS_FREE:
		var m msg.S_BA_SCENE_FREE
		m.TimeLeave = 2
		a.WriteMsg(&m)
	case STATUS_CHIP:
		var m msg.S_BA_SCENE_CHIP
		m.TimeLeave = 3
		a.WriteMsg(&m)
	case STATUS_OPEN:
		var m msg.S_BA_SCENE_OPEN
		m.TimeLeave = 4
		a.WriteMsg(&m)
	}
}

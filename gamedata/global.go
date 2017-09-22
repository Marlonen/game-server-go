package gamedata

import (
	"server/msg"
	"sync"

	"github.com/name5566/leaf/gate"
)

var (
	mu           sync.RWMutex
	userData     = make(map[string]msg.M_UserData)
	agentSession = make(map[gate.Agent]string)
)

func SetAgentSession(agent gate.Agent, session string) {
	mu.RLock()
	defer mu.RUnlock()
	agentSession[agent] = session
	if _, ok := userData[session]; !ok {
		userData[session] = msg.M_UserData{
			GameID: -1,
		}
	}
}

func RemoveAgent(agent gate.Agent) {
	mu.RLock()
	defer mu.RUnlock()
	if _, ok := agentSession[agent]; ok {
		delete(agentSession, agent)
	}
}

func GetUserDataByAgent(agent gate.Agent) (msg.M_UserData, bool) {
	mu.RLock()
	defer mu.RUnlock()
	if _, ok1 := agentSession[agent]; ok1 {
		if _, ok2 := userData[agentSession[agent]]; ok2 {
			return userData[agentSession[agent]], true
		}
	}

	return msg.M_UserData{}, false
}

func GetUserDataBySession(session string) (msg.M_UserData, bool) {
	mu.RLock()
	defer mu.RUnlock()
	if _, ok := userData[session]; ok {
		return userData[session], true
	}
	return msg.M_UserData{}, false
}

package internal

import (
	"server/base"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	log.Debug("game baccarat opening...")

	startGame()
}

func (m *Module) OnDestroy() {

}

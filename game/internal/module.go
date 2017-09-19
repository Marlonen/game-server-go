package internal

import (
	"server/base"
	"server/conf"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

//
const (
	GameBA = 1001
	GameHC = 1002
)

//
var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	log.Debug("tcp %s / ws %s", conf.Server.TCPAddr, conf.Server.WSAddr)
	log.Debug("game kernal server opening...")
}

func (m *Module) OnDestroy() {

}

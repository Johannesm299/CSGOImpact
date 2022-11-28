package match

import "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"

type PlayerWrapper struct {
	Player *common.Player
	Impact float64
}

func (p *PlayerWrapper) Equals(playerCompare *common.Player) bool {
	if (p.Player.UserID) == (playerCompare.UserID) {
		return true
	}
	return false
}

func (p *PlayerWrapper) SetImpact(impact float64) {
	p.Impact = impact
}

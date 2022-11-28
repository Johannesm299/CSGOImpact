package match

import "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"

type KillImpact struct {
	Kill            *events.Kill
	NumEnemiesAlive int
	NumAlliesAlive  int
	killImpact      float64
}

func (k *KillImpact) KillImpact() float64 {
	return k.killImpact
}

func (k *KillImpact) CalculateImpact() {
	k.killImpact = float64(k.NumEnemiesAlive) / float64(k.NumAlliesAlive)
}

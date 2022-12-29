package match

import (
	"fmt"

	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
)

type RoundImpact struct {
	Team1          [5]PlayerWrapper
	SumTeam1Impact float64
	Team2          [5]PlayerWrapper
	SumTeam2Impact float64
	Kills          []KillImpact
}

func (r *RoundImpact) ResetImpact(team int) {
	for i := range r.Team1 {
		if team == -1 {
			r.Team1[i].SetImpact(0)
			r.Team2[i].SetImpact(0)
		} else if team == 1 {
			r.Team1[i].SetImpact(0)
		} else if team == 2 {
			r.Team2[i].SetImpact(0)
		}

	}
	r.SumTeam1Impact = 0
	r.SumTeam2Impact = 0
}

func (r *RoundImpact) ResetKills() {
	r.Kills = []KillImpact{}
}

func (r *RoundImpact) AddImpact(player *common.Player, impact float64) {
	for i := range r.Team1 {
		if r.Team1[i].Equals(player) {
			r.Team1[i].Impact += impact
			r.SumTeam1Impact += impact
		}
		if r.Team2[i].Equals(player) {
			r.Team2[i].Impact += impact
			r.SumTeam2Impact += impact
		}
	}
}

func (r *RoundImpact) NormalizeImpact() {
	for i := range r.Team1 {
		r.Team1[i].Impact = r.Team1[i].Impact / r.SumTeam1Impact
		r.Team2[i].Impact = r.Team2[i].Impact / r.SumTeam2Impact
	}
}

func (r *RoundImpact) ValidateNormalizedImpact() {
	sumImpact := 0.0
	for i := range r.Team1 {
		sumImpact += r.Team1[i].Impact
		sumImpact += r.Team2[i].Impact
	}
	fmt.Printf("\tRound has sum of normalized Impacts of %.6f\n", sumImpact)
}

package calc

import (
	"fmt"
	"log"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
	common "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	events "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"

	"github.com/Johannesm299/CSGOImpact/pkg/match"
)

func SumAlive(p []*common.Player) int {
	sumAlive := 0
	for _, e := range p {
		if e.IsAlive() {
			sumAlive++
		}
	}
	return sumAlive
}

func CalculateImpact(filepath string) {
	f, err := os.Open(filepath) //os.Open("../match730_003403354556619293004_0695644917_191.dem")
	if err != nil {
		log.Panic("failed to open demo file: ", err)
	}
	defer f.Close()

	p := dem.NewParser(f)
	defer p.Close()

	var CurMatch match.MatchImpact
	var CurRound match.RoundImpact

	p.RegisterEventHandler(func(e events.IsWarmupPeriodChanged) {
		if !e.NewIsWarmupPeriod {
			fmt.Printf("\n\n ## Match Starting ## \n\n")
			TeamTerrorists := p.GameState().TeamTerrorists().Members()
			TeamCounterTerrorists := p.GameState().TeamCounterTerrorists().Members()

			for i := 0; i < 5; i++ {
				CurRound.Team1[i] = match.PlayerWrapper{Player: TeamTerrorists[i], Impact: 0}
				fmt.Printf("Added Player %s to team1\n", TeamTerrorists[i])
				CurRound.Team2[i] = match.PlayerWrapper{Player: TeamCounterTerrorists[i], Impact: 0}
				fmt.Printf("Added Player %s to team2\n", TeamCounterTerrorists[i])
			}

			CurMatch.Rounds = make([]match.RoundImpact, 0)
			fmt.Print("testest")
		}
	})

	p.RegisterEventHandler(func(e events.RoundStart) {
		if !p.GameState().IsMatchStarted() {
			return
		}
		fmt.Printf("\n\nRound starting!\n")
	})

	// Register handler on kill events
	p.RegisterEventHandler(func(e events.Kill) {
		if !p.GameState().IsMatchStarted() {
			return
		}
		var hs string
		if e.IsHeadshot {
			hs = " (HS)"
		}
		var wallBang string
		if e.PenetratedObjects > 0 {
			wallBang = " (WB)"
		}
		kill := match.KillImpact{
			Kill:            &e,
			NumEnemiesAlive: SumAlive(e.Victim.TeamState.Members()),
			NumAlliesAlive:  SumAlive(e.Killer.TeamState.Members()),
		}
		kill.CalculateImpact()
		fmt.Printf("%s <%v%s%s> %s\n", e.Killer, e.Weapon, hs, wallBang, e.Victim)
		fmt.Printf("\t\tNumber of opponents alive: %d\n", kill.NumEnemiesAlive)
		fmt.Printf("\t\tNumber of own Team alive: %d\n", kill.NumAlliesAlive)
		fmt.Printf("\t\t -> Impact of Kill: %.4f\n", kill.KillImpact())
		CurRound.AddImpact(e.Killer, kill.KillImpact())
		CurRound.Kills = append(CurRound.Kills, kill)
	})

	p.RegisterEventHandler(func(e events.RoundEnd) {
		gs := p.GameState()

		CurRound.NormalizeImpact()
		switch e.Winner {
		case common.TeamTerrorists:
			// Winner's score + 1 because it hasn't actually been updated yet
			fmt.Printf("Round finished: winnerSide=T  ; score=%d:%d\n\n", gs.TeamTerrorists().Score()+1, gs.TeamCounterTerrorists().Score())
			if CurRound.Team1[0].Player.Team == 2 {
				fmt.Printf("\t-> Impacts: %s : %.4f, %s : %.4f, %s : %.4f, %s : %.4f, %s : %.4f", CurRound.Team1[0].Player.Name, CurRound.Team1[0].Impact, CurRound.Team1[1].Player.Name, CurRound.Team1[1].Impact, CurRound.Team1[2].Player.Name, CurRound.Team1[2].Impact, CurRound.Team1[3].Player.Name, CurRound.Team1[3].Impact, CurRound.Team1[4].Player.Name, CurRound.Team1[4].Impact)
				CurRound.ResetImpact(2)
			} else if CurRound.Team2[0].Player.Team == 2 {
				fmt.Printf("\t-> Impacts: %s : %.4f, %s : %.4f, %s : %.4f, %s : %.4f, %s : %.4f", CurRound.Team2[0].Player.Name, CurRound.Team2[0].Impact, CurRound.Team2[1].Player.Name, CurRound.Team2[1].Impact, CurRound.Team2[2].Player.Name, CurRound.Team2[2].Impact, CurRound.Team2[3].Player.Name, CurRound.Team2[3].Impact, CurRound.Team2[4].Player.Name, CurRound.Team2[4].Impact)
				CurRound.ResetImpact(1)
			}
		case common.TeamCounterTerrorists:
			fmt.Printf("Round finished: winnerSide=CT ; score=%d:%d\n\n", gs.TeamCounterTerrorists().Score()+1, gs.TeamTerrorists().Score())
			if CurRound.Team1[0].Player.Team == 3 {
				fmt.Printf("\t-> Impacts: %s : %.4f, %s : %.4f, %s : %.4f, %s : %.4f, %s : %.4f", CurRound.Team1[0].Player.Name, CurRound.Team1[0].Impact, CurRound.Team1[1].Player.Name, CurRound.Team1[1].Impact, CurRound.Team1[2].Player.Name, CurRound.Team1[2].Impact, CurRound.Team1[3].Player.Name, CurRound.Team1[3].Impact, CurRound.Team1[4].Player.Name, CurRound.Team1[4].Impact)
				CurRound.ResetImpact(2)
			} else if CurRound.Team2[0].Player.Team == 3 {
				fmt.Printf("\t-> Impacts: %s : %.4f, %s : %.4f, %s : %.4f, %s : %.4f, %s : %.4f", CurRound.Team2[0].Player.Name, CurRound.Team2[0].Impact, CurRound.Team2[1].Player.Name, CurRound.Team2[1].Impact, CurRound.Team2[2].Player.Name, CurRound.Team2[2].Impact, CurRound.Team2[3].Player.Name, CurRound.Team2[3].Impact, CurRound.Team2[4].Player.Name, CurRound.Team2[4].Impact)
				CurRound.ResetImpact(1)
			}
		default:
			// Probably match medic or something similar
			fmt.Printf("Round finished: No winner (tie)\n\n")
		}

		CurMatch.Rounds = append(CurMatch.Rounds, CurRound)
		CurRound.ResetImpact(-1)
		CurRound.ResetKills()
	})

	// Parse to end
	err = p.ParseToEnd()
	if err != nil {
		log.Panic("failed to parse demo: ", err)
	}

	CurMatch.CalculateImpact()
	CurMatch.ValidateNormalizedImpact()

	fmt.Println("Finished parsming demo")
}

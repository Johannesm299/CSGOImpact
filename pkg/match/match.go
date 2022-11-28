package match

import "fmt"

type MatchImpact struct {
	Rounds      []RoundImpact
	MatchImpact [10]PlayerWrapper
}

func (m *MatchImpact) ValidateNormalizedImpact() {
	sumImpact := 0.0
	for i := range m.MatchImpact {
		sumImpact += m.MatchImpact[i].Impact
	}
	fmt.Printf("Match has sum of normalized Impacts of %.6f\n", sumImpact)
}

func (m *MatchImpact) CalculateImpact() {
	for r := range m.Rounds {
		for p := range m.Rounds[r].Team1 {
			if m.MatchImpact[p].Player == nil {
				m.MatchImpact[p] = PlayerWrapper{
					Player: m.Rounds[0].Team1[p].Player,
					Impact: 0,
				}
			}
			if m.MatchImpact[p+5].Player == nil {
				m.MatchImpact[p+5] = PlayerWrapper{
					Player: m.Rounds[0].Team2[p].Player,
					Impact: 0,
				}
			}
			//only works as long as roundImpact.team is filled in similar for loop
			m.MatchImpact[p].Impact += m.Rounds[r].Team1[p].Impact
			m.MatchImpact[p+5].Impact += m.Rounds[r].Team2[p].Impact

			if r == len(m.Rounds)-1 {
				m.MatchImpact[p].Impact = m.MatchImpact[p].Impact / float64(len(m.Rounds))
				m.MatchImpact[p+5].Impact = m.MatchImpact[p+5].Impact / float64(len(m.Rounds))
			}
		}
		fmt.Printf("Checking Impact of Round %d", r)
		m.Rounds[r].ValidateNormalizedImpact()
	}
}

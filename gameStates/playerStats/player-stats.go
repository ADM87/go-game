package playerStats

import (
	"fmt"
	"go-game/data"

	tea "github.com/charmbracelet/bubbletea"
)

type State struct {
	gameModel *data.GameModel
	hpBar     ProgressBar
	xpBar     ProgressBar
}

func NewState(mdl *data.GameModel) State {
	return State{
		gameModel: mdl,
		hpBar:     NewProgressBar(40, "#990202"),
		xpBar:     NewProgressBar(40, "#0f9902"),
	}
}

func (s *State) Init() {
}

func (s *State) OnKeyPressed(key string) tea.Cmd {
	return nil
}

func (s *State) Render() string {
	hp := s.gameModel.PlayerModel.HP
	lvl := s.gameModel.PlayerModel.Lvl

	output := "Player Stats\n"
	output += "---------------------------------------------\n"
	output += fmt.Sprintf("HP: %d \\ %d\n", hp.Current, hp.Max)
	output += s.hpBar.Render(hp.Current, hp.Max) + "\n\n"
	output += fmt.Sprintf("Level: %s\n", lvl.String())
	output += s.xpBar.Render(lvl.GetCurrentXP(), lvl.GetNextXP()) + "\n"
	output += fmt.Sprintf("XP: %d \\ %d\n", lvl.GetCurrentXP(), lvl.GetNextXP())

	return output
}

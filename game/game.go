package game

import (
	"go-game/data"
	"go-game/gameStates/gameplay"
	"go-game/gameStates/mapView"
	"go-game/stateMachine"

	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type Game struct {
	gameModel    data.GameModel
	stateMachine stateMachine.Model
}

func NewGame() Game {
	mdl := data.NewGameModel()
	gps := gameplay.NewState(&mdl)
	mvs := mapView.NewState()
	return Game{
		stateMachine: stateMachine.NewStateMachine(
			[]stateMachine.State{
				&gps,
				&mvs,
			},
			0,
		),
	}
}

func (g Game) Init() tea.Cmd {
	return g.stateMachine.Init()
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "esc":
			return g, tea.Quit

		case "1", "2":
			s, e := strconv.Atoi(key)
			if e != nil {
				panic(e)
			}
			if g.stateMachine.HasState(s - 1) {
				g.stateMachine.SetState(s - 1)
			}

		default:
			return g, g.stateMachine.OnKeyPressed(key)
		}
	}
	return g, nil
}

func (g Game) View() string {
	return g.stateMachine.Render()
}

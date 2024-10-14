package data

import "fmt"

// XPModel represents the player's experience points
type XPModel struct {
	Current int
	Next    int
}

// LvlModel represents the player's level
type LvlModel struct {
	Current int
	Max     int
	XP      []XPModel
}

func (l *LvlModel) GetCurrentXP() int {
	return l.XP[l.Current].Current
}

func (l *LvlModel) GetNextXP() int {
	return l.XP[l.Current].Next
}

func (l *LvlModel) IsMaxLevel() bool {
	return l.Current == l.Max
}

func (l *LvlModel) String() string {
	if l.IsMaxLevel() {
		return "Max"
	}
	return fmt.Sprintf("%d", l.Current)
}

// HPModel represents the player's health points
type HPModel struct {
	Current int
	Max     int
}

// PlayerModel represents the player's stats
type PlayerModel struct {
	HP  HPModel
	Lvl LvlModel
}

func NewPlayerModel() PlayerModel {
	return PlayerModel{
		HP: HPModel{
			Current: 100,
			Max:     100,
		},
		Lvl: LvlModel{
			Current: 1,
			Max:     10,
			XP: []XPModel{
				{Current: 0, Next: 100},
				{Current: 0, Next: 200},
				{Current: 0, Next: 400},
				{Current: 0, Next: 800},
				{Current: 0, Next: 1600},
				{Current: 0, Next: 3200},
				{Current: 0, Next: 6400},
				{Current: 0, Next: 12800},
				{Current: 0, Next: 25600},
				{Current: 0, Next: 51200},
			},
		},
	}
}

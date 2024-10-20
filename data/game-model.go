package data

const (
	UnknownId = iota - 1
	EmptyId
	WallId
	PlayerId
)

type GameModel struct {
	WorldWidth  int
	WorldHeight int
	ViewWidth   int
	ViewHeight  int
	GameTokens  map[int]string
	PlayerModel PlayerModel
}

func NewGameModel() GameModel {
	return GameModel{
		WorldWidth:  100,
		WorldHeight: 100,
		ViewWidth:   75,
		ViewHeight:  20,
		GameTokens: map[int]string{
			UnknownId: "?",
			EmptyId:   " ",
			WallId:    "█",
			PlayerId:  "O",
		},
		PlayerModel: NewPlayerModel(),
	}
}

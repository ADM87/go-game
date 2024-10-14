package data

const (
	UnknownId int = -1
	EmptyId   int = 0
	WallId    int = 1
	PlayerId  int = 2
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
		ViewWidth:   50,
		ViewHeight:  15,
		GameTokens: map[int]string{
			UnknownId: "?",
			EmptyId:   " ",
			WallId:    "█",
			PlayerId:  "☺",
		},
		PlayerModel: NewPlayerModel(),
	}
}

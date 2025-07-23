package player

type Player struct {
	ID    string
	Name  string
	Score int
}

func NewPlayer(id, name string) *Player {
	return &Player{
		ID:    id,
		Name:  name,
		Score: 0,
	}
}

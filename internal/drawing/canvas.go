package game

type Canvas struct {
	Strokes []Stroke
}

type Stroke struct {
	Points   []Point
	Color    string
	Width    int
	PlayerID string
}

type Point struct {
	X, Y float64
}

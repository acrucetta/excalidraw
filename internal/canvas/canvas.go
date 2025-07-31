package canvas

type Canvas struct {
	Strokes []StrokeSegment
}

// Point struct must have exported fields (capitalized) and correct JSON tags.
type Point struct {
	X float64 `json:"X"`
	Y float64 `json:"Y"`
}

// StrokeSegment must also have exported fields and matching JSON tags.
type StrokeSegment struct {
	P0       Point  `json:"P0"`
	P1       Point  `json:"P1"`
	Color    string `json:"Color"`
	Width    int    `json:"Width"`
	PlayerID string `json:"PlayerID"`
}

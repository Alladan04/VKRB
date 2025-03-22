package calc_path

type CalculatePathRequest struct {
	Start      Point `json:"start"`
	End        Point `json:"end"`
	LabirintID int64 `json:"labirintID"`
}

type Point struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type Transition struct {
	From Point `json:"start"`
	To   Point `json:"end"`
}

type CalculatePathResponse struct {
	Path []Transition `json:"path"`
	Time int64        `json:"time"` //время вычисления маршрута в секундах
}

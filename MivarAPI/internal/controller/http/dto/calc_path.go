package dto

type CalculatePathRequest struct {
	Start      Point   `json:"start"`
	End        []Point `json:"end"`
	LabirintID int64   `json:"labirintID"`
}

type CalculatePathResponse struct {
	Path []Transition `json:"path"`
	Time int64        `json:"time"` //время вычисления маршрута в секундах
}

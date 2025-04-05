package dto

import "mivar_robot_api/internal/entity"

type Point struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type Transition struct {
	From Point `json:"start"`
	To   Point `json:"end"`
}

type Transitions []Transition

func (p Point) ToEntity() entity.Point {
	return entity.Point{X: p.X, Y: p.Y}
}

func (t Transition) ToEntity() entity.Transition {
	return entity.Transition{From: t.From.ToEntity(), To: t.To.ToEntity()}
}

func (t Transitions) ToEntity() []entity.Transition {
	var result []entity.Transition
	for _, transition := range t {
		result = append(result, transition.ToEntity())
	}
	return result
}

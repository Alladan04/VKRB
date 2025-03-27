package entity

type Point struct {
	X int64
	Y int64
}

type Transition struct {
	From Point
	To   Point
}

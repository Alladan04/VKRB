package dto

type UpdateMapIn struct {
	Points     []Point `json:"points"`
	LabirintID int64   `json:"labirint_id" binding:"required" validate:"required"`
}

type UpdateMapOut struct {
	Labirint [][]uint8 `json:"labirint"`
}

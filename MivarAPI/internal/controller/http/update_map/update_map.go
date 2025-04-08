package update_map

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"mivar_robot_api/internal/controller/http/dto"
	"mivar_robot_api/internal/entity"
	"mivar_robot_api/utils"
)

type UpdateMapHandler struct {
	log *logrus.Logger
	uc  Usecase
}

func New(log *logrus.Logger, uc Usecase) *UpdateMapHandler {
	return &UpdateMapHandler{
		log: log,
		uc:  uc,
	}
}

func (h *UpdateMapHandler) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var bodyDTO dto.UpdateMapIn
	err = json.Unmarshal(body, &bodyDTO)
	if err != nil {
		h.log.Errorf("Error parsing body: %v", err)
		http.Error(w, "can't parse body", http.StatusBadRequest)
		return
	}

	labirint, err := h.uc.UpdateMap(
		r.Context(),
		h.convertDTOToEntity(bodyDTO),
		strconv.FormatInt(bodyDTO.LabirintID, 10),
	)
	if err != nil {
		h.log.Errorf("Error updating labirint: %v", err)
		http.Error(w, "can't update labirint", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(dto.MapOut{Labirint: utils.Uint8ToInt(labirint)})
	if err != nil {
		h.log.Errorf("Error encoding response: %v", err)
		return
	}

	err = json.NewEncoder(w).Encode(string(jsonData))
	if err != nil {
		h.log.Errorf("Error encoding response: %v", err)
		return
	}
}

func (h *UpdateMapHandler) convertDTOToEntity(dto dto.UpdateMapIn) []entity.Point {
	points := make([]entity.Point, 0, len(dto.Points))
	for _, p := range dto.Points {
		points = append(points, entity.Point{
			X: p.X,
			Y: p.Y,
		})
	}

	return points
}

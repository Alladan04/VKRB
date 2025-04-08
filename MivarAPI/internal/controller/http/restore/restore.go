package restore

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"mivar_robot_api/internal/controller/http/dto"
	"mivar_robot_api/internal/entity"
	"mivar_robot_api/utils"
)

type RestoreHandler struct {
	log *logrus.Logger
	uc  Usecase
}

func New(log *logrus.Logger, uc Usecase) *RestoreHandler {
	return &RestoreHandler{
		log: log,
		uc:  uc,
	}
}

func (h *RestoreHandler) Handle(w http.ResponseWriter, r *http.Request) {
	modelID := r.URL.Query().Get("labirint_id")

	labirint, err := h.uc.RestoreModel(
		r.Context(),
		modelID,
	)
	if err != nil {
		h.log.Errorf("uc.RestoreModel: %v", err)
		http.Error(w, "can't restore", http.StatusInternalServerError)
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

func (h *RestoreHandler) convertDTOToEntity(dto dto.UpdateMapIn) []entity.Point {
	points := make([]entity.Point, 0, len(dto.Points))
	for _, p := range dto.Points {
		points = append(points, entity.Point{
			X: p.X,
			Y: p.Y,
		})
	}

	return points
}

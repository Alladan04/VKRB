package calc_path

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"mivar_robot_api/internal/controller/http/dto"
	"mivar_robot_api/internal/entity"
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

	w.Header().Set("Content-Type", "application/json")

	//err = json.NewEncoder(w).Encode(resp)
	//if err != nil {
	//	h.log.Errorf("Error encoding response: %v", err)
	//	return
	//}
}

func (h *UpdateMapHandler) convertDTOToEntity(dto dto.CalculatePathRequest) (
	start entity.Point,
	end []entity.Point,
	modelID string) {
	endpoints := make([]entity.Point, 0, len(dto.End))
	for _, p := range dto.End {
		endpoints = append(endpoints, entity.Point{
			X: p.X,
			Y: p.Y,
		})
	}

	return entity.Point{
		X: dto.Start.X,
		Y: dto.Start.Y,
	}, endpoints, strconv.Itoa(int(dto.LabirintID))
}

func (h *UpdateMapHandler) convertEntityToDTO(path []entity.Transition, timing int64) (dto.CalculatePathResponse, error) {
	dtoPath := make([]dto.Transition, 0, len(path))
	for _, p := range path {
		dtoPath = append(dtoPath, dto.Transition{
			From: dto.Point{
				X: p.From.X,
				Y: p.From.Y,
			},
			To: dto.Point{
				X: p.To.X,
				Y: p.To.Y,
			},
		})
	}

	return dto.CalculatePathResponse{
		Path: dtoPath,
		Time: timing,
	}, nil
}

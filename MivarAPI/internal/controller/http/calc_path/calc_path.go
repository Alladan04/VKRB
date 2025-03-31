package calc_path

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"mivar_robot_api/internal/entity"
)

type CalcPathHandler struct {
	log *logrus.Logger
	uc  Usecase
}

func NewCalcPathHandler(log *logrus.Logger, uc Usecase) *CalcPathHandler {
	return &CalcPathHandler{
		log: log,
		uc:  uc,
	}
}

func (h *CalcPathHandler) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var bodyDTO CalculatePathRequest
	err = json.Unmarshal(body, &bodyDTO)
	if err != nil {
		h.log.Errorf("Error parsing body: %v", err)
		http.Error(w, "can't parse body", http.StatusBadRequest)
		return
	}

	start, end, modelID := h.convertDTOToEntity(bodyDTO)
	path, timing, err := h.uc.CalculatePath(r.Context(), start, end, modelID)
	if err != nil {
		h.log.Errorf("Error calculating path: %v", err)
		http.Error(w, "can't calculate path", http.StatusInternalServerError)
		return
	}

	resp, err := h.convertEntityToDTO(path, timing)
	if err != nil {
		h.log.Errorf("Error converting entity to dto: %v", err)
		http.Error(w, "can't convert entity to dto", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		h.log.Errorf("Error encoding response: %v", err)
		return
	}
}

func (h *CalcPathHandler) convertDTOToEntity(dto CalculatePathRequest) (
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

func (h *CalcPathHandler) convertEntityToDTO(path []entity.Transition, timing int64) (CalculatePathResponse, error) {
	dtoPath := make([]Transition, 0, len(path))
	for _, p := range path {
		dtoPath = append(dtoPath, Transition{
			From: Point{
				X: p.From.X,
				Y: p.From.Y,
			},
			To: Point{
				X: p.To.X,
				Y: p.To.Y,
			},
		})
	}

	return CalculatePathResponse{
		Path: dtoPath,
		Time: timing,
	}, nil
}

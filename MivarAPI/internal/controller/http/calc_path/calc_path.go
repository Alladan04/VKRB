package calc_path

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type CalcPathHandler struct {
	log *logrus.Logger
}

func NewCalcPathHandler(log *logrus.Logger) *CalcPathHandler {
	return &CalcPathHandler{
		log: log,
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

}

package calc_path

import (
	"encoding/json"
	"net/http"
)

type CalcPathHandler struct{}

func NewCalcPathHandler() *CalcPathHandler {
	return &CalcPathHandler{}
}

func (CalcPathHandler) Handle(w http.ResponseWriter, r *http.Request) {
	err := json.Unmarshal(r.Body, body)
}

package generator

import (
	"fmt"

	"github.com/ECUST-XX/xml"
)

func (g *Generator) UnmarshalModel(xmlData []byte) (Model, error) {
	var model Model
	err := xml.Unmarshal(xmlData, &model)
	if err != nil {
		return Model{}, fmt.Errorf("xml.Unmarshal:%v", err)
	}

	return model, nil
}

func (g *Generator) MarshalModel(model Model) ([]byte, error) {

	data, err := xml.Marshal(model)
	if err != nil {
		return nil, fmt.Errorf("xml.Marshal:%v", err)
	}

	return data, nil
}

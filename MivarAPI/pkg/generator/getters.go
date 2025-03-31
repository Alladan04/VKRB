package generator

import (
	"fmt"
	"slices"
	"strings"
)

func (g *Generator) GetParameterIDsByCoordinates(model Model, coordinates []Coordinate) ([]string, error) {
	modelCoordinates := make(map[Coordinate]string, len(model.Class.Parameters.Parameters))
	for _, p := range model.Class.Parameters.Parameters {
		strCoordinates := strings.Split(p.ShortName, ",")
		if len(strCoordinates) != 2 {
			return nil, fmt.Errorf("invalid parameter name: %s", p.ShortName)
		}

		modelCoordinates[Coordinate{
			X: strCoordinates[0],
			Y: strCoordinates[1],
		}] = p.ID
	}

	result := make([]string, 0, len(coordinates))

	for _, c := range coordinates {
		if _, ok := modelCoordinates[c]; ok {
			result = append(result, modelCoordinates[c])
		}
	}

	return result, nil
}

func (g *Generator) GetCoordinatesByParameterIDs(model Model, ids []string) (map[string]Coordinate, error) {
	modelCoordinates := make(map[string]Coordinate, len(model.Class.Parameters.Parameters))
	for _, p := range model.Class.Parameters.Parameters {
		if !slices.Contains(ids, p.ID) {
			continue
		}

		strCoordinates := strings.Split(p.ShortName, ",")
		if len(strCoordinates) != 2 {
			return nil, fmt.Errorf("invalid parameter name: %s", p.ShortName)
		}

		modelCoordinates[p.ID] = Coordinate{
			X: strCoordinates[0],
			Y: strCoordinates[1],
		}
	}

	return modelCoordinates, nil
}

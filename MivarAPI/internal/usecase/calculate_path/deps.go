package calculate_path

import "mivar_robot_api/pkg/generator"

type ModelManager interface {
	GetModel() generator.Model
	UnmarshalModel(xmlData []byte) (generator.Model, error)
	GetParameterIDsByCoordinates(model generator.Model, coordinates []generator.Coordinate) ([]string, error)
	GetCoordinatesByParameterIDs(model generator.Model, ids []string) (map[string]generator.Coordinate, error)
}

type ModelRepo interface {
	GetFromCache(key string) ([]byte, error)
}

package dto

const (
	OutputField_TIMING    = "timing"
	OutputField_ALGORITHM = "algorithm"

	ERR_MODEL_EXISTS = 5704
)

type AddModelRequest struct {
	ModelID       string `json:"modelID" binding:"required" example:"12345678901234567890123456789012"`
	ModelPoolSize string `json:"modelPoolSize" binding:"required" example:"1000000"`
	ModelXML      string `json:"modelXML" binding:"required" example:"<model>...</model>"`
}

type AddModelResponse struct {
	ModelID string `json:"modelID"`
	ErrorID int64  `json:"errorID"`
}

type GetModelResponse struct {
	ConstraintsCount int `json:"constraintsCount"`
	ParametersCount  int `json:"parametersCount"`
	PoolSize         int `json:"poolSize"`
	RelationsCount   int `json:"relationsCount"`
	RulesCount       int `json:"rulesCount"`
}

type CalculateRrequest struct {
	ModelID             string            `json:"modelID"`
	IncommingParameters []CalculateInItem `json:"incommingParameters"`
	OutputParameters    []string          `json:"outputParameters"`
	Service             Service           `json:"service"`
}

type Service struct {
	OutputFields []string `json:"outputFields"`
}

type CalculateInItem struct {
	Value int    `json:"value"`
	Id    string `json:"id"`
}

type CalculateResponse struct {
	Algorithm                  []ModelCalculateData        `json:"algorithm"`
	ExploredParameters         []ExploredParameter         `json:"exploredParameters"`
	RequiredExploredParameters []RequiredExploredParameter `json:"requiredExploredParameters"`
	Timing                     Timing                      `json:"timing"`
}

type ModelCalculateData struct {
	InputParameters  []InputParameter  `json:"inputParameters"`
	OutputParameters []OutputParameter `json:"outputParameters"`
	Rule             Rule              `json:"rule"`
	Script           string            `json:"script"`
}

type InputParameter struct {
	ModelParameterID    string `json:"modelParameterID"`
	RelationParameterID string `json:"relationParameterID"`
	Value               string `json:"value"`
}

type OutputParameter struct {
	ModelParameterID    string `json:"modelParameterID"`
	RelationParameterID string `json:"relationParameterID"`
	Value               string `json:"value"`
}

type Rule struct {
	Id string `json:"id"`
}

type ExploredParameter struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type RequiredExploredParameter struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

type Timing struct {
	RequestOutputGeneration int `json:"requestOutputGeneration"`
	RequestParsing          int `json:"requestParsing"`
	RequestProcessing       int `json:"requestProcessing"`
}

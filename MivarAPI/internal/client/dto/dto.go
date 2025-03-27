package dto

type AddModelRequest struct {
	ModelID       string `json:"modelID" binding:"required" example:"12345678901234567890123456789012"`
	ModelPoolSize string `json:"modelPoolSize" binding:"required" example:"1000000"`
	ModelXML      string `json:"modelXML" binding:"required" example:"<model>...</model>"`
}

type AddModelResponse struct {
	ModelID string `json:"modelID"`
}

type GetModelResponse struct {
	ConstraintsCount int `json:"constraintsCount"`
	ParametersCount  int `json:"parametersCount"`
	PoolSize         int `json:"poolSize"`
	RelationsCount   int `json:"relationsCount"`
	RulesCount       int `json:"rulesCount"`
}

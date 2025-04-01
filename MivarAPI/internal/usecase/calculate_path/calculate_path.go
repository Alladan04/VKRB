package calculate_path

import (
	"context"
	"errors"
	"strconv"

	"github.com/AlekSi/pointer"

	"mivar_robot_api/internal/client/dto"
	"mivar_robot_api/internal/entity"
	"mivar_robot_api/pkg/generator"
)

func (u *Usecase) CalculatePath(
	ctx context.Context,
	start entity.Point,
	end []entity.Point,
	modelID string,
) ([]entity.Transition, int64, error) {
	modelBytes, err := u.inMemRepo.GetFromCache(modelID)
	if err != nil {
		return nil, 0, err
	}

	model, err := u.modelGenerator.UnmarshalModel(modelBytes)
	if err != nil {
		return nil, 0, err
	}

	wimiReq, err := u.getMivarCalculateRequest(start, end, modelID, model)
	if err != nil {
		return nil, 0, err
	}

	clientResp, err := u.wimiCli.CalculatePath(ctx, wimiReq)
	if err != nil {
		return nil, 0, err
	}

	path, err := u.dtoAlgorithmToPath(pointer.Get(clientResp), model)
	if err != nil {
		return nil, 0, err
	}

	return path, int64(clientResp.Timing.RequestOutputGeneration +
		clientResp.Timing.RequestParsing +
		clientResp.Timing.RequestProcessing), nil
}

func (u *Usecase) getMivarCalculateRequest(
	start entity.Point,
	end []entity.Point,
	modelID string,
	model generator.Model,
) (dto.CalculateRrequest, error) {
	modelOutputCoordinates := u.pointsToCoordinates(end)
	if len(end) == 0 {
		points, err := u.manager.GetExitsByModelID(modelID)
		u.log.Info(points)
		if err != nil {
			u.log.Errorf("failed to get exits by model id %v", modelID)
			return dto.CalculateRrequest{}, err
		}

		modelOutputCoordinates = u.pointsToCoordinates(points)
	}

	u.log.Infof("model output coordinates: %v", modelOutputCoordinates)

	outputParams, err := u.modelGenerator.GetParameterIDsByCoordinates(model, modelOutputCoordinates)
	if err != nil {
		return dto.CalculateRrequest{}, err
	}

	if len(outputParams) == 0 {
		return dto.CalculateRrequest{}, errors.New("all end points are unavailable")
	}
	u.log.Infof("model output parameters: %v", outputParams)

	incomingParams, err := u.modelGenerator.GetParameterIDsByCoordinates(model, []generator.Coordinate{{
		X: strconv.Itoa(int(start.X)),
		Y: strconv.Itoa(int(start.Y)),
	}})
	if err != nil {
		return dto.CalculateRrequest{}, err
	}

	return dto.CalculateRrequest{
		ModelID: modelID,
		IncommingParameters: []dto.CalculateInItem{
			{
				Value: 0,
				Id:    incomingParams[0],
			},
		},
		OutputParameters: outputParams,
		Service: dto.Service{
			OutputFields: []string{
				dto.OutputField_ALGORITHM,
				dto.OutputField_TIMING,
			},
		},
	}, nil
}

func (u *Usecase) dtoAlgorithmToPath(dto dto.CalculateResponse, model generator.Model) ([]entity.Transition, error) {
	transitions := make([]entity.Transition, 0, len(dto.Algorithm))
	for _, transition := range dto.Algorithm {
		if len(transition.InputParameters) != 1 || len(transition.OutputParameters) != 1 {
			u.log.Warningf("unexpected number of input or output parameters for transition %v, %v",
				transition.InputParameters, transition.OutputParameters)

			continue
		}

		inputParam := transition.InputParameters[0]
		outputParam := transition.OutputParameters[0]

		coordinates, err := u.modelGenerator.GetCoordinatesByParameterIDs(model, []string{
			inputParam.ModelParameterID,
			outputParam.ModelParameterID,
		})
		if err != nil {
			return nil, err
		}

		points := u.coordinatesToPoints(coordinates)

		if len(points) != 2 {
			u.log.Warningf("unexpected number of points for transition %v, %v", inputParam, outputParam)
			continue
		}

		transitions = append(transitions, entity.Transition{
			From: points[0],
			To:   points[1],
		})
	}

	return transitions, nil
}

func (u *Usecase) coordinatesToPoints(coordinates map[string]generator.Coordinate) []entity.Point {
	points := make([]entity.Point, 0, len(coordinates))
	for _, coordinate := range coordinates {
		x, err := strconv.ParseInt(coordinate.X, 10, 64)
		if err != nil {
			u.log.Errorf("unexpected coordinate %v", coordinate)
			break
		}
		y, err := strconv.ParseInt(coordinate.Y, 10, 64)
		if err != nil {
			u.log.Errorf("unexpected coordinate %v", coordinate)
			break
		}

		points = append(points, entity.Point{
			X: x,
			Y: y,
		})
	}

	return points
}

func (u *Usecase) pointsToCoordinates(points []entity.Point) []generator.Coordinate {
	coordinates := make([]generator.Coordinate, 0, len(points))
	for _, point := range points {
		coordinates = append(coordinates, generator.Coordinate{
			X: strconv.FormatInt(point.X, 10),
			Y: strconv.FormatInt(point.Y, 10),
		})
	}
	return coordinates
}

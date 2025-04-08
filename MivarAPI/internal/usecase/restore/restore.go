package restore

import (
	"context"
	"fmt"

	"mivar_robot_api/internal/client/dto"
)

func (u *Usecase) RestoreModel(
	ctx context.Context,
	modelID string,
) ([][]uint8, error) {
	labirint, err := u.repo.GetLabirintByID(modelID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetLabirintByID: %v", err)
	}

	model, err := u.repo.GetModelByID(modelID)
	if err != nil {
		return nil, fmt.Errorf("repo.GetModelByID: %v", err)
	}

	err = u.wimiCli.DeleteModel(ctx, modelID)
	if err != nil {
		return nil, fmt.Errorf("wimiCli.DeleteModel: %v", err)
	}

	resp, err := u.wimiCli.AddModel(ctx, dto.AddModelRequest{
		ModelID:       modelID,
		ModelPoolSize: dto.DefaultModelPoolSize,
		ModelXML:      string(model),
	})
	if err != nil {
		return nil, fmt.Errorf("wimiCli.AddModel: %v, details: %v", err, resp)
	}

	u.inMemRepo.UpsertLabirintToCache(labirint, modelID)
	u.inMemRepo.UpsertModelToCache(model, modelID)

	return labirint, nil
}

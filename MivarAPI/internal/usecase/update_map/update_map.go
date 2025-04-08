package update

import (
	"context"
	"fmt"

	"mivar_robot_api/internal/client/dto"
	"mivar_robot_api/internal/entity"
)

func (u *Usecase) UpdateMap(
	ctx context.Context,
	points []entity.Point,
	modelID string,
) ([][]uint8, error) {
	//Изменить точки points на противоположное значение в лабиринте (рабочей версии, берем из кэша)
	labirint, err := u.inMemRepo.GetLabirintFromCache(modelID)
	if err != nil {
		return nil, fmt.Errorf("inMemRepo.GetLabirintFromCache: %v", err)
	}

	newLabirint := flipPoints(labirint, points)

	newModel, err := u.modelGenerator.GenerateModelFromLabyrinth(newLabirint, modelID)
	if err != nil {
		return nil, fmt.Errorf("modelGenerator.GenerateModelFromLabyrinth: %v", err)
	}

	u.inMemRepo.UpsertLabirintToCache(newLabirint, modelID)
	u.inMemRepo.UpsertModelToCache(newModel, modelID)

	err = u.wimiCli.DeleteModel(ctx, modelID)
	if err != nil {
		return nil, fmt.Errorf("wimiCli.DeleteModel: %v", err)
	}

	resp, err := u.wimiCli.AddModel(ctx, dto.AddModelRequest{
		ModelID:       modelID,
		ModelPoolSize: dto.DefaultModelPoolSize,
		ModelXML:      string(newModel),
	})
	if err != nil {
		return nil, fmt.Errorf("wimiCli.AddModel: %v, details: %v", err, resp)
	}

	return newLabirint, nil
}

func flipPoints(maze [][]uint8, points []entity.Point) [][]uint8 {
	newMaze := make([][]uint8, len(maze))
	for i := range maze {
		newMaze[i] = make([]uint8, len(maze[i]))
		copy(newMaze[i], maze[i])
	}

	// Инвертируем значения по указанным координатам
	for _, p := range points {
		newMaze[p.Y] = make([]uint8, len(maze[p.Y]))
		if p.Y >= 0 && p.Y < int64(len(newMaze)) && p.X >= 0 && p.X < int64(len(newMaze[p.Y])) {
			if maze[p.Y][p.X] == 0 {
				newMaze[p.Y][p.X] = 1
				continue
			}

			newMaze[p.Y][p.X] = 0
		}
	}

	return newMaze
}

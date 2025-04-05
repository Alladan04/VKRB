package model_manager

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"

	"mivar_robot_api/internal/client/dto"
	configer "mivar_robot_api/internal/config"
)

func (m *Manager) LoadModels(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(m.cfg.Model))

	for i, modelCfg := range m.cfg.Model {
		wg.Add(1)

		go func(i int, modelCfg configer.ModelCfg) {
			defer wg.Done()

			modelID := modelCfg.ModelID
			if modelID == "" {
				modelID = strconv.Itoa(i)
			}

			if modelCfg.ModelXmlPath == "" {
				errCh <- fmt.Errorf("modelXmlPath is required for model %d", i)
				return
			}

			matrix, err := readMatrixFromFile(modelCfg.FilePath)
			if err != nil {
				errCh <- fmt.Errorf("failed to read matrix for %s: %w", modelCfg.FilePath, err)
				return
			}

			var model []byte

			// Check if model XML file exists
			if _, statErr := os.Stat(modelCfg.ModelXmlPath); os.IsNotExist(statErr) {
				// File doesn't exist — generate and write
				model, err = m.modelManager.GenerateModelFromLabyrinth(matrix, modelID)
				if err != nil {
					errCh <- fmt.Errorf("failed to generate model %s: %w", modelID, err)
					return
				}

				if writeErr := os.WriteFile(modelCfg.ModelXmlPath, model, 0644); writeErr != nil {
					errCh <- fmt.Errorf("failed to write generated model XML for %s: %w", modelID, writeErr)
					return
				}
			} else {
				// Load existing model
				model, err = os.ReadFile(modelCfg.ModelXmlPath)
				if err != nil {
					errCh <- fmt.Errorf("failed to read existing model XML for %s: %w", modelID, err)
					return
				}
			}

			// Transaction-like parallel upserts and WiMi push
			g, grCtx := errgroup.WithContext(ctx)

			g.Go(func() error {
				m.inMemRepo.UpsertModelToCache(model, modelID)
				return nil
			})

			g.Go(func() error {
				res, grErr := m.wimiCli.AddModel(grCtx, dto.AddModelRequest{
					ModelID:       modelID,
					ModelPoolSize: dto.DefaultModelPoolSize,
					ModelXML:      string(model),
				})

				m.log.Info(res)

				if res.ErrorID == dto.ERR_MODEL_EXISTS {
					m.log.Info(fmt.Sprintf("Model %s already exists", modelID))
					return nil
				}

				return grErr
			})

			if err := g.Wait(); err != nil {
				errCh <- fmt.Errorf("failed to process model %s: %w", modelID, err)
				return
			}

			log.Printf("Successfully loaded model %s", modelID)
		}(i, modelCfg)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

func readMatrixFromFile(filename string) ([][]uint8, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	// Парсинг матрицы (примерная реализация)
	var matrix [][]uint8
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var row []uint8
		for _, char := range strings.Split(line, " ") {
			num, err := strconv.Atoi(char)
			if err != nil {
				return nil, fmt.Errorf("invalid matrix format: %w", err)
			}
			row = append(row, uint8(num))
		}
		matrix = append(matrix, row)
	}

	return matrix, nil
}

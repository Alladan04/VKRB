package model_manager

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"

	"mivar_robot_api/internal/client/dto"
)

func (m *Manager) LoadModels(ctx context.Context) error {
	var wg sync.WaitGroup //можно тоже на errgroup заменить по идее
	errCh := make(chan error, len(m.cfg.Model))

	for i, modelCfg := range m.cfg.Model {
		wg.Add(1)

		go func() {
			defer wg.Done()

			//ctx, cancel := context.WithTimeout(context.Background(), mc.InitTimeout)
			//defer cancel()

			matrix, err := readMatrixFromFile(modelCfg.FilePath)
			if err != nil {
				errCh <- fmt.Errorf("failed to read matrix for %s: %w", modelCfg.FilePath, err)
				return
			}

			model, err := m.modelManager.GenerateModelFromLabyrinth(matrix, strconv.Itoa(i))
			if err != nil {
				errCh <- fmt.Errorf("failed to generate model %s: %w", modelCfg.FilePath, err)
				return
			}

			err = os.WriteFile(fmt.Sprintf("%s_%s.xml", uuid.NewString(), strconv.Itoa(i)), model, 777)
			if err != nil {
				panic(err)
			}

			// Параллельная загрузка в кэш и сервис
			g, grCtx := errgroup.WithContext(ctx)
			g.Go(func() error {
				m.inMemRepo.AddToCache(model, strconv.Itoa(i))
				return nil
			})

			g.Go(func() error {
				m.inMemRepo.AddLabirintToCache(matrix, strconv.Itoa(i))
				return nil
			})

			g.Go(func() error {
				res, grErr := m.wimiCli.AddModel(grCtx, dto.AddModelRequest{
					ModelID:       strconv.Itoa(i),
					ModelPoolSize: "100",
					ModelXML:      string(model),
				})

				m.log.Info(res)

				if res.ErrorID == dto.ERR_MODEL_EXISTS {
					m.log.Info(fmt.Sprintf("Model %s already exists", strconv.Itoa(i)))
					return nil
				}

				return grErr
			})

			if err := g.Wait(); err != nil {
				errCh <- fmt.Errorf("failed to process model %s: %w", modelCfg.FilePath, err)
				return
			}

			log.Printf("Successfully loaded model %s", modelCfg.FilePath)
		}()
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

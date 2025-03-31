package model_manager

import (
	"github.com/sirupsen/logrus"

	configer "mivar_robot_api/internal/config"
)

type Manager struct {
	log          *logrus.Logger
	inMemRepo    ModelRepo
	modelManager ModelGenerator
	wimiCli      MivarClient
	cfg          configer.Config
}

func New(log *logrus.Logger, inMemRepo ModelRepo, modelManager ModelGenerator, wimiCli MivarClient, cfg configer.Config) *Manager {
	return &Manager{
		log:          log,
		inMemRepo:    inMemRepo,
		modelManager: modelManager,
		wimiCli:      wimiCli,
		cfg:          cfg,
	}
}

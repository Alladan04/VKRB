package calculate_path

import (
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	log            *logrus.Logger
	inMemRepo      ModelRepo
	modelGenerator ModelGenerator
	manager        ModelManager
	wimiCli        Client
}

func New(log *logrus.Logger, inMemRepo ModelRepo, modelGenerator ModelGenerator, wimiCli Client, manager ModelManager) *Usecase {
	return &Usecase{
		log:            log,
		inMemRepo:      inMemRepo,
		modelGenerator: modelGenerator,
		wimiCli:        wimiCli,
		manager:        manager,
	}
}

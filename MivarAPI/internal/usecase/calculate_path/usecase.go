package calculate_path

import (
	"github.com/sirupsen/logrus"

	"mivar_robot_api/internal/client/mivar"
)

type Usecase struct {
	log          *logrus.Logger
	inMemRepo    ModelRepo
	modelManager ModelManager
	wimiCli      mivar.Client
}

func New(log *logrus.Logger, inMemRepo ModelRepo, modelManager ModelManager, wimiCli mivar.Client) *Usecase {
	return &Usecase{
		log:          log,
		inMemRepo:    inMemRepo,
		modelManager: modelManager,
		wimiCli:      wimiCli,
	}
}

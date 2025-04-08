package update

import (
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	log            *logrus.Logger
	inMemRepo      Repo
	modelGenerator ModelGenerator
	wimiCli        Client
}

func New(log *logrus.Logger, inMemRepo Repo, modelGenerator ModelGenerator, wimiCli Client) *Usecase {
	return &Usecase{
		log:            log,
		inMemRepo:      inMemRepo,
		modelGenerator: modelGenerator,
		wimiCli:        wimiCli,
	}
}

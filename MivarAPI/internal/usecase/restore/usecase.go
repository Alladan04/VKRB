package restore

import (
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	log       *logrus.Logger
	inMemRepo Repo
	repo      PersistentRepo
	wimiCli   Client
}

func New(log *logrus.Logger, inMemRepo Repo, repo PersistentRepo, wimiCli Client) *Usecase {
	return &Usecase{
		log:       log,
		inMemRepo: inMemRepo,
		wimiCli:   wimiCli,
		repo:      repo,
	}
}

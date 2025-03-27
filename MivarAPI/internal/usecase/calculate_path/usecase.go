package calculate_path

import (
	"github.com/sirupsen/logrus"

	"mivar_robot_api/internal/controller/http/calc_path"
	"mivar_robot_api/internal/entity"
	"mivar_robot_api/internal/repo/cache"
	"mivar_robot_api/pkg/generator"
)

type Usecase struct {
	log          *logrus.Logger
	inMemRepo    cache.Repo
	modelManager generator.Generator
}

func New(log *logrus.Logger, inMemRepo cache.Repo, modelManager generator.Generator) *Usecase {
	return &Usecase{
		log:          log,
		inMemRepo:    inMemRepo,
		modelManager: modelManager,
	}
}

func (u *Usecase) CalculatePath(start entity.Point, end entity.Point, modelID int64) (calc_path.CalculatePathResponse, error) {

}

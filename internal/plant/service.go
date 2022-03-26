package plant

import (
	"artificial-catalog/internal/entity"
	"context"
	"fmt"
	"io/ioutil"

	"go.uber.org/zap"
)

const limit = "3"
const page = "1"

type PlantService interface {
	GetPage(ctx context.Context) ([]entity.Plant, error)
	GetImage(ctx context.Context, iamgeNamge string) ([]byte, error)
}

type plantService struct {
	repo   PlantRepository
	logger *zap.Logger
}

func NewPlantService(repo PlantRepository, logger *zap.Logger) PlantService {
	plantService := new(plantService)
	plantService.repo = repo
	plantService.logger = logger
	return plantService
}

func (r plantService) GetPage(ctx context.Context) ([]entity.Plant, error) {
	plants, err := r.repo.GetPage(ctx, page, limit)
	if err != nil {
		return []entity.Plant{}, err
	}
	return plants, err
}

func (r plantService) GetImage(ctx context.Context, iamgeNamge string) ([]byte, error) {
	f, err := ioutil.ReadFile(fmt.Sprintf("images/%s", iamgeNamge))
	if err != nil {
		return nil, err
	}
	return f, err
}

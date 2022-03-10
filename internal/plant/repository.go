package plant

import (
	"artificial-catalog/internal/entity"
	"context"
	"database/sql"

	//"fmt"
	"strings"

	"go.uber.org/zap"
)

type PlantRepository interface {
	GetPage(ctx context.Context, page string, limit string) ([]entity.Plant, error)
}

type plantRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewPlantRepository(db *sql.DB, logger *zap.Logger) PlantRepository {
	plantRepository := new(plantRepository)
	plantRepository.db = db
	plantRepository.logger = logger
	return plantRepository
}

func (r *plantRepository) GetPage(ctx context.Context, page string, limit string) ([]entity.Plant, error) {
	rows, err := r.db.Query("select product.id, product.Name, product.Image_name, product.Height, product.Preparation, GROUP_CONCAT(tag.Name) from product join product_tag on product_tag.product_id = product.id join tag on tag.id = product_tag.tag_id GROUP BY id")
	if err != nil {
		return []entity.Plant{}, err
	}
	defer rows.Close()

	var result []entity.Plant
	var tags string

	for rows.Next() {
		var plants entity.Plant

		err := rows.Scan(
			&plants.Id,
			&plants.Name,
			&plants.Image_name,
			&plants.Height,
			&plants.Preparation,
			&tags,
		)
		if err != nil {
			return []entity.Plant{}, err
		}
		plants.Tags = strings.Split(tags, ",")

		result = append(result, plants)
	}

	return result, err
}

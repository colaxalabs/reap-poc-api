package resolvers

import (
	"context"
	"errors"
	models1 "github.com/3dw1nM0535/Byte/db/models"
	"github.com/3dw1nM0535/Byte/models"
)

func (r *mutationResolver) UpdateFarmHarvests(ctx context.Context, input *models.HarvestInput) (*models1.Season, error) {
	season := &models1.Season{}
	r.ORM.DB.Where("season_number = ? AND token = ?", input.SeasonNumber, input.Token).First(&season)
	if season.ID.String() == specialUUID && r.ORM.DB.NewRecord(season) {
		notFound := errors.New("no season data. start from preparations")
		return nil, notFound
	}
	r.ORM.DB.Model(&season).Updates(&models1.Season{
		HarvestYield: input.TotalSupply,
		HarvestPrice: input.Price,
		HarvestUnit:  input.SupplyUnit,
	})
	return season, nil
}

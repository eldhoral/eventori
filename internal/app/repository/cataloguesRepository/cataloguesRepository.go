package cataloguesRepository

import (
	datapaging "eldho/eventori/internal/app/commons/dataPagingHelper"
	"eldho/eventori/internal/app/model/cataloguesModel"
	"eldho/eventori/internal/app/model/schedulesModel"

	param "eldho/eventori/internal/pkg/params"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type cataloguesRepository struct {
	db *gorm.DB
}

func NewCataloguesRepository(db *gorm.DB) ICataloguesRepository {
	return &cataloguesRepository{db: db}
}

func (c cataloguesRepository) ListCatalogues(paging datapaging.Datapaging) (data *[]cataloguesModel.Catalogues, count int64, err error) {
	model := []cataloguesModel.Catalogues{}
	db := c.db.Preload("Schedules").Model(&model)

	db.Count(&count)
	if paging.Page != 0 {
		pg := datapaging.New(paging.Limit, paging.Page, []string{})
		db.Offset(pg.GetOffset()).Limit(paging.Limit)
	}
	db.Order("updated_at desc").Find(&model)

	return &model, count, db.Error
}

func (c cataloguesRepository) CreateCatalogues(parameter param.Params) (err error) {
	data := cataloguesModel.Catalogues{
		ModelName:   parameter.GetString("model_name"),
		ImageUrl:    parameter.GetString("image_url"),
		Description: parameter.GetString("description"),
	}
	err = c.db.Create(&data).Error
	return
}

func (c cataloguesRepository) GetCataloguesByModelId(modelId int64, paging datapaging.Datapaging) (data *[]cataloguesModel.Catalogues, count int64, err error) {
	db := c.db.Preload("Schedules").Model(&data)
	db.Where(cataloguesModel.Catalogues{Id: modelId})

	db.Count(&count)
	if paging.Page != 0 {
		pg := datapaging.New(paging.Limit, paging.Page, []string{})
		db.Offset(pg.GetOffset()).Limit(paging.Limit)
	}
	db.Order("updated_at desc").Find(&data)

	return
}

func (c cataloguesRepository) UpdateCataloguesByModelId(modelId int64, parameter param.Params) (err error) {
	data := cataloguesModel.Catalogues{
		ModelName:   parameter.GetString("model_name"),
		ImageUrl:    parameter.GetString("image_url"),
		Description: parameter.GetString("description"),
	}

	db := c.db.Model(&cataloguesModel.Catalogues{})

	db.Where(cataloguesModel.Catalogues{Id: modelId})

	err = db.Updates(&data).Error

	return
}

func (c cataloguesRepository) DeleteCataloguesByModelId(modelId int64) (err error) {
	dataCatalogues := cataloguesModel.Catalogues{}
	dataSchedules := schedulesModel.Schedules{}

	tx := c.db.Begin()
	defer func() {
		// if error, rollback all changes
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				log.Error().Err(errRollback.Error).Msg("err, Rollback transaction at DeleteCataloguesByModelId")
			}
			return
		}

		// commit the db transaction
		if err = tx.Commit().Error; err != nil {
			log.Error().Msg("err, Commit transaction at DeleteCataloguesByModelId")
		}

	}()

	// Delete Catalogues data as child
	tx.Model(&dataSchedules)
	err = tx.Delete(&dataSchedules, schedulesModel.Schedules{ModelId: modelId}).Error

	// Delete Catalogues data as parent
	tx.Model(&dataCatalogues)
	err = tx.Delete(&dataCatalogues, cataloguesModel.Catalogues{Id: modelId}).Error

	return
}

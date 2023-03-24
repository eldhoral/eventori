package modelService

import (
	datapaging "eldho/eventori/internal/app/commons/dataPagingHelper"
	"eldho/eventori/internal/app/model/cataloguesModel"
	"eldho/eventori/internal/app/model/schedulesModel"

	param "eldho/eventori/internal/pkg/params"
)

type IModelService interface {
	ListCatalogues(paging datapaging.Datapaging) (httpStatus int, result *[]cataloguesModel.Catalogues, count int64, err error)
	CreateCatalogues(parameter param.Params) (httpStatus int, err error)
	ListSchedulesByModelId(modelId int64, paging datapaging.Datapaging) (httpStatus int, result *[]schedulesModel.Schedules, count int64, err error)
	CreateSchedules(parameter param.Params) (httpStatus int, err error)
	ListCataloguesByModelId(modelId int64, paging datapaging.Datapaging) (httpStatus int, result *[]cataloguesModel.Catalogues, count int64, err error)
	UpdateCataloguesByModelId(modelId int64, parameter param.Params) (httpStatus int, err error)
	UpdateSchedulesById(id int64, parameter param.Params) (httpStatus int, err error)
	DeleteSchedulesById(id int64) (httpStatus int, err error)
	DeleteCataloguesById(id int64) (httpStatus int, err error)
}

package modelService

import (
	datapaging "eldho/eventori/internal/app/commons/dataPagingHelper"
	"eldho/eventori/internal/app/model/cataloguesModel"
	"eldho/eventori/internal/app/model/schedulesModel"
	"eldho/eventori/internal/app/repository/cataloguesRepository"
	"eldho/eventori/internal/app/repository/schedulesRepository"
	"errors"
	"net/http"

	param "eldho/eventori/internal/pkg/params"

	"github.com/rs/zerolog/log"
)

type modelService struct {
	cataloguesRepository cataloguesRepository.ICataloguesRepository
	schedulesRepository  schedulesRepository.ISchedulesRepository
}

func NewModelServiceImpl(
	cataloguesRepository cataloguesRepository.ICataloguesRepository,
	schedulesRepository schedulesRepository.ISchedulesRepository,
) IModelService {
	return &modelService{
		cataloguesRepository,
		schedulesRepository,
	}
}

func (m modelService) ListCatalogues(paging datapaging.Datapaging) (httpStatus int, result *[]cataloguesModel.Catalogues, count int64, err error) {
	result, count, err = m.cataloguesRepository.ListCatalogues(paging)
	if err != nil {
		log.Error().Msg("[ListCatalogues] Failed to list the data with message : " + err.Error())
		return http.StatusInternalServerError, nil, count, errors.New("[ListCatalogues] Failed to list the data with message : error")
	}
	if len((*result)) == 0 {
		log.Error().Msg("[ListCatalogues] Failed to list the data with message : not found")
		return http.StatusNotFound, nil, count, errors.New("[ListCatalogues] Failed to list the data with message : not found")
	}
	return
}

func (m modelService) CreateCatalogues(parameter param.Params) (httpStatus int, err error) {
	err = m.cataloguesRepository.CreateCatalogues(parameter)
	if err != nil {
		log.Error().Msg("[CreateCatalogues] Failed to create the data with message : " + err.Error())
		return http.StatusInternalServerError, errors.New("[CreateCatalogues] Failed to create the data with message : error")
	}
	return
}

func (m modelService) ListSchedulesByModelId(modelId int64, paging datapaging.Datapaging) (httpStatus int, result *[]schedulesModel.Schedules, count int64, err error) {
	result, count, err = m.schedulesRepository.GetSchedulesByModelId(modelId, paging)
	if err != nil {
		log.Error().Msg("[ListSchedulesByModelId] Failed to list the data with message : " + err.Error())
		return http.StatusInternalServerError, nil, count, errors.New("[ListSchedulesByModelId] Failed to list the data with message : error")
	}

	if len((*result)) == 0 {
		log.Error().Msg("[ListSchedulesByModelId] Failed to list the data with message : not found")
		return http.StatusNotFound, nil, count, errors.New("[ListSchedulesByModelId] Failed to list the data with message : not found")
	}
	return
}

func (m modelService) CreateSchedules(parameter param.Params) (httpStatus int, err error) {
	err = m.schedulesRepository.CreateSchedules(parameter)
	if err != nil {
		log.Error().Msg("[CreateSchedules] Failed to create the data with message : " + err.Error())
		return http.StatusInternalServerError, errors.New("[CreateSchedules] Failed to create the data with message : error")
	}
	return
}

func (m modelService) ListCataloguesByModelId(modelId int64, paging datapaging.Datapaging) (httpStatus int, result *[]cataloguesModel.Catalogues, count int64, err error) {
	result, count, err = m.cataloguesRepository.GetCataloguesByModelId(modelId, paging)
	if err != nil {
		log.Error().Msg("[ListCataloguesByModelId] Failed to list the data with message : " + err.Error())
		return http.StatusInternalServerError, nil, count, errors.New("[ListCataloguesByModelId] Failed to list the data with message : error")
	}

	if len((*result)) == 0 {
		log.Error().Msg("[ListCataloguesByModelId] Failed to list the data with message : not found")
		return http.StatusNotFound, nil, count, errors.New("[ListCataloguesByModelId] Failed to list the data with message : not found")
	}
	return
}

func (m modelService) UpdateCataloguesByModelId(modelId int64, parameter param.Params) (httpStatus int, err error) {
	err = m.cataloguesRepository.UpdateCataloguesByModelId(modelId, parameter)
	if err != nil {
		log.Error().Msg("[UpdateCataloguesByModelId] Failed to update the data with message : " + err.Error())
		return http.StatusInternalServerError, errors.New("[UpdateCataloguesByModelId] Failed to update the data with message : error")
	}
	return
}

func (m modelService) UpdateSchedulesById(id int64, parameter param.Params) (httpStatus int, err error) {
	err = m.schedulesRepository.UpdateSchedulesById(id, parameter)
	if err != nil {
		log.Error().Msg("[UpdateSchedulesByModelId] Failed to update the data with message : " + err.Error())
		return http.StatusInternalServerError, errors.New("[UpdateSchedulesByModelId] Failed to update the data with message : error")
	}
	return
}

func (m modelService) DeleteSchedulesById(id int64) (httpStatus int, err error) {
	err = m.schedulesRepository.DeleteSchedulesById(id)
	if err != nil {
		log.Error().Msg("[DeleteSchedulesById] Failed to delete the data with message : " + err.Error())
		return http.StatusInternalServerError, errors.New("[DeleteSchedulesById] Failed to delete the data with message : error")
	}
	return
}

func (m modelService) DeleteCataloguesById(id int64) (httpStatus int, err error) {
	err = m.cataloguesRepository.DeleteCataloguesByModelId(id)
	if err != nil {
		log.Error().Msg("[DeleteCataloguesById] Failed to delete the data with message : " + err.Error())
		return http.StatusInternalServerError, errors.New("[DeleteCataloguesById] Failed to delete the data with message : error")
	}
	return
}

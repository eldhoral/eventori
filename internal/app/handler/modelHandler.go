package handler

import (
	datapaging "eldho/eventori/internal/app/commons/dataPagingHelper"
	"eldho/eventori/internal/app/commons/uploadFile"
	"eldho/eventori/internal/app/model/cataloguesModel"
	"eldho/eventori/internal/app/model/schedulesModel"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"

	utils "eldho/eventori/internal/pkg"
	"eldho/eventori/internal/pkg/loggers"
	"eldho/eventori/internal/pkg/params"
	param "eldho/eventori/internal/pkg/params"
)

type ModelHandler struct {
	HandlerOption
}

func (handler ModelHandler) ListCatalogues(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	pageNumber := cast.ToInt(c.Query("page_number"))
	pageSize := cast.ToInt(c.Query("page_size"))

	paging := datapaging.Datapaging{Page: pageNumber, Limit: pageSize}

	httpStatus, result, count, err := handler.ModelService.ListCatalogues(paging)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, httpStatus, err.Error(), "Failed retrive data")
		return
	}

	Data := map[string]interface{}{
		"page_number":        pageNumber,
		"page_size":          pageSize,
		"total_record_count": count,
		"records":            result,
	}

	utils.BasicResponse(record, c.Writer, true, http.StatusOK, Data, "Success retrive data")
}

func (handler ModelHandler) CreateCatalogues(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	var request cataloguesModel.CataloguesRequest
	c.ShouldBind(&request)

	err := request.Validate()
	if err != nil {
		log.Error().Msg(err.Error())
		utils.BasicResponse(record, c.Writer, false, http.StatusBadRequest, err.Error(), "Failed validate params")
		return
	}

	// Upload file
	newFileName, err := uploadFile.UploadFIleInstance(request.Image, c)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "Failed retrieve image")
		return
	}

	params := param.NewParamsWrapper()
	params.Add("model_name", request.ModelName)
	params.Add("image_url", newFileName)
	params.Add("description", request.Description)

	httpStatus, err := handler.ModelService.CreateCatalogues(params)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.BasicResponse(record, c.Writer, false, httpStatus, err.Error(), "Failed create data")
		return
	}

	utils.BasicResponse(record, c.Writer, true, http.StatusOK, nil, "Success create data")
}

func (handler ModelHandler) ListSchedulesByModelId(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	pageNumber := cast.ToInt(c.Query("page_number"))
	pageSize := cast.ToInt(c.Query("page_size"))
	modelId := cast.ToInt64(c.Param("model_id"))

	paging := datapaging.Datapaging{Page: pageNumber, Limit: pageSize}

	httpStatus, result, count, err := handler.ModelService.ListSchedulesByModelId(modelId, paging)
	if err != nil {
		utils.BasicResponse(record, c.Writer, false, httpStatus, err.Error(), "Failed retrive data")
		return
	}

	Data := map[string]interface{}{
		"page_number":        pageNumber,
		"page_size":          pageSize,
		"total_record_count": count,
		"records":            result,
	}

	utils.BasicResponse(record, c.Writer, true, http.StatusOK, Data, "Success retrive data")
}

func (handler ModelHandler) CreateSchedules(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	var request schedulesModel.SchedulesRequest
	c.ShouldBind(&request)

	err := request.Validate()
	if err != nil {
		log.Error().Msg(err.Error())
		utils.BasicResponse(record, c.Writer, false, http.StatusBadRequest, err.Error(), "Failed validate params")
		return
	}

	params := params.NewParamsWrapper()
	params.Add("model_id", request.ModelId)
	params.Add("schedule_date", request.ScheduleDate)

	httpStatus, err := handler.ModelService.CreateSchedules(params)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.BasicResponse(record, c.Writer, false, httpStatus, err.Error(), "Failed create data")
		return
	}

	utils.BasicResponse(record, c.Writer, true, http.StatusOK, nil, "Success create data")
}

func (handler ModelHandler) ListCataloguesByModelId(c *gin.Context) {
	record := loggers.StartRecord(c.Request)
	pageNumber := cast.ToInt(c.Query("page_number"))
	pageSize := cast.ToInt(c.Query("page_size"))
	modelId := cast.ToInt64(c.Param("model_id"))

	paging := datapaging.Datapaging{Page: pageNumber, Limit: pageSize}

	httpStatus, result, count, err := handler.ModelService.ListCataloguesByModelId(modelId, paging)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.BasicResponse(record, c.Writer, false, httpStatus, err.Error(), "Failed retrive data")
		return
	}

	if len((*result)) == 0 {
		log.Error().Msg(err.Error())
		utils.BasicResponse(record, c.Writer, false, httpStatus, "No data was found", "Failed retrive data")
		return
	}

	Data := map[string]interface{}{
		"page_number":        pageNumber,
		"page_size":          pageSize,
		"total_record_count": count,
		"records":            result,
	}

	utils.BasicResponse(record, c.Writer, true, http.StatusOK, Data, "Success retrive data")
}

func (handler ModelHandler) UpdateCatalogues(c *gin.Context) {
	modelId := cast.ToInt64(c.Param("model_id"))

	record := loggers.StartRecord(c.Request)
	var (
		destinationFile string
		err             error
		request         cataloguesModel.CataloguesRequest
	)

	c.ShouldBind(&request)

	if request.Image != nil {
		// Upload file
		destinationFile, err = uploadFile.UploadFIleInstance(request.Image, c)
		if err != nil {
			utils.BasicResponse(record, c.Writer, false, http.StatusInternalServerError, err.Error(), "Failed retrieve image")
			return
		}
	}

	params := params.NewParamsWrapper()
	params.Add("model_name", request.ModelName)
	params.Add("image_url", destinationFile)
	params.Add("description", request.Description)

	httpStatus, err := handler.ModelService.UpdateCataloguesByModelId(modelId, params)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.BasicResponse(record, c.Writer, false, httpStatus, err.Error(), "Failed create data")
		return
	}

	utils.BasicResponse(record, c.Writer, true, http.StatusOK, nil, "Success create data")
}

func (handler ModelHandler) UpdateSchedules(c *gin.Context) {
	scheduleId := cast.ToInt64(c.Param("schedule_id"))

	record := loggers.StartRecord(c.Request)
	var request schedulesModel.SchedulesRequest
	c.ShouldBind(&request)

	params := params.NewParamsWrapper()
	params.Add("model_id", request.ModelId)
	params.Add("schedule_date", request.ScheduleDate)

	httpStatus, err := handler.ModelService.UpdateSchedulesById(scheduleId, params)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.BasicResponse(record, c.Writer, false, httpStatus, err.Error(), "Failed update data")
		return
	}

	utils.BasicResponse(record, c.Writer, true, http.StatusOK, nil, "Success update data")
}

func (handler ModelHandler) DeleteSchedules(c *gin.Context) {
	scheduleId := cast.ToInt64(c.Param("schedule_id"))

	record := loggers.StartRecord(c.Request)

	httpStatus, err := handler.ModelService.DeleteSchedulesById(scheduleId)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.BasicResponse(record, c.Writer, false, httpStatus, err.Error(), "Failed delete data")
		return
	}

	utils.BasicResponse(record, c.Writer, true, http.StatusOK, nil, "Success delete data")
}

func (handler ModelHandler) DeleteCatalogues(c *gin.Context) {
	modelId := cast.ToInt64(c.Param("model_id"))

	record := loggers.StartRecord(c.Request)

	httpStatus, err := handler.ModelService.DeleteCataloguesById(modelId)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.BasicResponse(record, c.Writer, false, httpStatus, err.Error(), "Failed delete data")
		return
	}

	utils.BasicResponse(record, c.Writer, true, http.StatusOK, nil, "Success delete data")
}

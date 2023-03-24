package schedulesRepository

import (
	datapaging "eldho/eventori/internal/app/commons/dataPagingHelper"
	"eldho/eventori/internal/app/model/schedulesModel"

	param "eldho/eventori/internal/pkg/params"
)

type ISchedulesRepository interface {
	GetSchedulesByModelId(modelId int64, paging datapaging.Datapaging) (data *[]schedulesModel.Schedules, count int64, err error)
	CreateSchedules(parameter param.Params) (err error)
	UpdateSchedulesById(id int64, parameter param.Params) (err error)
	DeleteSchedulesById(id int64) (err error)
}

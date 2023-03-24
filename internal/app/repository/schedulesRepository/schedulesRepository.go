package schedulesRepository

import (
	datapaging "eldho/eventori/internal/app/commons/dataPagingHelper"
	"eldho/eventori/internal/app/model/schedulesModel"

	param "eldho/eventori/internal/pkg/params"

	"gorm.io/gorm"
)

type schedulesRepository struct {
	db *gorm.DB
}

func NewSchedulesRepository(db *gorm.DB) ISchedulesRepository {
	return &schedulesRepository{db: db}
}

func (s schedulesRepository) GetSchedulesByModelId(modelId int64, paging datapaging.Datapaging) (data *[]schedulesModel.Schedules, count int64, err error) {
	db := s.db.Model(&data).Where(schedulesModel.Schedules{ModelId: modelId})

	db.Count(&count)
	if paging.Page != 0 {
		pg := datapaging.New(paging.Limit, paging.Page, []string{})
		db.Offset(pg.GetOffset()).Limit(paging.Limit)
	}
	db.Order("updated_at desc").Find(&data)

	return
}

func (s schedulesRepository) CreateSchedules(parameter param.Params) (err error) {
	scheduleDate, err := parameter.GetLocalTime("schedule_date")
	if err != nil {
		return err
	}

	data := schedulesModel.Schedules{
		ModelId:      parameter.GetInt64("model_id"),
		ScheduleDate: scheduleDate,
	}
	err = s.db.Create(&data).Error

	return
}

func (s schedulesRepository) UpdateSchedulesById(id int64, parameter param.Params) (err error) {
	scheduleDate, err := parameter.GetLocalTime("schedule_date")
	if err != nil {
		return err
	}

	data := schedulesModel.Schedules{
		Id:           parameter.GetInt64("id"),
		ModelId:      parameter.GetInt64("model_id"),
		ScheduleDate: scheduleDate,
	}

	db := s.db.Model(&schedulesModel.Schedules{})
	db.Where(schedulesModel.Schedules{Id: id})

	err = db.Updates(&data).Error

	return
}

func (s schedulesRepository) DeleteSchedulesById(id int64) (err error) {
	data := schedulesModel.Schedules{}

	db := s.db.Model(&data)
	err = db.Delete(&data, schedulesModel.Schedules{Id: id}).Error

	return
}

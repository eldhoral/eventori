package schedulesModel

import (
	"eldho/eventori/internal/app/model/helperModel"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Schedules struct {
	Id                                  int64     `json:"id"`
	ModelId                             int64     `json:"model_id"`
	ScheduleDate                        time.Time `json:"schedule_date"`
	helperModel.DateAuditModelTimeStamp `json:"-"`
}

type SchedulesRequest struct {
	ModelId      int64  `form:"model_id"`
	ScheduleDate string `form:"schedule_date"`
}

func (Schedules) TableName() string {
	return "model_schedules"
}

func (sr SchedulesRequest) Validate() error {
	return validation.ValidateStruct(&sr,
		validation.Field(&sr.ModelId, validation.Required),
		validation.Field(&sr.ScheduleDate, validation.Required),
	)
}

package cataloguesModel

import (
	"eldho/eventori/internal/app/model/helperModel"
	"eldho/eventori/internal/app/model/schedulesModel"
	"mime/multipart"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Catalogues struct {
	Id                                  int64                       `json:"id"`
	ModelName                           string                      `json:"model_name"`
	ImageUrl                            string                      `json:"image_url"`
	Description                         string                      `json:"description"`
	Schedules                           *[]schedulesModel.Schedules `gorm:"foreignKey:ModelId;references:Id" json:"model_schedules,omitempty"`
	helperModel.DateAuditModelTimeStamp `json:"-"`
}

type CataloguesRequest struct {
	ModelName   string                `form:"model_name"`
	Image       *multipart.FileHeader `form:"image"`
	Description string                `form:"description"`
}

func (Catalogues) TableName() string {
	return "model_catalogues"
}

func (cr CataloguesRequest) Validate() error {
	return validation.ValidateStruct(&cr,
		validation.Field(&cr.ModelName, validation.Required),
		validation.Field(&cr.Image, validation.Required),
		validation.Field(&cr.Description, validation.Required),
	)
}

package cataloguesRepository

import (
	datapaging "eldho/eventori/internal/app/commons/dataPagingHelper"
	"eldho/eventori/internal/app/model/cataloguesModel"

	param "eldho/eventori/internal/pkg/params"
)

type ICataloguesRepository interface {
	ListCatalogues(paging datapaging.Datapaging) (data *[]cataloguesModel.Catalogues, count int64, err error)
	CreateCatalogues(parameter param.Params) (err error)
	GetCataloguesByModelId(modelId int64, paging datapaging.Datapaging) (data *[]cataloguesModel.Catalogues, count int64, err error)
	UpdateCataloguesByModelId(modelId int64, parameter param.Params) (err error)
	DeleteCataloguesByModelId(modelId int64) (err error)
}

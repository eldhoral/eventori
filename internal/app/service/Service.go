package service

import (
	"eldho/eventori/internal/app/commons"
	"eldho/eventori/internal/app/repository"
	"eldho/eventori/internal/app/service/modelService"
)

// Option anything any service object needed
type Option struct {
	commons.Options
	*repository.Repositories
}

type Services struct {
	ModelService modelService.IModelService
}

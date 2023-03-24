package repository

import (
	"eldho/eventori/internal/app/commons"
	"eldho/eventori/internal/app/repository/cataloguesRepository"
	"eldho/eventori/internal/app/repository/schedulesRepository"
)

// Option anything any repo object needed
type Option struct {
	commons.Options
}

type Repositories struct {
	CataloguesRepository cataloguesRepository.ICataloguesRepository
	SchedulesRepository  schedulesRepository.ISchedulesRepository
}

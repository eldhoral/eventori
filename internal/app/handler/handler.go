package handler

import (
	"eldho/eventori/internal/app/commons"
	"eldho/eventori/internal/app/service"
)

// HandlerOption option for handler, including all service
type HandlerOption struct {
	commons.Options
	*service.Services
}

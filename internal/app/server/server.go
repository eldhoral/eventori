package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"eldho/eventori/internal/app/commons"
	"eldho/eventori/internal/app/handler"
	"eldho/eventori/internal/app/service"

	"github.com/rs/zerolog/log"
)

// IServer interface for server
type IServer interface {
	StartApp()
}

type server struct {
	opt      commons.Options
	services *service.Services
}

// NewServer create object server
func NewServer(opt commons.Options, services *service.Services) IServer {
	return &server{
		opt:      opt,
		services: services,
	}
}

func (s *server) StartApp() {
	var srv http.Server

	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		appPort = 8000
	}

	srv.Addr = fmt.Sprintf("%s:%d", os.Getenv("APP_HOST"), appPort)
	hOpt := handler.HandlerOption{
		Options:  s.opt,
		Services: s.services,
	}

	r := Router(hOpt)
	srv.Handler = r

	log.Info().Msgf("[API] HTTP serve at %s\n", srv.Addr)

	if errHTTP := r.Run(":" + strconv.Itoa(appPort)); errHTTP != nil {
		log.Error().Msg(errHTTP.Error())
	}

	log.Info().Msg("Bye")
}

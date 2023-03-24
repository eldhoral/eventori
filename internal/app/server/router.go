package server

import (
	"os"
	"time"

	"eldho/eventori/internal/app/commons/uploadFile"
	"eldho/eventori/internal/app/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/cast"
)

func Router(opt handler.HandlerOption) *gin.Engine {
	modelHandler := handler.ModelHandler{
		HandlerOption: opt,
	}

	setMode := cast.ToBool(os.Getenv("DEBUG_MODE"))
	if setMode {
		gin.SetMode(gin.ReleaseMode)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		gin.SetMode(gin.DebugMode)
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	//routes
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"POST", "DELETE", "GET", "OPTIONS", "PUT"},
		AllowHeaders:           []string{"Origin", "Content-Type", "Authorization", "userid", "REQUEST-ID", "X-SIGNATURE", "Referer", "User-Agent"},
		AllowCredentials:       true,
		ExposeHeaders:          []string{"Content-Length"},
		MaxAge:                 120 * time.Second,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))

	// return 500 if system got panic
	r.Use(gin.Recovery())

	//Maximum memory limit for Multipart forms
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	modelGroup := r.Group("/api/v1")
	{
		modelGroup.GET("/models/list", modelHandler.ListCatalogues)
		modelGroup.GET("/models/:model_id", modelHandler.ListCataloguesByModelId)
		modelGroup.GET("/models/schedules/:model_id", modelHandler.ListSchedulesByModelId)
		modelGroup.POST("/models/create", modelHandler.CreateCatalogues)
		modelGroup.POST("/models/schedules/create", modelHandler.CreateSchedules)
		modelGroup.PATCH("/models/update/:model_id", modelHandler.UpdateCatalogues)
		modelGroup.PATCH("/models/schedules/update/:schedule_id", modelHandler.UpdateSchedules)
		modelGroup.DELETE("/models/delete/:model_id", modelHandler.DeleteCatalogues)
		modelGroup.DELETE("/models/schedules/delete/:schedule_id", modelHandler.DeleteSchedules)
	}

	staticGroup := r.Group("/api/v1")
	{
		staticGroup.Static("/assets/upload/image/", uploadFile.DocumentPath[uploadFile.OtherDocument])
	}

	return r
}

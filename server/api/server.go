package api

import (
	"context"
	"melp-back/data"
	"melp-back/ent"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

type Database struct {
	Map *ent.Client
}

type Controller struct {
	Conn *gin.Engine
	Data Database
}

func setDevStgState(state string) {
	// Read `state` from the environment
	switch {
	case state == "deploy":
		gin.SetMode(gin.ReleaseMode)
	case state == "test":
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func CreateController(state string) *Controller {
	// Set `gin` status
	setDevStgState(state)

	// Start up gin
	router := gin.Default()
	router.Use(corsMiddleware())

	// Start database connection engine
	database := data.New()

	src := Controller{
		Conn: router,
		Data: Database{
			Map: database,
		},
	}

	return &src
}

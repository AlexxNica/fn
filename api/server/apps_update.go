package server

import (
	"net/http"

	"github.com/fnproject/fn/api"
	"github.com/fnproject/fn/api/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleAppUpdate(c *gin.Context) {
	ctx := c.Request.Context()

	wapp := models.AppWrapper{}

	err := c.BindJSON(&wapp)
	if err != nil {
		handleErrorResponse(c, models.ErrInvalidJSON)
		return
	}

	if wapp.App == nil {
		handleErrorResponse(c, models.ErrAppsMissingNew)
		return
	}

	wapp.App.ID = c.MustGet(api.ID).(string)

	err = s.FireBeforeAppUpdate(ctx, wapp.App)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	app, err := s.datastore.UpdateApp(ctx, wapp.App)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	err = s.FireAfterAppUpdate(ctx, wapp.App)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, appResponse{"App successfully updated", app})
}

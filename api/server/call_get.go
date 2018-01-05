package server

import (
	"net/http"

	"github.com/fnproject/fn/api"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleCallGet(c *gin.Context) {
	ctx := c.Request.Context()

	appID := c.MustGet(api.ID).(string)
	callID := c.Param(api.Call)
	callObj, err := s.datastore.GetCall(ctx, appID, callID)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, callResponse{"Successfully loaded call", callObj})
}

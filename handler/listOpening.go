package handler

import (
	"Api_first_1_19/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error list openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}

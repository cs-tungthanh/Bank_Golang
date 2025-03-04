package util

import (
	"net/http"

	"github.com/cs-tungthanh/Bank_Golang/pkg/core"
	"github.com/gin-gonic/gin"
)

func WriteErrorResponse(c *gin.Context, err error) {
	if errSt, ok := err.(core.StatusCodeCarrier); ok {
		c.JSON(errSt.StatusCode(), errSt)
		return
	}

	c.JSON(http.StatusInternalServerError, core.ErrInternalServerError.WithError(err.Error()))
}

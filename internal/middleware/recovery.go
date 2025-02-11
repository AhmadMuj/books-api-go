package middleware

import (
	"fmt"
	"net/http"

	"github.com/AhmadMuj/books-api-go/internal/errors"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		err, ok := recovered.(error)
		if !ok {
			err = fmt.Errorf("%v", recovered)
		}

		appErr := errors.NewInternalError(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, appErr)
	})
}

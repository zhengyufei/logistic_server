package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				c.Abort()
				detail := fmt.Sprint(err)

				c.JSON(http.StatusInternalServerError, &Response{
					Code:    ErrorCodeInternalError,
					Message: detail,
					Data:    struct{}{},
				})

				//if e, ok := err.(error); ok {
				//	logger.Errorf("PANIC: %+v", errors.WithStack(e))
				//} else {
				//	logger.Errorf("PANIC: %v", errors.WithStack(errors.Errorf("%v", err)))
				//}
			}
		}()
		c.Next()
	}
}

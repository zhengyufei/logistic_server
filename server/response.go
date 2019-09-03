package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Write ...
func Write(c *gin.Context, obj interface{}) {
	if obj == nil {
		obj = struct{}{}
	}
	data, err := json.Marshal(Response{
		Code:    ErrorCodeNull,
		Message: ErrorCodeNull.String(),
		Data:    obj,
	})
	// todo
	//log.Debug("return response %s",)
	if err != nil {
		c.JSON(http.StatusOK, &Response{
			Code:    ErrorCodeInternalError,
			Message: err.Error(),
			Data:    struct{}{},
		})
		return
	}
	c.Data(http.StatusOK, "application/json", data)
}

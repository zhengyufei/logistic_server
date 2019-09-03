package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/handler"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/kuaidi100"
)

func K1Callback(c *gin.Context) {
	h := handler.Kuaidi100Handler{}
	req := &kuaidi100.CallbackRequest{}
	res := &kuaidi100.CallbackResponse{}
	err := h.Callback(req, res)
	if err != nil {
		res.Result = false
		res.ReturnCode = "500"
		res.Message = err.Error()
	} else {
		res.Result = true
		res.ReturnCode = "200"
	}

	Write(c, res)
	return
}

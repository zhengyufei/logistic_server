package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/handler"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/log"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/proto"
)

func LogisticQuery(c *gin.Context) {
	h := handler.LogisticHandler{}
	req := &proto.LogisticQueryRequest{}
	res := &proto.LogisticQueryResponse{}
	if err := c.ShouldBind(req); err != nil {
		Write(c, res)
		return
	}

	err := h.QueryLogistic(req, res)
	if err != nil {
		log.Errorf("err %v", err.Error())
	}

	Write(c, res)
	return
}

func PermissionQuery(c *gin.Context) {
	h := handler.LogisticHandler{}
	req := &proto.LogisticPermissionRequest{}
	res := &proto.LogisticPermissionResponse{}
	if err := c.ShouldBind(req); err != nil {
		Write(c, res)
		return
	}

	log.Debugf("shop_id %v", req.ShopId)
	err := h.QueryPermission(req, res)
	if err != nil {
		log.Errorf("err %v", err.Error())
	}

	Write(c, res)
	return
}

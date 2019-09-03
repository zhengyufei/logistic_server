package handler

import (
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/kuaidi100"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/log"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/models"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/mongo"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/proto"
)

type LogisticHandler struct {
}

func (h *LogisticHandler) QueryLogistic(req *proto.LogisticQueryRequest, res *proto.LogisticQueryResponse) error {
	// first judge in whitelist, if has subscribed, query mongo else subscribe now, and query kuaidi100
	b, err := h.GetPermission(req.ShopId)
	if err != nil {
		return err
	}
	if !b {
		return errors.New("No Permission")
	}

	hasSubscribe, err := h.hasSubscribe(req)
	if err != nil {
		return err
	}

	if hasSubscribe {
		err = h.queryLogisticByMongo(req, res)
		if err != nil {
			return err
		}
		return nil
	} else {
		err = h.subscribe(req)
		if err != nil {
			return err
		}

		err = h.queryLogistic(req, res)
		if err != nil {
			return err
		}

		return nil
	}
}

func (h *LogisticHandler) GetPermission(shopId string) (bool, error) {
	session := mongo.GetSession()
	defer session.Close()

	if !bson.IsObjectIdHex(shopId) {
		log.Errorf("shopid %v is not bson", shopId)
		return false, errors.New("shopid is not bson")
	}

	b, err := models.GetLogisticWhiteList(session, bson.ObjectIdHex(shopId))
	if err != nil && err.Error() != mongo.NOT_FOUND {
		return false, err
	}

	return b, nil
}

func (h *LogisticHandler) queryLogisticByMongo(req *proto.LogisticQueryRequest, res *proto.LogisticQueryResponse) error {
	session := mongo.GetSession()
	defer session.Close()

	logistic, err := models.GetLogistic(session, req.Company, req.Number)
	if err != nil {
		return err
	}

	res.State = logistic.State
	res.IsCheck = logistic.IsCheck
	res.State = logistic.State

	for _, s := range logistic.Tracks {
		var data = new(proto.LogisticQueryData)
		data.Context = s.Context
		data.Time = s.Time
		data.FTime = s.FTime
		data.Status = s.Status
		data.AreaCode = s.AreaCode
		data.AreaName = s.AreaName

		res.Data = append(res.Data, data)
	}

	return nil
}

func (h *LogisticHandler) hasSubscribe(req *proto.LogisticQueryRequest) (bool, error) {
	session := mongo.GetSession()
	defer session.Close()

	b, err := models.GetSubscribe(session, req.Company, req.Number)
	if err != nil && err.Error() != mongo.NOT_FOUND {
		return false, err
	}

	return b, nil
}

func (h *LogisticHandler) subscribe(req *proto.LogisticQueryRequest) error {
	res, err := kuaidi100.Subscribe(req.Company, req.Number, req.Phone)
	if err != nil {
		return err
	}

	switch res.ReturnCode {
	case "200", "501":
		session := mongo.GetSession()
		defer session.Close()

		err := models.InsertSubscribe(session, bson.ObjectIdHex(req.ShopId), req.Company, req.Number)
		if err != nil {
			return err
		}
	default:
		err := errors.New("this is a new error")
		return err
	}

	return nil
}

func (h *LogisticHandler) queryLogistic(req *proto.LogisticQueryRequest, res *proto.LogisticQueryResponse) error {
	k1Res, err := kuaidi100.Query(req.Company, req.Number, req.Phone)
	if err != nil {
		return err
	}

	err = saveToMongo(k1Res)
	if err != nil {
		log.Errorf("to mongo error")
	}

	res.Company = k1Res.Company
	res.Number = k1Res.Number
	res.IsCheck = k1Res.IsCheck
	res.State = k1Res.State

	for _, s := range k1Res.Tracks {
		var data = new(proto.LogisticQueryData)
		data.Context = s.Context
		data.Time = s.Time
		data.FTime = s.FTime
		data.Status = s.Status
		data.AreaCode = s.AreaCode
		data.AreaName = s.AreaName

		res.Data = append(res.Data, data)
	}

	return nil
}

func (h *LogisticHandler) QueryPermission(req *proto.LogisticPermissionRequest, res *proto.LogisticPermissionResponse) error {
	b, err := h.GetPermission(req.ShopId)
	if err != nil {
		return err
	}

	res.IsAllow = b

	return nil
}

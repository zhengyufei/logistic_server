package handler

import (
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/kuaidi100"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/models"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/mongo"
)

func saveToMongo(k1Res *kuaidi100.QueryResponse) error {
	session := mongo.GetSession()
	defer session.Close()

	err := models.UpsertLogistic(session, k1Res)
	if err != nil {
		return err
	}
	return nil
}

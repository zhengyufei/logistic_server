package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type LogisticWhiteList struct {
	Id     bson.ObjectId `bson:"_id" json:"id"`
	ShopID bson.ObjectId `bson:"shop_id" json:"shop_id"`
}

func (h *LogisticWhiteList) Collection(s *mgo.Session) *mgo.Collection {
	return s.DB("xdmp").C("logistic_whitelist")
}

func GetLogisticWhiteList(s *mgo.Session, shopId bson.ObjectId) (bool, error) {
	whiteList := &LogisticWhiteList{}
	find := bson.M{"shop_id": shopId}
	if err := whiteList.Collection(s).Find(&find).One(whiteList); err != nil {
		return false, err
	} else {
		return whiteList.ShopID != "", nil
	}
}

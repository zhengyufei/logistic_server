package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/kuaidi100"
)

type Track struct {
	Context  string `bson:"context,omitempty" json:"context,omitempty"`
	Time     string `bson:"time,omitempty" json:"time,omitempty"`
	FTime    string `bson:"ftime,omitempty" json:"ftime,omitempty"`
	Status   string `bson:"status,omitempty" json:"status,omitempty"`
	AreaCode string `bson:"area_code,omitempty" json:"area_code,omitempty"`
	AreaName string `bson:"area_name,omitempty" json:"area_name,omitempty"`
}

type Logistic struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	ShopID       bson.ObjectId `bson:"shop_id,omitempty" json:"shop_id,omitempty"`
	HasSubscribe bool          `bson:"has_subscribe,omitempty" json:"has_subscribe,omitempty"`
	State        string        `bson:"state,omitempty" json:"state,omitempty"`
	IsCheck      string        `bson:"is_check,omitempty" json:"is_check,omitempty"`
	Company      string        `bson:"company,omitempty" json:"company,omitempty"`
	Number       string        `bson:"number,omitempty" json:"number,omitempty"`
	Tracks       []*Track      `bson:"tracks,omitempty" json:"tracks,omitempty"`
}

func (h *Logistic) Collection(s *mgo.Session) *mgo.Collection {
	return s.DB("xdmp").C("logistic")
}

func GetLogistic(s *mgo.Session, company string, number string) (*Logistic, error) {
	logistic := &Logistic{}
	find := bson.M{"company": company, "number": number}

	if err := logistic.Collection(s).Find(&find).One(logistic); err != nil {
		return nil, err
	} else {
		return logistic, nil
	}
}

func UpsertLogistic(s *mgo.Session, k1Res *kuaidi100.QueryResponse) error {
	logistic := &Logistic{}
	logistic.HasSubscribe = true
	logistic.State = k1Res.State
	logistic.IsCheck = k1Res.IsCheck
	logistic.Company = k1Res.Company
	logistic.Number = k1Res.Number

	for _, s := range k1Res.Tracks {
		var track Track
		track.Context = s.Context
		track.Time = s.Time
		track.FTime = s.FTime
		track.Status = s.Status
		track.AreaCode = s.AreaCode
		track.AreaName = s.AreaName
		logistic.Tracks = append(logistic.Tracks, &track)
	}

	filter := bson.M{"company": k1Res.Company, "number": k1Res.Number}
	doc := bson.M{"$set": logistic}
	_, err := logistic.Collection(s).Upsert(filter, doc)
	if err != nil {
		return err
	}

	return nil
}

func GetSubscribe(s *mgo.Session, company string, number string) (bool, error) {
	logistic := &Logistic{}
	find := bson.M{"company": company, "number": number}

	if err := logistic.Collection(s).Find(&find).One(logistic); err != nil {
		return false, err
	} else {
		return logistic.HasSubscribe, nil
	}
}

func InsertSubscribe(s *mgo.Session, shopId bson.ObjectId, company string, number string) error {
	logistic := &Logistic{}
	doc := bson.M{"shop_id": shopId, "has_subscribe": true, "company": company, "number": number}
	err := logistic.Collection(s).Insert(doc)
	if err != nil {
		return err
	}

	return nil
}

type Logistic2 struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	ShopID       bson.ObjectId `bson:"shop_id,omitempty" json:"shop_id,omitempty"`
	HasSubscribe bool          `bson:"has_subscribe,omitempty" json:"has_subscribe,omitempty"`
	State        string        `bson:"state,omitempty" json:"state,omitempty"`
	IsCheck      string        `bson:"is_check,omitempty" json:"is_check,omitempty"`
	Company      string        `bson:"company,omitempty" json:"company,omitempty"`
	Number       string        `bson:"number,omitempty" json:"number,omitempty"`
	Tracks       []*Track      `bson:"tracks,omitempty" json:"tracks,omitempty"`
}

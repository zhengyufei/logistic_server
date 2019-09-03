package handler

import (
	"encoding/json"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/kuaidi100"
	"gitlab.xiaoduoai.com/ecrobot/logistic_server/log"
)

type Kuaidi100Handler struct {
}

func (h *Kuaidi100Handler) Callback(req *kuaidi100.CallbackRequest, res *kuaidi100.CallbackResponse) error {
	cbParam := &kuaidi100.CallbackParam{}
	err := json.Unmarshal([]byte(req.Param), cbParam)
	if err != nil {
		log.Errorf("unmarshal CallbackParm %+v", err)
		return err
	}

	return nil
}

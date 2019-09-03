package kuaidi100

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
)

func Query(com string, num string, phone string) (*QueryResponse, error) {
	var req QueryRequest
	var param QueryParam
	//var res proto.QueryResponse

	req.Customer = viper.GetString("kuaidi100.Customer")
	param.Com = com
	param.Num = num
	param.Phone = phone
	param.From = ""
	param.To = ""
	param.Resultv2 = "1"
	paramJson, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	req.Param = string(paramJson)

	tmpStr := req.Param + viper.GetString("kuaidi100.Key") + req.Customer
	req.Sign = fmt.Sprintf("%X", md5.Sum([]byte(tmpStr)))
	req.Sign = strings.ToUpper(req.Sign)

	req2, err := sling.New().Post(QueryUri).QueryStruct(req).Request()
	if err != nil {
		return nil, err
	}
	var client http.Client
	resp, err := client.Do(req2)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var res = new(QueryResponse)
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Subscribe(com string, num string, phone string) (*SubscribeResponse, error) {
	var req SubscribeRequest
	var param SubscribeParam

	req.Schema = "json"
	param.Company = com
	param.Number = num
	param.Key = viper.GetString("kuaidi100.Key")
	param.Parameters = new(SubscribeParamParameters)
	param.Parameters.Phone = phone
	param.Parameters.ResultV2 = "1"
	param.Parameters.CallbackUrl = viper.GetString("kuaidi100.Callback")
	paramJson, err := json.Marshal(param)
	if err != nil {
		fmt.Println("生成json字符串错误")
		return nil, err
	}
	req.Param = string(paramJson)

	req2, err := sling.New().Post(SubscribeUri).QueryStruct(req).Request()
	if err != nil {
		return nil, err
	}
	var client http.Client
	resp, err := client.Do(req2)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var res = new(SubscribeResponse)
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

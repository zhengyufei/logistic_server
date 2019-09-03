package kuaidi100

const (
	QueryUri     = "https://poll.kuaidi100.com/poll/query.do"
	SubscribeUri = "https://poll.kuaidi100.com/poll"
)

type QueryParam struct {
	Com      string `json:"com"`
	Num      string `json:"num"`
	Phone    string `json:"phone,omitempty"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Resultv2 string `json:"resultv2,omitempty"`
}

//*
// 快递100实时查询接口
// Post
type QueryRequest struct {
	Customer string `url:"customer"`
	Sign     string `url:"sign"`
	Param    string `url:"param"`
}

type QueryData struct {
	Context  string `json:"context,omitempty"`
	Time     string `json:"time,omitempty"`
	FTime    string `json:"ftime,omitempty"`
	Status   string `json:"status,omitempty"`
	AreaCode string `json:"areaCode,omitempty"`
	AreaName string `json:"areaName,omitempty"`
}

type QueryResponse struct {
	State   string       `json:"state,omitempty"`
	IsCheck string       `json:"ischeck,omitempty"`
	Company string       `json:"com,omitempty"`
	Number  string       `json:"nu,omitempty"`
	Tracks  []*QueryData `json:"data,omitempty"`
}

type SubscribeParamParameters struct {
	CallbackUrl string `json:"callbackurl,omitempty"`
	//string salt       // 可选 签名用随机字符串。32位自定义字符串。添加该参数，则推送的时候会增加sign给贵司校验消息的可靠性
	ResultV2 string `json:"resultv2,omitempty"`
	AutoCom  string `json:"autoCom,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type SubscribeParam struct {
	Company string `json:"company,omitempty"`
	Number  string `json:"number,omitempty"`
	//string from;        // 可选 出发地城市，省-市-区，非必填，填了有助于提升签收状态的判断的准确率，请尽量提供
	//string to;          // 可选 目的地城市，省-市-区，非必填，填了有助于提升签收状态的判断的准确率，且到达目的地后会加大监控频率，请尽量提供
	Key        string                    `json:"key,omitempty"`
	Parameters *SubscribeParamParameters `json:"parameters,omitempty"`
}

//*
// 快递100订阅请求接口
// Post
type SubscribeRequest struct {
	Schema string `url:"schema"`
	Param  string `url:"param"`
}

type SubscribeResponse struct {
	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	//
	// 200: 提交成功
	// 701: 拒绝订阅的快递公司
	// 700: 订阅方的订阅数据存在错误（如不支持的快递公司、单号为空、单号超长等）或错误的回调地址
	// 702: POLL:识别不到该单号对应的快递公司
	// 600: 您不是合法的订阅者（即授权Key出错）
	// 601: POLL:KEY已过期
	// 500: 服务器错误（即快递100的服务器出理间隙或临时性异常，有时如果因为不按规范提交请求，比如快递公司参数写错等，也会报此错误）
	// 501:重复订阅（请格外注意，501表示这张单已经订阅成功且目前还在跟踪过程中（即单号的status=polling），快递100的服务器会因此忽略您最新的此次订阅请求，从而返回501。一个运单号只要提交一次订阅即可，若要提交多次订阅，请在收到单号的status=abort或shutdown后隔半小时再提交订阅
	ReturnCode string `json:"returnCode,omitempty"`
	Message    string `json:"message,omitempty"`
}

type CallbackParam struct {
	// 监控状态:
	// polling:监控中，
	// shutdown:结束，
	// abort:中止，
	// updateall：重新推送。
	// 其中当快递单为已签收时status=shutdown，当message为“3天查询无记录”或“60天无变化时”status= abort ，对于stuatus=abort的状度，需要增加额外的处理逻辑
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	// 快递公司编码是否出错
	// 0为本推送信息对应的是贵司提交的原始快递公司编码，
	// 1为本推送信息对应的是我方纠正后的新的快递公司编码。
	// 一个单如果我们连续3天都查不到结果，我方会（1）判断一次贵司提交的快递公司编码是否正确，如果正确，给贵司的回调接口（callbackurl）推送带有如下字段的信息：autoCheck=0、comOld与comNew都为空；（2）如果贵司提交的快递公司编码出错，我们会帮忙用正确的快递公司编码+原来的运单号重新提交订阅并开启监控（后续如果监控到单号有更新就给贵司的回调接口（callbackurl）推送带有如下字段的信息：autoCheck=1、comOld=原来的公司编码、comNew=新的公司编码）；并且给贵方的回调接口（callbackurl）推送一条含有如下字段的信息：status=abort、autoCheck=0、comOld为空、comNew=纠正后的快递公司编码。
	AutoCheck  string         `json:"autoCheck,omitempty"`
	ComOld     string         `json:"comOld,omitempty"`
	ComNew     string         `json:"comNew,omitempty"`
	LastResult *QueryResponse `json:"lastResult,omitempty"`
	DestResult *QueryResponse `json:"destResult,omitempty"`
}

//*
// 快递100 callback接口
// Post
type CallbackRequest struct {
	Param string `json:"param,omitempty" form:"param"`
}

type CallbackResponse struct {
	Result bool `json:"result,omitempty"`
	//
	// 200: 接收成功
	// 500: 服务器错误
	// 其他错误请自行定义
	ReturnCode string `json:"returnCode,omitempty"`
	Message    string `json:"message,omitempty"`
}

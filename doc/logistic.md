# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [logistic.proto](#logistic.proto)
    - [LogisticPermissionRequest](#.LogisticPermissionRequest)
    - [LogisticPermissionResponse](#.LogisticPermissionResponse)
    - [LogisticQueryData](#.LogisticQueryData)
    - [LogisticQueryRequest](#.LogisticQueryRequest)
    - [LogisticQueryResponse](#.LogisticQueryResponse)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="logistic.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## logistic.proto



<a name=".LogisticPermissionRequest"></a>

### LogisticPermissionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shop_id | [string](#string) |  | @inject_tag: form:&#34;shop_id&#34;

必须 |






<a name=".LogisticPermissionResponse"></a>

### LogisticPermissionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_allow | [bool](#bool) |  |  |






<a name=".LogisticQueryData"></a>

### LogisticQueryData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| context | [string](#string) |  | 物流轨迹节点内容 |
| time | [string](#string) |  | 时间，原始格式 |
| fTime | [string](#string) |  | 格式化后时间 |
| status | [string](#string) |  | 本数据元对应的签收状态。只有在开通签收状态服务（见上面&#34;status&#34;后的说明）且在订阅接口中提交resultv2标记后才会出现 |
| areaCode | [string](#string) |  | 本数据元对应的行政区域的编码，只有在开通签收状态服务（见上面&#34;status&#34;后的说明）且在订阅接口中提交resultv2标记后才会出现 |
| areaName | [string](#string) |  | 本数据元对应的行政区域的名称，开通签收状态服务（见上面&#34;status&#34;后的说明）且在订阅接口中提交resultv2标记后才会出现 |






<a name=".LogisticQueryRequest"></a>

### LogisticQueryRequest
快递100实时查询接口
Post


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shop_id | [string](#string) |  | @inject_tag: form:&#34;shop_id&#34;

必须 |
| company | [string](#string) |  | @inject_tag: form:&#34;company&#34;

必须 查询的快递公司的编码，一律用小写字母 |
| number | [string](#string) |  | @inject_tag: form:&#34;number&#34;

必须 查询的快递单号， 单号的最大长度是32个字符 |
| phone | [string](#string) |  | @inject_tag: form:&#34;phone&#34;

可选 收件人或寄件人的手机号或固话（顺丰单号必填，也可以填写后四位，如果是固话，请不要上传分机号） |






<a name=".LogisticQueryResponse"></a>

### LogisticQueryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [string](#string) |  | 快递单当前签收状态，包括0在途中、1已揽收、2疑难、3已签收、4退签、5同城派送中、6退回、7转单等7个状态 |
| isCheck | [string](#string) |  | 是否签收标记 |
| company | [string](#string) |  | 快递公司编码,一律用小写字母，点击查看快递公司编码 |
| number | [string](#string) |  | 快递单号 |
| data | [LogisticQueryData](#LogisticQueryData) | repeated | 数组，包含多个对象，每个对象字段如展开所示 |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |


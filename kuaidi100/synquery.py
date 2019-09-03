# coding = utf-8
import sys,os
import requests,json,hashlib

key = 'jJKdiSLj9800'              #客户授权key
customer = 'FF2E166517654ED1C93A7C94FFB72795'    #查询公司编号
param = {}
param['com'] = 'shunfeng'         #快递公司编码
param['num'] = '232005784132'     #快递单号
param['phone'] = '13408544339'    #手机号
param['from'] = ''                #出发地城市
param['to'] = ''                  #目的地城市
param['resultv2'] = '1'           #开启行政区域解析
pjson = json.dumps(param)         #转json字符串

postdata = {}
postdata['customer'] = customer   #查询公司编号
postdata['param'] = pjson         #参数数据

#签名加密
str = pjson + key + customer
print(str)
md = hashlib.md5()
md.update(str.encode())
sign = md.hexdigest()
postdata['sign'] = sign.upper()  #加密签名
print(postdata['sign'])

url = 'http://poll.kuaidi100.com/poll/query.do'  #实时查询请求地址

result = requests.post(url, postdata)  #发送请求
print(result.text)                     #返回数据

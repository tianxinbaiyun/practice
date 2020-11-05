@[TOC]善行银联接入

## 1.需求简介

因为公司的服务对象寺院，部分受政策的限制，
无法自行申请微信商户平台支付以及支付宝商户平台支付；
善行平台不做资金沉淀，即款项直接打到寺院自己的账户。
为解决这些问题，公司接入银联平台，完成客户的在线支付，寺院的分账处理。

银联相关API接口查看<银商大华捷通平台综合支付接口规范_v2.0.6>


## 2.相关接口

### 2.1 在线支付

接口地址

```
POST /pay
```

文件目录
```
[$GOPATH]/git.liyula.net/shanxing/ts-core/api/pay/handler/api/payment/chinaums.go
```

对应方法
```
func (this *PayHandler) GetChinaumsPayURL(order *ordertype.OrderTypeHandler) (payURL string, rsp map[string]interface{}, err error) 
```

接口文档地址

[统一支付](https://wiki.dev.renrenfo.cn/document/index?document_id=190)


### 2.2 支付回调

接口地址

```
POST /pay/query
```

文件目录
```
[$GOPATH]/git.liyula.net/shanxing/ts-core/api/pay/handler/api/notify.go
```

对应方法
```
func (p *Chinaums) PayNotify(c *gin.Context) 
```

接口文档地址

无


### 2.3 支付查单

接口地址

```
POST /pay/query
```

文件目录
```
[$GOPATH]/git.liyula.net/shanxing/ts-core/api/pay/handler/api/payment/chinaums.go
```

对应方法
```
func (this QueryHandler) QueryChinaumsOrder() (code int, err error)
```

接口文档地址

[交易查询](https://wiki.dev.renrenfo.cn/document/index?document_id=144)

##  3.总结

1.目前已完成银联的在线支付,支付回调,交易查询测试(H5和APP类型)
2.寺院分账、退款暂未接入
3.相关账号配置,在ts_mall.payment_temple表中,type=4表示银联配置





package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
	"sort"
	"strings"
)

// ChinaumsOrder ChinaumsOrder
type ChinaumsOrder struct {
	Version        string `json:"version"`        //版本号,必填:固定值 v2
	OrderNo        string `json:"order_no"`       //支付单号,必填,单号唯一
	BusiOrderNo    string `json:"busi_order_no"`  //商户业务订单号,选填:业务订单号,支付通知会回传
	MerID          string `json:"mer_id"`         //商户 ID,必填,联系大华对接人员获取,为商户分配的唯一编号
	Cod            string `json:"cod"`            //金额,必填,单位元
	Qrtype         string `json:"qrtype"`         //使 用 方 式 :
	Payway         string `json:"payway"`         //支付方式
	Memo           string `json:"memo"`           //备注
	OrderDesc      string `json:"orderDesc"`      //订单信息
	SubOpenID      string `json:"subOpenID"`      //微信用户 id
	Dscode         string `json:"dscode"`         //分账编码
	BizCode        string `json:"bizCode"`        //分账编码
	Fee            string `json:"fee"`            //分账金额
	SubOrders      string `json:"subOrders"`      //分账商户明细,
	PlatformAmount string `json:"platformAmount"` //平台金额,单位分
	FixBuyer       string `json:"fixBuyer"`       //是否需要实名认证
	Name           string `json:"name"`           //实名认证姓名
	Mobile         string `json:"mobile"`         //实名认证手机号
	CertNo         string `json:"certNo"`         //实名认证证件号
	CertType       string `json:"certType"`       //实名认证证件类型
	TransMid       string `json:"transMid"`       //实际交易商户号
	TransTid       string `json:"transTid"`       //实际交易终端号
	TransType      string `json:"transType"`      //交易类型
	NotifyURL      string `json:"notifyURL"`      //异步通知地址
	ReturnURL      string `json:"returnURL"`      //同步跳转 url
	SubAppID       string `json:"subAppID"`       //申请微信支付 appid
	EmployeeNo     string `json:"employeeNo"`     //操作员工号
	SignType       string `json:"signType"`       //签名方式
	ExpireTime     string `json:"expireTime"`     //订单过期时间
	Mac            string `json:"mac"`            //签名
}

func main2() {
	data := &ChinaumsOrder{
		Version:        "v2",
		OrderNo:        "123456",
		BusiOrderNo:    "",
		MerID:          "123456",
		Cod:            "123456",
		Qrtype:         "123456",
		Payway:         "123456",
		Memo:           "123456",
		OrderDesc:      "123456",
		SubOpenID:      "adfadf",
		Dscode:         "asdf",
		BizCode:        "adf",
		Fee:            "adf",
		SubOrders:      "adf",
		PlatformAmount: "adf",
		FixBuyer:       "",
		Name:           "asdf",
		Mobile:         "adf",
		CertType:       "",
		TransMid:       "",
		TransTid:       "zxc",
		TransType:      "zxc",
		NotifyURL:      "zxc",
		ReturnURL:      "vcb",
		SubAppID:       "cvb",
		EmployeeNo:     "vb",
		SignType:       "SHA256",
		ExpireTime:     "789",
		Mac:            "789",
	}
	err := CreateSignData(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
	return
}

//CreateSignData 创建加密数据
func CreateSignData(order *ChinaumsOrder) (err error) {
	str, err := CreatePreSignString(order)
	if err != nil {
		fmt.Println(err)
		return
	}
	var sign string
	switch order.SignType {
	case "SM3":
		sign = SignSm3(str)
	case "MD5":
		sign = SignMd5(str)
	case "SHA256":
		sign = SignSha256(str)
	default:
		sign = SignSm3(str)
	}
	fmt.Println("sign:", sign)
	order.Mac = sign
	return
}

// SignSha256 SignSha256
func SignSha256(preSign string) (sign string) {
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write([]byte(preSign))
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	sign = hex.EncodeToString(bytes)
	return
}

// SignMd5 SignMd5
func SignMd5(preSign string) (sign string) {
	h := md5.New()
	h.Write([]byte(preSign))
	return hex.EncodeToString(h.Sum(nil))
}

// SignSm3 SignSm3
func SignSm3(preSign string) (sign string) {
	hash := sm3.New()
	hash.Write([]byte(preSign))
	result := hash.Sum(nil)
	sign = hex.EncodeToString(result)
	//println("sm3 hash = ",sign)
	//hash := sm3.Sm3Sum([]byte(preSign))
	//println("sm3 hash = ",hex.EncodeToString(hash))
	return
}

//CreatePreSignString 空值 和 下标是sign 去掉
func CreatePreSignString(order *ChinaumsOrder) (str string, err error) {
	b, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err)
		return
	}
	//转成map
	var mapOrder map[string]string
	err = json.Unmarshal(b, &mapOrder)
	if err != nil {
		fmt.Println("JsonToMap err: ", err)
	}
	//转成arr,进行排序
	var newMp = make([]string, 0)
	for i := range mapOrder {
		newMp = append(newMp, i)
	}
	sort.Strings(newMp)
	var sb strings.Builder
	for _, i2 := range newMp {
		if i2 == "mac" || mapOrder[i2] == "" {
			continue
		} else {
			sb.WriteString(i2)
			sb.WriteString("=")
			sb.WriteString(mapOrder[i2])
			sb.WriteString("&")
		}
	}
	str = sb.String()

	checkStr := "1111111111111111111111111111111111111111111111111111111111111111"
	str += checkStr
	fmt.Println("PreSignString:", str)
	return
}

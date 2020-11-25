package util

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"io"
	"math"
	"math/big"
	rand2 "math/rand"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

// GetCurrentPath 获取当前文件所在的目录
func GetCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// GetExecpath 获取当前程序运行目录
func GetExecpath() string {
	execpath, _ := os.Executable() // 获得程序路径
	path := filepath.Dir(execpath)
	return strings.Replace(path, "\\", "/", -1)
}

// Float64 Float64
func Float64(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

// GetPath 获取文件的路径
func GetPath(filePath string) string {
	mainPath := GetExecpath()
	path := mainPath + "/" + filePath
	_, err := os.Stat(path)
	if err != nil {
		path = filePath
	}
	return path
}

// GetTemplatesPath 获取模板路径
func GetTemplatesPath(name string) string {
	//当前程序运行的目录，获取文件
	filePath := GetExecpath() + "/../templates/"

	//目录不存在，从指定的目录找
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		filePath = GetCurrentPath() + "/../templates/"
	}

	return filePath + name
}

// StrToTime 字符串转换成时间
func StrToTime(str string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		panic(err)
	}
	return t
}

// StrTime StrTime
//* @des 时间转换函数
//* @param timeStr string
//* @return string
func StrTime(timeStr string) string {
	atime := TimeStringToInt(timeStr)
	var byTime = []int64{365 * 24 * 60 * 60, 30 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}
	var unit = []string{"年前", "个月前", "天前", "小时前", "分钟前", "秒前"}
	now := time.Now().Unix()
	ct := now - atime
	if ct < 0 {
		return "刚刚"
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = MergeString(tempStr, unit[i])
		}
		break
	}
	return res
}

// MergeString MergeString
//* @des 拼接字符串
//* @param args ...string 要被拼接的字符串序列
//* @return string
func MergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
	}
	return buffer.String()
}

// TimeStringToInt TimeStringToInt
func TimeStringToInt(timeSting string) int64 {
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeSting, time.Local)
	return theTime.Unix()
}

// DateStringToInt DateStringToInt
func DateStringToInt(timeSting string) int64 {
	theTime, _ := time.ParseInLocation("2006-01-02", timeSting, time.Local)
	return theTime.Unix()
}

// Md5 Md5
func Md5(str string) string {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(str))
	Result := Md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", Result)
}

// Md5File 文件MD5
func Md5File(file multipart.File) string {
	md5 := md5.New()
	io.Copy(md5, file)
	MD5Str := hex.EncodeToString(md5.Sum(nil))
	return MD5Str
}

// Hash Hash
func Hash(str string) string {
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(str))
	Result := Sha1Inst.Sum([]byte(""))

	return fmt.Sprintf("%x", Result)
}

// CreateDateDir 根据日期创建目录
func CreateDateDir(Path string) string {
	folderName := time.Now().Format("2006/01/02/1504")
	folderPath := filepath.Join(Path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, os.ModeAppend|os.ModePerm)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
	return folderName
}

// CreateDir 创建目录
func CreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModeAppend|os.ModePerm)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

// StdChars StdChars
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// NewLenChars NewLenChars
func NewLenChars(length int) string {
	if length == 0 {
		return ""
	}
	clen := len(StdChars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // Skip this number to avoid modulo bias.
			}
			b[i] = StdChars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

// VerifyEmailFormat email verify
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyMobileFormat mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// GetNum 根据奇偶数返回数字
func GetNum() uint8 {
	timeNum := time.Now().Unix()
	if timeNum%2 == 0 {
		return 1
	}
	return 2
}

// CreateCaptcha 生成6为随机数
func CreateCaptcha() string {
	return fmt.Sprintf("%06v", rand2.New(rand2.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// CreateOrderSn 创建订单号
func CreateOrderSn(pre string) string {
	str := time.Now().Format("20060102150405") + fmt.Sprintf("%06v", rand2.New(rand2.NewSource(time.Now().UnixNano())).Int31n(1000000))

	str = pre + str
	return str
}

// CreateNewOrderSn CreateNewOrderSn
//* 订单号
//* 下单渠道1位+支付渠道1位+订单类型1位+随机数1位+时间戳10位+用户4位
//* userID 分解成后2位和后第三位和第四位，位置颠倒
//* payChannel 支付渠道，9未知渠道
//11支付宝电脑网页支付,12支付宝手机网页支付,13支付宝APP支付,14支付宝刷脸付,15支付宝当面付,19支付宝未知渠道支付
//21微信付款码支付,22微信JSAPI支付,23微信native扫码支付,24微信app支付,25 H5支付,26小程序支付,27人脸支付,29支付宝未知渠道支付
//39 线下未知渠道支付
//* payCode 支付方式1位,1支付宝alipay，2微信wechat，3线下支付cod，9其他other
//* orderType 订单类型(1寺院供养，2经书助印，3寺院建设，4寺院捐助，5在线祈福，6佛事预约-供灯，7.佛事预约-祈福，8佛事预约-往生,9活动供养,10直播礼物,11直播供养)
func CreateNewOrderSn(userID uint32, payChannel uint32, payID uint32, orderType uint32) (orderSn string) {
	timestamp := time.Now().Unix()
	if payChannel == 0 {
		payChannel = 9
	} else {
		payChannel = payChannel % 10
	}
	if payID == 0 {
		payID = 9
	}
	//用户后两位
	userID1 := userID % 100
	//用户后第三位和第四位
	userID2 := (userID / 100) % 100
	orderSn = fmt.Sprintf("%1d%1d%1d%1d%010d%02d%02d", payChannel, payID, orderType, rand2.Intn(9), timestamp, userID1, userID2)
	return
}

// FilterEmoji 输入字符串，表情过滤
func FilterEmoji(content string) string {
	newContent := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			newContent += string(value)
		}
	}
	return newContent
}

// RandString 短信系统随机字符串
func RandString(len int) string {
	r := rand2.New(rand2.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	timeNow := time.Now().Format("20060102150405")
	return timeNow + string(bytes)
}

// Str2bytes 字符串转字符数组
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2str 字符数组转字符串
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StructToMap 结构体转map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

// Random Random
func Random(strings []interface{}, length int) []interface{} {

	for i := len(strings) - 1; i > 0; i-- {
		num := rand2.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	return strings[:length]
}

// CreateUUID 获取uuid
func CreateUUID() string {
	return uuid.New().String()
}

// PathExists PathExists
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// EnsureDir EnsureDir
func EnsureDir(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return
		}
	}
	return
}

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// GetLastDateOfMonth 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetTwoTimeDuration 获取两个时间相差时长
func GetTwoTimeDuration(startTime, endTime string) (dutarion string) {
	var (
		hour, min, sec int
	)
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if t1.Before(t2) {
		diff := int(t2.Unix() - t1.Unix())
		sec = diff % 60
		min = (diff / 60) % 60
		hour = diff / 3600
		dutarion = fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
	}
	return
}

// Krand 随机字符串
//kind 0:纯数字,1:小写字母,2:大写字母,3:数字、大小写字母
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand2.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand2.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand2.Intn(scope))
	}
	return result
}

// RangeRand RangeRand
func RangeRand(min, max int64) int64 {
	if min > max {
		return 0
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	return min + result.Int64()
}

// RangeSplit 将一个数字分成多个数
func RangeSplit(num int64, count int64) (res []int64) {
	if num < count {
		return
	}
	if count == 1 {
		res = append(res, num)
	} else {
		n1 := num - (count-1)*(num/count)

		//rand2.Seed(time.Now().Unix())
		//r := int64(rand2.Intn(int(n1)))

		r := RangeRand(num/count, n1)

		r2 := num - r
		res = RangeSplit(r2, count-1)
		res = append(res, r)
	}
	return
}

// FloatLessThan 小于 <
func FloatLessThan(f1, f2 float64) bool {
	n1 := decimal.NewFromFloat(f1)
	n2 := decimal.NewFromFloat(f2)
	return n1.LessThan(n2)
}

// FloatLessThanOrEqual 小于等于 <=
func FloatLessThanOrEqual(f1, f2 float64) bool {
	n1 := decimal.NewFromFloat(f1)
	n2 := decimal.NewFromFloat(f2)
	return n1.LessThanOrEqual(n2)
}

// FloatGreaterThan 大于 >
func FloatGreaterThan(f1, f2 float64) bool {
	n1 := decimal.NewFromFloat(f1)
	n2 := decimal.NewFromFloat(f2)
	return n1.GreaterThan(n2)
}

// FloatGreaterThanOrEqual 大于等于 >=
func FloatGreaterThanOrEqual(f1, f2 float64) bool {
	n1 := decimal.NewFromFloat(f1)
	n2 := decimal.NewFromFloat(f2)
	return n1.GreaterThanOrEqual(n2)
}

// FloatEqual 是否相等
func FloatEqual(f1, f2 float64) bool {
	n1 := decimal.NewFromFloat(f1)
	n2 := decimal.NewFromFloat(f2)

	if n1.Cmp(n2) == 0 {
		return true
	}
	return false
}

// FloatAdd 浮点加
func FloatAdd(x float64, y float64, more ...float64) float64 {

	floatX := new(big.Float).SetFloat64(x)
	floatY := new(big.Float).SetFloat64(y)
	result := new(big.Float).Add(floatX, floatY)
	if len(more) > 0 {
		for _, m := range more {
			floatM := new(big.Float).SetFloat64(m)
			result = new(big.Float).Add(result, floatM)
		}
	}

	f, _ := strconv.ParseFloat(result.String(), 64)
	return f
}

// FloatSub 浮点减
func FloatSub(x float64, y float64, more ...float64) float64 {

	floatX := new(big.Float).SetFloat64(x)
	floatY := new(big.Float).SetFloat64(y)
	result := new(big.Float).Sub(floatX, floatY)
	if len(more) > 0 {
		for _, m := range more {
			floatM := new(big.Float).SetFloat64(m)
			result = new(big.Float).Sub(result, floatM)
		}
	}

	f, _ := strconv.ParseFloat(result.String(), 64)
	return f
}

// FloatMul 浮点乘
func FloatMul(x float64, y float64, more ...float64) float64 {

	floatX := new(big.Float).SetFloat64(x)
	floatY := new(big.Float).SetFloat64(y)
	result := new(big.Float).Mul(floatX, floatY)
	if len(more) > 0 {
		for _, m := range more {
			floatM := new(big.Float).SetFloat64(m)
			result = new(big.Float).Mul(result, floatM)
		}
	}

	f, _ := strconv.ParseFloat(result.String(), 64)
	return f
}

// FloatQuo 浮点除
func FloatQuo(x float64, y float64, more ...float64) float64 {

	floatX := new(big.Float).SetFloat64(x)
	floatY := new(big.Float).SetFloat64(y)
	result := new(big.Float).Quo(floatX, floatY)
	if len(more) > 0 {
		for _, m := range more {
			floatM := new(big.Float).SetFloat64(m)
			result = new(big.Float).Quo(result, floatM)
		}
	}

	f, _ := strconv.ParseFloat(result.String(), 64)
	return f
}

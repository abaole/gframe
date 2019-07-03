package gno

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/rs/xid"

	"github.com/abaole/gframe/utils/gdate"
	uuid "github.com/satori/go.uuid"
)

//CreateTradeNo 生成系统交易号
//06123xxxxx
//sum 最少10位,sum 表示全部单号位数
func CreateTradeNo(sum int) string {
	//年
	strs := time.Now().Format("06")
	//一年中的第几天
	days := strconv.Itoa(gdate.GetDaysInYearByThisYear())
	count := len(days)
	if count < 3 {
		//重复字符0
		days = strings.Repeat("0", 3-count) + days
	}
	//组合
	strs += days
	//剩余随机数
	sum = sum - 5
	if sum < 1 {
		sum = 5
	}
	//0~9999999的随机数
	//rand.Seed(time.Now().UnixNano())
	pow := math.Pow(10, float64(sum)) - 1

	result := strconv.Itoa(rand.Intn(int(pow)))
	count = len(result)
	if count < sum {
		//重复字符0
		result = strings.Repeat("0", sum-count) + result
	}
	//组合
	strs += result
	return strs
}

// 生带前缀ID
func GenPrefixNo(code string) string {

	nt := strconv.FormatInt(time.Now().Unix(), 10)

	xid := xid.New()
	rs := []rune(xid.String())
	length := len(rs)
	xStr := string(rs[length-6:])
	return code + nt + strings.ToUpper(xStr)
}

// 生成流水号
func GenTrxNo() string {
	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	result := strconv.Itoa(rand.Intn(1000))
	return now + result
}

// 生成KEY
func GenPayKey() string {
	id := xid.New()
	return id.String()
}

func RandBetween(min, max float64) int {
	n := int(max - min)
	if n > 1000 {
		n = 1000
	}
	if n < 1 {
		n = 1
	}
	if max < 1000 {
		max = 1000
	}
	return rand.Intn(n) + int(max-1000)
}

func RandWeight(total int) int {
	return rand.Intn(total) + 1
}

func GetUUID() string {
	return uuid.NewV4().String()
}

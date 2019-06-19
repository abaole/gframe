package gno

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/abaole/gframe/utils/gdate"
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

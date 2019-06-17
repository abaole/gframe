package utils


import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母

)

var kinds = [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}} //初始化字符池

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, result := kind, make([]byte, size)
	isAll := kind > KC_RAND_KIND_UPPER || kind < KC_RAND_KIND_NUM
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

//随机5位数字
func RandString(size int) string {
	return string(Krand(size, KC_RAND_KIND_NUM))
}

// RandFileName 随机生成一个文件名
func RandFileName() string {
	return strconv.FormatInt(time.Now().Unix(), 10) + string(Krand(6, KC_RAND_KIND_NUM))
}

//随机5位数字
func RandMobile4() string {
	return string(Krand(4, KC_RAND_KIND_NUM))
}

//随机1位数字
func RandMobile1() string {
	return string(Krand(1, KC_RAND_KIND_NUM))
}

//生成32随机字符串
func RandAll32() string {
	return string(Krand(32, KC_RAND_KIND_ALL))
}

//	生成一个唯一账号
func RandAccount() string {
	return string(Krand(3, KC_RAND_KIND_LOWER)) + fmt.Sprint(time.Now().UnixNano())
}

//	生成一个唯一订单号
func RandOrderNo() string {
	return string(Krand(5, KC_RAND_KIND_LOWER)) + fmt.Sprint(time.Now().UnixNano())
}

func GenCode() string {
	now := time.Now()
	return now.Format("2006010215") + strconv.Itoa(now.Second()) + string(Krand(4, KC_RAND_KIND_NUM))
}
package gstr

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

//GenerateNonceString 生成随机字符串
func GenerateNonceString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//substr 截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	if length == 0 {
		length = rl
	}
	end = start + length - 1

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func SubBetween(str string, start, end int) string {
	rs := []byte(str)
	rl := len(rs)
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// 中文转码
func UTF82GBK(src []byte) (string, error) {
	reader := transform.NewReader(strings.NewReader(string(src)), simplifiedchinese.GBK.NewEncoder())
	if buf, err := ioutil.ReadAll(reader); err != nil {
		return "", err
	} else {
		return string(buf), nil
	}
}

func WordWrap(text string, column int) string {
	if len(text) <= column {
		return text
	}
	initialPart := text[:column]
	pos := strings.LastIndex(initialPart, " ")
	if pos >= 0 {
		return text[:pos] + "\n" + WordWrap(text[pos+1:], column)
	}
	return initialPart + "\n" + WordWrap(text[column:], column)
}

func ContainStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func WxRate(rate float64) string {
	TR := fmt.Sprintf("%.2f", rate)
	rs := []rune(TR)
	rl := len(rs)
	if string(rs[rl-1]) == "0" {
		rs = rs[:rl-1]
	}

	return string(rs) + "%"
}

func DeleteHttp(url string) string {
	rs := []rune(url)
	result := ""
	if strings.Contains(url, "http://") {
		result = string(rs[7:])
	}
	if strings.Contains(url, "https://") {
		result = string(rs[8:])
	}

	return result
}

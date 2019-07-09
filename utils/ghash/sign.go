package ghash

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sort"
	"strings"
)

const (
	EncryptMD5           = iota //md5加密
	EncryptSHA1withRsa          //SHA1withRsa
	EncryptSHA256withRsa        //SHA256withRsa
	EncryptMD5withRsa           //MD5withRsa
)

type kvPair struct {
	k, v string
}

type kvPairs []kvPair

func (t kvPairs) Sort() {
	sort.Slice(t, func(i, j int) bool {
		return t[i].k < t[j].k
	})
}

func (t kvPairs) Join() string {
	var args []string
	for _, kv := range t {
		args = append(args, kv.k+"="+kv.v)
	}
	return strings.Join(args, "&")
}

// 生成sign
func MakeSignMD5(data map[string]string) {
	p := kvPairs{}
	// 剔除空值 和 sign
	for k, v := range data {
		if !(v == "" || k == "sign") {
			p = append(p, kvPair{k, v})
		}
	}
	p.Sort()
	data["sign"] = Md5String(p.Join())
}

func MakeSignAbc(data map[string]string) {
	p := kvPairs{}
	// 剔除空值 和 sign
	for k, v := range data {
		if !(v == "" || k == "sign") {
			p = append(p, kvPair{k, v})
		}
	}
	p.Sort()

	param := data["partnerKey"]
	for _, v := range p {
		//if !(v.k == "partnerKey" || v.k == "sign") {
		param += v.v
		//}
	}
	data["sign"] = Md5String(param)
}

// 生成sign
func MakeSign(data map[string]string, secretKey string) {
	p := kvPairs{}
	// 剔除空值 和 sign
	for k, v := range data {
		if !(v == "" || k == "sign") {
			p = append(p, kvPair{k, v})
		}
	}
	p.Sort()
	str := p.Join()
	println(str)
	data["sign"] = sha256Sign(str, secretKey)

}

// 验证sign
func VerifySign(data map[string]string, secretKey string) error {
	p := kvPairs{}
	sign, ok := data["sign"]
	if !ok {
		return errors.New("sign not exist")
	}
	for k, v := range data {
		if !(v == "" || k == "sign") {
			p = append(p, kvPair{k, v})
		}
	}

	p.Sort()
	goSign := sha256Sign(p.Join(), secretKey)
	if sign == goSign {
		return nil
	} else {
		return errors.New("sign not same")
	}
}

func sha256Sign(str, key string) string {
	h := sha256.New()
	h.Write([]byte(str))
	h.Write([]byte(key))
	return fmt.Sprintf("%x", h.Sum(nil))
}

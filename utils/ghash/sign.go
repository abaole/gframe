package utils

import (
	"context"
	"crypto"
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"paymananger/internal/pkg/conf"
	pb "paymananger/message/proto"
	"sort"
	"strings"

	log "paymananger/pkg/logger"

	"google.golang.org/grpc"
)

const (
	EncryptMD5           = iota //md5加密
	EncryptSHA1withRsa          //SHA1withRsa
	EncryptSHA256withRsa        //SHA256withRsa
	EncryptMD5withRsa           //MD5withRsa
	SignAddr             = "127.0.0.1:10081"
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
	data["sign"] = MD5(p.Join())
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
	data["sign"] = MD5(param)
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
	//amount=1&bizType=16&merchantNo=112570436922904576&nonceStr=sadfasfasdf&orderNo=12313331231551&payKey=bi4v582693m1o31b0ptg&payType=N01&tranTime=123123&version=v1
	p.Sort()
	goSign := sha256Sign(p.Join(), secretKey)
	if sign == goSign {
		return nil
	} else {
		return errors.New("sign not same")
	}
}

// 验证sign
func CCBGrpcVerifySign(strSrc, strSign, pubKey string) error {
	// 连接服务
	conn, err := grpc.Dial(conf.GetConfig("").HTTP.CcbSign, grpc.WithInsecure())
	if err != nil {
		log.Panicf("建设银行回调验证服务连接失败：%v", err)
	}
	defer conn.Close()

	client := pb.NewCcbSignClient(conn)
	reqs := &pb.SignRequest{}
	reqs.Params = strSrc
	reqs.Sign = strSign
	reqs.Key = pubKey

	reply, err := client.CcbSignRsa(context.Background(), reqs)
	if err != nil {
		log.Errorf("签名GRPC请求签名错误")
		return err
	}
	if reply.Result == "Y" {
		return nil
	}
	log.Warnf("签名%s", reply.Result)
	return errors.New("N 签名失败")
}

// 验证sign
func CCBCallBackVerifySign(data, sign, gongyao string) error {

	tmp, _ := hex.DecodeString(gongyao)
	publicKey := string(tmp)
	publicKey = base64.StdEncoding.EncodeToString([]byte(publicKey))

	publicKey = "-----BEGIN PUBLIC KEY-----\n" + WordWrap(string(publicKey), 64) + "\n-----END PUBLIC KEY-----"

	var genSign string

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return errors.New("private key error!")
	}
	//asn1.Unmarshal(block.Bytes, &pk)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		fmt.Println("pub is of type RSA:", pub)
	case *dsa.PublicKey:
		fmt.Println("pub is of type DSA:", pub)
	case *ecdsa.PublicKey:
		fmt.Println("pub is of type ECDSA:", pub)
	default:
		panic("unknown type of public key")
	}

	rsaKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		fmt.Printf("got unexpected key type: %T", rsaKey)
	}

	hashMD5 := md5.New()
	hashMD5.Write([]byte(data))
	Digest := hashMD5.Sum(nil)

	tmm, _ := hex.DecodeString(sign)
	publiSign := string(tmm)
	publiSign = base64.StdEncoding.EncodeToString([]byte(publiSign))

	signature, _ := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	errRsa := rsa.VerifyPKCS1v15(rsaKey, crypto.MD5, Digest, signature)

	print(errRsa)

	if sign == genSign {
		return nil
	} else {
		return errors.New("sign not same")
	}
}

func md5Sign(str, key string) string {
	h := md5.New()
	h.Write([]byte(str))
	h.Write([]byte(key))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func sha256Sign(str, key string) string {
	h := sha256.New()
	h.Write([]byte(str))
	h.Write([]byte(key))
	return fmt.Sprintf("%x", h.Sum(nil))
}

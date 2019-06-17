package ghash

import "framework/utils"

/* 对字符串取hash值 */

// Md5String 获取字符串md5值
func Md5String(s string) string {
	return utils.Md5Byte([]byte(s))
}

// Sha1String 获取字符串sha1值
func Sha1String(s string) string {
	return utils.Sha1Byte([]byte(s))
}

// Sha256String 获取字符串sha256值
func Sha256String(s string) string {
	return utils.Sha256Byte([]byte(s))
}

// Sha512String 获取字符串sha512值
func Sha512String(s string) string {
	return utils.Sha512Byte([]byte(s))
}

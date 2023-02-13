package symmetric

import "encoding/base64"

// base64.URLEncoding.EncodeToString(rawStr string) 方法来对应用于URL中的base64编码进行了一些特殊处理，
// 也就是将 '+' 替换为 '-'，将 '/' 替换为 '_'符号。
const (
	base64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"
)

var coder = base64.NewEncoding(base64Table)

// Base64Encode 自定义密钥加密
func Base64Encode(src []byte) []byte { //编码
	return []byte(coder.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) { //解码
	return coder.DecodeString(string(src))
}

// Base64StdEncode 标准加解密
func Base64StdEncode(src []byte) string { //编码
	return base64.StdEncoding.EncodeToString(src)
}

func Base64StdDecode(src string) ([]byte, error) { //解码
	return base64.StdEncoding.DecodeString(src)
}

// Base64UrlEncode url加解密
func Base64UrlEncode(src []byte) string { //编码
	return base64.URLEncoding.EncodeToString(src)
}

func Base64UrlDecode(src string) ([]byte, error) { //解码
	return base64.URLEncoding.DecodeString(src)
}

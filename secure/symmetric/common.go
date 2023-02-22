package symmetric

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// 注：
// aes加密，分组长度128比特，密钥长度以32比特为单位在128比特-256比特之间选择，默认规格中只有128，192，256，
// go提供的接口中，密钥长度只能为16字节
// ecb/ctr:不需要初始化向量（go接口中的iv可以理解为随机数种子, iv的长度 == 明文分组的长度）
// cbc/cfb/ofb:需要初始化向量（des/3des->8字节，aes->16字节，加解密向量相同）
// 如果使用ecb/cbc分组模式需要对明文分组进行填充,cfb/ofb/ctr都不需要对分组明文填充

// 密文先分组，再加密
// des/3des按8字节分组
// 分组密码模式(ecb/cbc模式需要填充 )

// 1、创建des/3des/aes密码接口
// crypto/des   (des/3des)
// crypto/aes   (aes)
// NewCipher(key []byte)(cipher.Block,error)

// 2、如果使用cbc/ecb分组模式需要对密文分组进行填充,cfb/ofb/ctr都不需要对分组密文填充

// 3、创建分组模式接口对象
// cbc:NewCBCEncrypter(b block,iv []byte) BlockMode 加密
// cfb:NewCFBEncrypter(b Block, iv []byte) Stream 加密

func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// block, _ := des.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式 初始化向量默认为一个字节数组的block.BlockSize()长度
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式，初始化向量默认为一个字节数组的block.BlockSize()长度
	//blockMode := cipher.NewCFBDecrypter(block, k[:blockSize])
	//blockMode := cipher.NewOFB(block, k[:blockSize])
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	//blockMode.XORKeyStream(orig, crytedByte)
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

// 私钥生成
// openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQDcGsUIIAINHfRTdMmgGwLrjzfMNSrtgIf4EGsNaYwmC1GjF/bM
h0Mcm10oLhNrKNYCTTQVGGIxuc5heKd1gOzb7bdTnCDPPZ7oV7p1B9Pud+6zPaco
qDz2M24vHFWYY2FbIIJh8fHhKcfXNXOLovdVBE7Zy682X1+R1lRK8D+vmQIDAQAB
AoGAeWAZvz1HZExca5k/hpbeqV+0+VtobMgwMs96+U53BpO/VRzl8Cu3CpNyb7HY
64L9YQ+J5QgpPhqkgIO0dMu/0RIXsmhvr2gcxmKObcqT3JQ6S4rjHTln49I2sYTz
7JEH4TcplKjSjHyq5MhHfA+CV2/AB2BO6G8limu7SheXuvECQQDwOpZrZDeTOOBk
z1vercawd+J9ll/FZYttnrWYTI1sSF1sNfZ7dUXPyYPQFZ0LQ1bhZGmWBZ6a6wd9
R+PKlmJvAkEA6o32c/WEXxW2zeh18sOO4wqUiBYq3L3hFObhcsUAY8jfykQefW8q
yPuuL02jLIajFWd0itjvIrzWnVmoUuXydwJAXGLrvllIVkIlah+lATprkypH3Gyc
YFnxCTNkOzIVoXMjGp6WMFylgIfLPZdSUiaPnxby1FNM7987fh7Lp/m12QJAK9iL
2JNtwkSR3p305oOuAz0oFORn8MnB+KFMRaMT9pNHWk0vke0lB1sc7ZTKyvkEJW0o
eQgic9DvIYzwDUcU8wJAIkKROzuzLi9AvLnLUrSdI6998lmeYO9x7pwZPukz3era
zncjRK3pbVkv0KrKfczuJiRlZ7dUzVO0b6QJr8TRAA==
-----END RSA PRIVATE KEY-----
`)

// 公钥: 根据私钥生成
// openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDcGsUIIAINHfRTdMmgGwLrjzfM
NSrtgIf4EGsNaYwmC1GjF/bMh0Mcm10oLhNrKNYCTTQVGGIxuc5heKd1gOzb7bdT
nCDPPZ7oV7p1B9Pud+6zPacoqDz2M24vHFWYY2FbIIJh8fHhKcfXNXOLovdVBE7Z
y682X1+R1lRK8D+vmQIDAQAB
-----END PUBLIC KEY-----
`)

func RsaSign(origData string) (sign string) {

	return
}

// 加密
func RsaEncrypt(origData string) (codeStr string, err error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err = errors.New("public key error")
		return
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	codeByte := make([]byte, 0)
	codeByte, err = rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData))
	if err != nil {
		return
	}
	return base64.StdEncoding.EncodeToString(codeByte), nil
}

// 解密
func RsaDecrypt(ciphertext string) (codeStr string, err error) {
	cipherByte := make([]byte, 0)
	cipherByte, err = base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return
	}
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = errors.New("private key error!")
		return
	}
	//解析PKCS1格式的私钥
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	// 解密
	codeByte := make([]byte, 0)
	codeByte, err = rsa.DecryptPKCS1v15(rand.Reader, private, cipherByte)
	if err != nil {
		return
	}
	codeStr = string(codeByte)
	return
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	paddingText := bytes.Repeat([]byte{0}, padding) // 用0填充
	return append(ciphertext, paddingText...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

// PKCS5填充块的大小为8bytes(64位)
// PKCS7填充块的大小可以在1-255bytes之间。
// PKCS7是兼容PKCS5的，PKCS5相当于PKCS7的一个子集

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

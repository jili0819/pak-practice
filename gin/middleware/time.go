package middleware

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

func TimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		fmt.Println("start time: ", startTime.Format(time.DateTime))
		c.Next()
		fmt.Println("end time: ", time.Since(startTime))
		return
	}
}
func CheckAliMnsSi() gin.HandlerFunc {
	return func(c *gin.Context) {
		var StringToSign string
		method := c.Request.Method
		contentMd5 := c.GetHeader("Content-MD5")
		contentType := c.GetHeader("Content-Type")
		date := c.GetHeader("Date")
		header := c.Request.Header

		if method == "" || contentMd5 == "" || contentType == "" || date == "" || header == nil {
			c.Abort()
			return
		}
		StringToSign = strings.ToUpper(c.Request.Method) + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n"
		mnsHeaderKey := make([]string, 0)
		mnsHeaderMap := make(map[string]string)
		for k, _ := range header {
			if strings.HasPrefix(strings.ToLower(k), "x-mns-") {
				mnsHeaderKey = append(mnsHeaderKey, strings.ToLower(k))
				mnsHeaderMap[strings.ToLower(k)] = k
			}
		}
		sort.Strings(mnsHeaderKey)
		for _, mnsKey := range mnsHeaderKey {
			StringToSign += mnsKey + ":" + c.GetHeader(mnsHeaderMap[mnsKey]) + "\n"
		}
		StringToSign += c.Request.URL.RequestURI()
		// 获取证书的URL
		certURLBase64 := c.GetHeader("X-Mns-Signing-Cert-Url")
		if certURLBase64 == "" {
			c.Abort()
			return
		}
		certURLBytes, err := base64.StdEncoding.DecodeString(certURLBase64)
		if err != nil {
			c.Abort()
			return
		}

		certURL := string(certURLBytes)
		if !strings.HasPrefix(certURL, "https://mnstest.oss-cn-hangzhou.aliyuncs.com/") {
			c.Abort()
			return
		}

		// 根据URL获取证书，并从证书中获取公钥
		resp, err := http.Get(certURL)
		if err != nil {
			c.Abort()
			return
		}
		//goland:noinspection GoUnhandledErrorResult
		defer resp.Body.Close()
		//goland:noinspection GoDeprecation
		certData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.Abort()
			return
		}

		block, _ := pem.Decode(certData)
		if block == nil || block.Type != "CERTIFICATE" {
			c.Abort()
			return
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			c.Abort()
			return
		}

		pubKey, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			c.Abort()
			return
		}

		// 对Authorization字段做Base64解码
		signatureBase64 := c.GetHeader("Authorization")
		if !ok {
			c.Abort()
			return
		}

		signature, err := base64.StdEncoding.DecodeString(signatureBase64)
		if err != nil {
			fmt.Println("Failed to decode base64 signature:", err)
			c.Abort()
			return
		}

		// 认证
		hash := sha1.New()
		hash.Write([]byte(StringToSign))
		digest := hash.Sum(nil)
		err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA1, digest, signature)
		if err != nil {
			fmt.Println("Signature verification failed:", err)
			c.Abort()
			return
		}
		c.Next()
	}
}

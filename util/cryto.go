// package 工具类
package util

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/big"
	"strings"

	"github.com/forgoer/openssl"
)

var iv = []byte("0102030405060708")
var presetKey = []byte("0CoJUm6Qyw8W8jud")

// linuxapiKey linuxAPI加密固定key
var linuxapiKey = []byte("rFgB&h#%2?^eDg:Q")
var stdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// publicKey 公钥
var publicKey = []byte("-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB\n-----END PUBLIC KEY-----")
var eapiKey = []byte("e82ckenh8dichen8")

func aesEncryptCBC(buffer []byte, key []byte, ivv []byte) []byte {
	dst, _ := openssl.AesCBCEncrypt(buffer, key, ivv, openssl.PKCS7_PADDING)
	return dst
	// base64 解码
	//fmt.Println(base64.StdEncoding.EncodeToString(dst))

	// 解密
	//dst, _ = openssl.AesCBCDecrypt(dst, presetKey, iv, openssl.PKCS7_PADDING)
	//fmt.Println(string(dst)) // 123456
}

func aesEncryptECB(buffer []byte, key []byte) []byte {
	dst, _ := openssl.AesECBEncrypt(buffer, key, openssl.PKCS7_PADDING)
	return dst
	//fmt.Println(base64.StdEncoding.EncodeToString(dst))  // yXVUkR45PFz0UfpbDB8/ew==
	// hex 解码
	//fmt.Println(hex.EncodeToString(dst))

	//解密
	//dst, _ = openssl.AesECBDecrypt(dst, linuxapiKey, openssl.PKCS7_PADDING)
	//fmt.Println(string(dst)) // 123456
}

// NewLen16Rand 16 位随机数生成
//  @return []byte 随机字节
//  @return []byte 随机字节取反
func NewLen16Rand() ([]byte, []byte) {

	// randByte 随机字节
	randByte := make([]byte, 16)

	//随机字节取反
	randByteReverse := make([]byte, 16)
	for i := 0; i < 16; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(62))
		randByte[i] = stdChars[result.Int64()]
		randByteReverse[15-i] = stdChars[result.Int64()]
	}
	return randByte, randByteReverse
}

// aesEncrypt AES加密
//  @param buffer 要加密的内容
//  @param mod 模式 cbc/ecb
//  @param key 私钥
//  @param ivv 偏移量
//  @return []byte 密文
func aesEncrypt(buffer []byte, mod string, key []byte, ivv []byte) []byte {
	if mod == "cbc" {
		return aesEncryptCBC(buffer, key, ivv)
	} else if mod == "ecb" {
		return aesEncryptECB(buffer, key)
	}
	return nil
}

// rsaEncrypt RSA加密
//  @param buffer 加密的内容
//  @param key 公钥
//  @return []byte 密文
func rsaEncrypt(buffer []byte, key []byte) []byte {
	// buffers []byte 初始创建112 最大容量128
	buffers := make([]byte, 128-16, 128)
	buffers = append(buffers, buffer...) //合并切片
	block, _ := pem.Decode(key)
	if block == nil {
		return nil
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)

	// 加密 因为网易采用的是无padding加密故直接进行计算
	c := new(big.Int).SetBytes([]byte(buffers))
	//计算
	encryptedBytes := c.Exp(c, big.NewInt(int64(pub.E)), pub.N).Bytes()
	return encryptedBytes
}

// Weapi Weapi方式加密
//  @param data 数据
//  @return map 加密结果
func Weapi(data map[string]string) map[string]string {
	text, _ := json.Marshal(data)            //生成json格式
	secretKey, reSecretKey := NewLen16Rand() //生成长度为16的随机数
	weapiType := make(map[string]string, 2)  //生成长度为2的map
	weapiType["params"] = base64.StdEncoding.EncodeToString(
		aesEncrypt([]byte(base64.StdEncoding.EncodeToString(
			aesEncrypt(text, "cbc", presetKey, iv))), "cbc", reSecretKey, iv))
	weapiType["encSecKey"] = hex.EncodeToString(rsaEncrypt(secretKey, publicKey))
	return weapiType
}

// Linuxapi  LinuxAPI方式加密
//  @param data 要加密的数据
//  @return map linuxapiType
func Linuxapi(data map[string]interface{}) map[string]string {
	text, _ := json.Marshal(data) //生成json
	linuxapiType := make(map[string]string, 1)
	//使用aes ecb 方式加密 生成hex
	linuxapiType["eparams"] = strings.ToUpper(hex.EncodeToString(aesEncrypt(text, "ecb", linuxapiKey, nil)))
	return linuxapiType
}

func Eapi(url string, data map[string]interface{}) map[string]string {
	textByte, _ := json.Marshal(data)
	fmt.Println(string(textByte))
	message := "nobody" + url + "use" + string(textByte) + "md5forencrypt"
	h := md5.New()
	h.Write([]byte(message))
	digest := hex.EncodeToString(h.Sum(nil))
	dd := url + "-36cd479b6b5-" + string(textByte) + "-36cd479b6b5-" + digest
	eapiType := make(map[string]string, 1)
	eapiType["params"] = strings.ToUpper(hex.EncodeToString(aesEncrypt([]byte(dd), "ecb", eapiKey, nil)))
	return eapiType
}

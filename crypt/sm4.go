package crypt

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm4"
	"gvadmin_v3/core/config"
)

// SM4加密
func SM4Encrypt(data string) (result string, err error) {
	//字符串转byte切片
	plainText := []byte(data)
	SM4Key := config.Instance().Crypt.SM4Key //建议从配置文件中读取秘钥，进行统一管理, 16位字符串, 128bit
	SM4Iv := config.Instance().Crypt.SM4Iv   //iv建议是随机的，进一步保证加密的安全性，将iv的值和加密后的数据一起返回给外部, 16位字符串, 128bit
	iv := []byte(SM4Iv)
	key := []byte(SM4Key)
	//实例化sm4加密对象
	block, err := sm4.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//明文数据填充
	paddingData := paddingLastGroup(plainText, block.BlockSize())
	//声明SM4的加密工作模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//为填充后的数据进行加密处理
	cipherText := make([]byte, len(paddingData))
	//使用CryptBlocks这个核心方法，将paddingData进行加密处理，将加密处理后的值赋值到cipherText中
	blockMode.CryptBlocks(cipherText, paddingData)
	//加密结果使用hex转成字符串，方便外部调用
	cipherString := hex.EncodeToString(cipherText)
	return cipherString, nil
}

// SM4解密 传入string 输出string
func SM4Decrypt(data string) (res string, err error) {
	//秘钥
	SM4Key := config.Instance().Crypt.SM4Key
	//iv是Initialization Vector，初始向量，
	SM4Iv := config.Instance().Crypt.SM4Iv
	iv := []byte(SM4Iv)
	key := []byte(SM4Key)
	block, err := sm4.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//使用hex解码
	decodeString, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}
	//CBC模式 优点：具有较好的安全性，能够隐藏明文的模式和重复性。 缺点：加密过程是串行的，不适合并行处理。
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//下文有详解这段代码的含义
	blockMode.CryptBlocks(decodeString, decodeString)
	//去掉明文后面的填充数据
	plainText := unPaddingLastGroup(decodeString)
	//直接返回字符串类型，方便外部调用
	return string(plainText), nil
}

// 明文数据填充
func paddingLastGroup(plainText []byte, blockSize int) []byte {
	//1.计算最后一个分组中明文后需要填充的字节数
	padNum := blockSize - len(plainText)%blockSize
	//2.将字节数转换为byte类型
	char := []byte{byte(padNum)}
	//3.创建切片并初始化
	newPlain := bytes.Repeat(char, padNum)
	//4.将填充数据追加到原始数据后
	newText := append(plainText, newPlain...)
	return newText
}

// 去掉明文后面的填充数据
func unPaddingLastGroup(plainText []byte) []byte {
	//1.拿到切片中的最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1]
	//2.将最后一个数据转换为整数
	number := int(lastChar)
	return plainText[:length-number]
}

package crypt

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm2"
	"gvadmin_core/config"
	"math/big"
)

var chipMode = config.Instance().Crypt.ChipMode

// GetSm2Keys 用于生成sm2密钥对
func GetSm2Keys() (string, string, error) {
	privKey, err := sm2.GenerateKey(rand.Reader) // 生成密钥对
	if err != nil {
		return "", "", err
	}
	pubKey := &privKey.PublicKey
	return publicKeyToString(pubKey), privateKeyToString(privKey), nil
}

// privateKeyToString 私钥sm2.PrivateKey 转字符串(与java中org.bouncycastle.crypto生成的公私钥完全互通使用)
func privateKeyToString(privKey *sm2.PrivateKey) string {
	return hex.EncodeToString(privKey.D.Bytes())
}

// publicKeyToString 公钥sm2.PublicKey转字符串(与java中org.bouncycastle.crypto生成的公私钥完全互通使用)
func publicKeyToString(publicKey *sm2.PublicKey) string {
	xBytes := publicKey.X.Bytes()
	yBytes := publicKey.Y.Bytes()

	// 确保坐标字节切片长度相同
	byteLen := len(xBytes)
	if len(yBytes) > byteLen {
		byteLen = len(yBytes)
	}

	// 为坐标补齐前导零
	xBytes = append(make([]byte, byteLen-len(xBytes)), xBytes...)
	yBytes = append(make([]byte, byteLen-len(yBytes)), yBytes...)

	// 添加 "04" 前缀
	publicKeyBytes := append([]byte{0x04}, append(xBytes, yBytes...)...)
	return hex.EncodeToString(publicKeyBytes)
}

// stringToPublicKey 公钥字符串还原为 sm2.PublicKey 对象(与java中org.bouncycastle.crypto生成的公私钥完全互通使用)
func stringToPublicKey(publicKeyStr string) (*sm2.PublicKey, error) {
	publicKeyBytes, err := hex.DecodeString(publicKeyStr)
	if err != nil {
		return nil, err
	}

	// 提取 x 和 y 坐标字节切片
	curve := sm2.P256Sm2().Params()
	byteLen := (curve.BitSize + 7) / 8
	xBytes := publicKeyBytes[1 : byteLen+1]
	yBytes := publicKeyBytes[byteLen+1 : 2*byteLen+1]

	// 将字节切片转换为大整数
	x := new(big.Int).SetBytes(xBytes)
	y := new(big.Int).SetBytes(yBytes)

	// 创建 sm2.PublicKey 对象
	publicKey := &sm2.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}

	return publicKey, nil
}

// stringToPrivateKey 私钥还原为 sm2.PrivateKey对象(与java中org.bouncycastle.crypto生成的公私钥完全互通使用)
func stringToPrivateKey(privateKeyStr string, publicKey *sm2.PublicKey) (*sm2.PrivateKey, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyStr)
	if err != nil {
		return nil, err
	}

	// 将字节切片转换为大整数
	d := new(big.Int).SetBytes(privateKeyBytes)

	// 创建 sm2.PrivateKey 对象
	privateKey := &sm2.PrivateKey{
		PublicKey: *publicKey,
		D:         d,
	}
	return privateKey, nil
}

// SM2Encrypt 公钥字符串, text 待加密明文字符串
func SM2Encrypt(publicKeyStr, text string) (string, error) {
	publicKey, err := stringToPublicKey(publicKeyStr)
	if err != nil {
		return "", err
	}
	encryptStr, _ := sm2.Encrypt(publicKey, []byte(text), rand.Reader, chipMode)
	encodeToString := hex.EncodeToString(encryptStr)
	//fmt.Println("加密后的字符串：", encodeToString)
	return encodeToString, nil
}

// SM2Decrypt 私钥字符串, cipherText 待解密字符串
func SM2Decrypt(publicKeyStr, privateKeyStr, cipherText string) (string, error) {
	publicKeyObj, err := stringToPublicKey(publicKeyStr)
	if err != nil {
		return "", err
	}
	privateKeyObj, err := stringToPrivateKey(privateKeyStr, publicKeyObj)
	if err != nil {
		return "", err
	}
	decodeString, err := hex.DecodeString(cipherText)
	decrypt, err := sm2.Decrypt(privateKeyObj, decodeString, chipMode)
	if err != nil {
		return "", err
	}
	resultStr := string(decrypt)
	//fmt.Println("解密后的字符串：", resultStr)
	return resultStr, nil
}

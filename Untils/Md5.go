package Untils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func MD5IsOK() string {
	data := []byte("ashjk521007712123432--99") // 要加密的数据
	// 创建MD5哈希
	hash := md5.New()
	hash.Write(data)
	hashInBytes := hash.Sum(nil)

	// 将哈希转为16进制格式
	encodedData := hex.EncodeToString(hashInBytes)

	fmt.Println("MD5 加密结果:", encodedData)
	return encodedData
}

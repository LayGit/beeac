package utils

import (
  "bytes"
  "crypto/aes"
  "crypto/cipher"
  "crypto/md5"
  "crypto/sha1"
  "encoding/base64"
  "encoding/hex"
  "strings"
  "io"
)

// Base64 加密
func Base64Encode(source []byte) string {
  bytearr := base64.StdEncoding.EncodeToString(source)
  retval := strings.Replace(string(bytearr), "/", "_", -1)
  retval = strings.Replace(retval, "+", "-", -1)
  retval = strings.Replace(retval, "=", "", -1)
  return retval
}

// Base64 解密
func Base64Decode(data string) ([]byte, error) {
  var missing = (4 - len(data)%4) % 4
  data += strings.Repeat("=", missing)
  return base64.URLEncoding.DecodeString(data)
}

// AES 加密
func AesEncrypt(src, key string) ([]byte, error) {
  bKey := []byte(key)
  origData := []byte(src)
  block, err := aes.NewCipher(bKey)
  if err != nil {
    return nil, err
  }
  blockSize := block.BlockSize()
  origData = PKCS5Padding(origData, blockSize)
  // origData = ZeroPadding(origData, block.BlockSize())
  blockMode := cipher.NewCBCEncrypter(block, bKey[:blockSize])
  crypted := make([]byte, len(origData))
  // 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
  // crypted := origData
  blockMode.CryptBlocks(crypted, origData)
  return crypted, nil
}

func AesDecrypt(crypted []byte, key string) (string, error) {
  sKey := []byte(key)
  block, err := aes.NewCipher(sKey)
  if err != nil {
    return "", err
  }
  blockSize := block.BlockSize()
  blockMode := cipher.NewCBCDecrypter(block, sKey[:blockSize])
  origData := make([]byte, len(crypted))
  // origData := crypted
  blockMode.CryptBlocks(origData, crypted)
  origData = PKCS5UnPadding(origData)
  // origData = ZeroUnPadding(origData)
  return string(origData), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
  padding := blockSize - len(ciphertext)%blockSize
  padtext := bytes.Repeat([]byte{byte(padding)}, padding)
  return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
  length := len(origData)
  // 去掉最后一个字节 unpadding 次
  unpadding := int(origData[length-1])
  return origData[:(length - unpadding)]
}

func MD5Encrypt(src string) string {
  h := md5.New()
  io.WriteString(h, src)
  return hex.EncodeToString(h.Sum(nil))
}

func SHA1Encrypt(src string) string {
  h := sha1.New()
  h.Write([]byte(src))
  return hex.EncodeToString(h.Sum(nil))
}

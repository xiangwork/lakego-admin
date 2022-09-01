package encrypt

import (
    "errors"
    "crypto/rand"
    "crypto/cipher"
    "encoding/asn1"
)

// OFB 模式加密参数
type ofbParams []byte

// OFB 模式加密
type CipherOFB struct {
    cipherFunc func(key []byte) (cipher.Block, error)
    keySize    int
    blockSize  int
    identifier asn1.ObjectIdentifier
}

// 值大小
func (this CipherOFB) KeySize() int {
    return this.keySize
}

// oid
func (this CipherOFB) OID() asn1.ObjectIdentifier {
    return this.identifier
}

// 加密
func (this CipherOFB) Encrypt(key, plaintext []byte) ([]byte, []byte, error) {
    // 加密数据补码
    plaintext = pkcs7Padding(plaintext, this.blockSize)

    // 随机生成 iv
    iv := make(ofbParams, this.blockSize)
    if _, err := rand.Read(iv); err != nil {
        return nil, nil, errors.New("pkcs8:" + err.Error() + " failed to generate IV")
    }

    block, err := this.cipherFunc(key)
    if err != nil {
        return nil, nil, errors.New("pkcs8:" + err.Error() + " failed to create cipher")
    }

    // 需要保存的加密数据
    encrypted := make([]byte, len(plaintext))

    enc := cipher.NewOFB(block, iv)
    enc.XORKeyStream(encrypted, plaintext)

    // 编码 iv
    paramBytes, err := asn1.Marshal(iv)
    if err != nil {
        return nil, nil, err
    }

    return encrypted, paramBytes, nil
}

// 解密
func (this CipherOFB) Decrypt(key, params, ciphertext []byte) ([]byte, error) {
    // 解析出 iv
    var iv ofbParams
    if _, err := asn1.Unmarshal(params, &iv); err != nil {
        return nil, errors.New("pkcs8: invalid iv parameters")
    }

    plaintext := make([]byte, len(ciphertext))

    block, err := this.cipherFunc(key)
    if err != nil {
        return nil, err
    }

    // 判断数据是否为填充数据
    blockSize := block.BlockSize()
    dlen := len(ciphertext)
    if dlen == 0 || dlen%blockSize != 0 {
        return nil, errors.New("pkcs8: invalid padding")
    }

    mode := cipher.NewOFB(block, iv)
    mode.XORKeyStream(plaintext, ciphertext)

    // 解析加密数据
    plaintext = pkcs7UnPadding(plaintext)

    return plaintext, nil
}
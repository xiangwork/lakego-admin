package cryptobin

// 私钥/公钥
func (this EdDSA) ToKeyBytes() []byte {
    return this.keyData
}

// 私钥/公钥
func (this EdDSA) ToKeyString() string {
    return string(this.keyData)
}

// ==========

// 输出字节
func (this EdDSA) ToBytes() []byte {
    return this.paredData
}

// 输出字符
func (this EdDSA) ToString() string {
    return string(this.paredData)
}

// 输出Base64
func (this EdDSA) ToBase64String() string {
    return NewEncoding().Base64Encode(this.paredData)
}

// 输出Hex
func (this EdDSA) ToHexString() string {
    return NewEncoding().HexEncode(this.paredData)
}

// ==========

// 验证结果
func (this EdDSA) ToVeryed() bool {
    return this.veryed
}

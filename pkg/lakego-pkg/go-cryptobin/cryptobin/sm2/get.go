package sm2

import (
    "crypto/elliptic"

    "github.com/deatil/go-cryptobin/tool"
    "github.com/deatil/go-cryptobin/gm/sm2"
)

// get PrivateKey
func (this SM2) GetPrivateKey() *sm2.PrivateKey {
    return this.privateKey
}

// get PrivateKey Curve
func (this SM2) GetPrivateKeyCurve() elliptic.Curve {
    return this.privateKey.Curve
}

// get PrivateKey D hex string
func (this SM2) GetPrivateKeyDString() string {
    data := this.privateKey.D

    dataHex := tool.HexEncode(data.Bytes())

    return dataHex
}

// get PrivateKey data hex string
func (this SM2) GetPrivateKeyString() string {
    return this.GetPrivateKeyDString()
}

// get PublicKey
func (this SM2) GetPublicKey() *sm2.PublicKey {
    return this.publicKey
}

// get PublicKey Curve
func (this SM2) GetPublicKeyCurve() elliptic.Curve {
    return this.publicKey.Curve
}

// get PublicKey X hex string
func (this SM2) GetPublicKeyXString() string {
    data := this.publicKey.X

    dataHex := tool.HexEncode(data.Bytes())

    return dataHex
}

// get PublicKey Y hex string
func (this SM2) GetPublicKeyYString() string {
    data := this.publicKey.Y

    dataHex := tool.HexEncode(data.Bytes())

    return dataHex
}

// get PublicKey X and Y Hex string
func (this SM2) GetPublicKeyXYString() string {
    dataHex := this.GetPublicKeyXString() + this.GetPublicKeyYString()

    return dataHex
}

// get PublicKey Uncompress Hex string
func (this SM2) GetPublicKeyUncompressString() string {
    dataHex := "04" + this.GetPublicKeyXString() + this.GetPublicKeyYString()

    return dataHex
}

// get PublicKey Compress Hex string
func (this SM2) GetPublicKeyCompressString() string {
    data := sm2.Compress(this.publicKey)

    dataHex := tool.HexEncode(data)

    return dataHex
}

// get key Data
func (this SM2) GetKeyData() []byte {
    return this.keyData
}

// get mode
func (this SM2) GetMode() sm2.Mode {
    return this.mode
}

// get data
func (this SM2) GetData() []byte {
    return this.data
}

// get parsedData
func (this SM2) GetParsedData() []byte {
    return this.parsedData
}

// get uid
func (this SM2) GetUID() []byte {
    return this.uid
}

// get verify data
func (this SM2) GetVerify() bool {
    return this.verify
}

// get errors
func (this SM2) GetErrors() []error {
    return this.Errors
}

package cryptobin

import (
    "errors"
    "crypto/dsa"
    "crypto/x509"
    "encoding/pem"
)

var (
    ErrNotDSAPrivateKey = errors.New("key is not a valid DSA private key")
    ErrNotDSAPublicKey  = errors.New("key is not a valid DSA public key")
)

// 解析私钥
func (this DSA) ParsePrivateKeyFromPEM(key []byte) (*dsa.PrivateKey, error) {
    var err error

    // Parse PEM block
    var block *pem.Block
    if block, _ = pem.Decode(key); block == nil {
        return nil, ErrKeyMustBePEMEncoded
    }

    // Parse the key
    var parsedKey any
    if parsedKey, err = this.ParsePrivateKey(block.Bytes); err != nil {
        return nil, err
    }

    var pkey *dsa.PrivateKey
    var ok bool
    if pkey, ok = parsedKey.(*dsa.PrivateKey); !ok {
        return nil, ErrNotDSAPrivateKey
    }

    return pkey, nil
}

// 解析私钥带密码
func (this DSA) ParsePrivateKeyFromPEMWithPassword(key []byte, password string) (*dsa.PrivateKey, error) {
    var err error

    // Parse PEM block
    var block *pem.Block
    if block, _ = pem.Decode(key); block == nil {
        return nil, ErrKeyMustBePEMEncoded
    }

    var blockDecrypted []byte
    if blockDecrypted, err = x509.DecryptPEMBlock(block, []byte(password)); err != nil {
        return nil, err
    }

    // Parse the key
    var parsedKey any
    if parsedKey, err = this.ParsePrivateKey(blockDecrypted); err != nil {
        return nil, err
    }

    var pkey *dsa.PrivateKey
    var ok bool
    if pkey, ok = parsedKey.(*dsa.PrivateKey); !ok {
        return nil, ErrNotDSAPrivateKey
    }

    return pkey, nil
}

// 解析公钥
func (this DSA) ParsePublicKeyFromPEM(key []byte) (*dsa.PublicKey, error) {
    var err error

    // Parse PEM block
    var block *pem.Block
    if block, _ = pem.Decode(key); block == nil {
        return nil, ErrKeyMustBePEMEncoded
    }

    // Parse the key
    var parsedKey any
    if parsedKey, err = this.ParsePublicKey(block.Bytes); err != nil {
        return nil, err
    }

    var pkey *dsa.PublicKey
    var ok bool
    if pkey, ok = parsedKey.(*dsa.PublicKey); !ok {
        return nil, ErrNotDSAPublicKey
    }

    return pkey, nil
}
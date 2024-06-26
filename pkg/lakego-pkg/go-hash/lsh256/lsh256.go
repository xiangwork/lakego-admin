package lsh256

import (
    "hash"
)

const (
    // The size of a LSH-256 checksum in bytes.
    Size = 32

    // The size of a LSH-224 checksum in bytes.
    Size224 = 28

    // The blocksize of LSH-256 and LSH-224 in bytes.
    BlockSize = 128
)

// New returns a new hash.Hash computing the LSH-256 checksum.
func New() hash.Hash {
    return newDigest(Size)
}

// New224 returns a new hash.Hash computing the LSH-224 checksum.
func New224() hash.Hash {
    return newDigest(Size224)
}

// Sum256 returns the LSH-256 checksum of the data.
func Sum256(data []byte) (sum256 [Size]byte) {
    b := New()
    b.Write(data)
    sum := b.Sum(nil)

    copy(sum256[:], sum)
    return
}

// Sum224 returns the LSH-224 checksum of the data.
func Sum224(data []byte) (sum224 [Size224]byte) {
    b := New224()
    b.Write(data)
    sum := b.Sum(nil)

    copy(sum224[:], sum)
    return
}

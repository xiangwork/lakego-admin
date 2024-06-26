package skein512

import (
    "io"
    "errors"
)

// Args can be used to configure hash function for different purposes.
// All fields are optional: if a field is nil, it will not be used.
type Args struct {
    // Key is a secret key for MAC, KDF, or stream cipher
    Key []byte
    // Person is a personalization string
    Person []byte
    // PublicKey is a public key for signature hashing
    PublicKey []byte
    // KeyId is a key identifier for KDF
    KeyId []byte
    // Nonce for stream cipher or randomized hashing
    Nonce []byte
    // NoMsg indicates whether message input is used by the function.
    NoMsg bool
}

const (
    // hash size
    Size224 = 28
    Size256 = 32
    Size384 = 48
    Size512 = 64
)

// BlockSize is the block size of Skein-512 in bytes.
const BlockSize = 64

// Hash represents a state of Skein hash function.
// It implements hash.Hash interface.
type Hash struct {
    k [8]uint64 // chain value
    t [2]uint64 // tweak

    x  [64]byte // buffer
    nx int      // number of bytes in buffer

    outLen uint64 // output length in bytes
    noMsg  bool   // true if message block argument should not be used

    ik [8]uint64 // copy of initial chain value
}

// Reset resets hash to its state after initialization.
// If hash was initialized with arguments, such as key,
// these arguments are preserved.
func (h *Hash) Reset() {
    // Restore initial chain value.
    h.k = h.ik
    // Reset buffer.
    h.nx = 0
    // Init tweak to first message block.
    h.t[0] = 0
    h.t[1] = messageArg<<56 | firstBlockFlag
}

// Size returns the number of bytes Sum will return.
// If the hash was created with output size greater than the maximum
// size of int, the result is undefined.
func (h *Hash) Size() int {
    return int(h.outLen)
}

// BlockSize returns the hash's underlying block size.
func (h *Hash) BlockSize() int {
    return BlockSize
}

func (h *Hash) hashLastBlock() {
    // Pad buffer with zeros.
    for i := h.nx; i < len(h.x); i++ {
        h.x[i] = 0
    }
    // Set last block flag.
    h.t[1] |= lastBlockFlag
    // Process last block.
    h.hashBlock(h.x[:], uint64(h.nx))
    h.nx = 0
}

func (h *Hash) outputBlock(dst *[64]byte, counter uint64) {
    var u [8]uint64
    u[0] = counter
    block(&h.k, &outTweak, &u, &u)
    for i, v := range u {
        dst[i*8+0] = byte(v)
        dst[i*8+1] = byte(v >> 8)
        dst[i*8+2] = byte(v >> 16)
        dst[i*8+3] = byte(v >> 24)
        dst[i*8+4] = byte(v >> 32)
        dst[i*8+5] = byte(v >> 40)
        dst[i*8+6] = byte(v >> 48)
        dst[i*8+7] = byte(v >> 56)
    }
}

func (h *Hash) appendOutput(length uint64) []byte {
    var b [64]byte
    var counter uint64

    var out []byte

    for length > 0 {
        h.outputBlock(&b, counter)
        counter++ // increment counter
        if length < 64 {
            out = append(out, b[:length]...)
            break
        }

        out = append(out, b[:]...)
        length -= 64
    }

    return out
}

func (h *Hash) update(b []byte) {
    left := 64 - h.nx
    if len(b) > left {
        // Process leftovers.
        copy(h.x[h.nx:], b[:left])
        b = b[left:]
        h.hashBlock(h.x[:], 64)
        h.nx = 0
    }

    // Process full blocks except for the last one.
    for len(b) > 64 {
        h.hashBlock(b, 64)
        b = b[64:]
    }

    // Save leftovers.
    h.nx += copy(h.x[h.nx:], b)
}

// Write adds more data to the running hash.
// It never returns an error.
func (h *Hash) Write(b []byte) (n int, err error) {
    if h.noMsg {
        return 0, errors.New("Skein: can't write to a function configured with NoMsg")
    }
    h.update(b)
    return len(b), nil
}

// Sum appends the current hash to in and returns the resulting slice.
// It does not change the underlying hash state.
func (h *Hash) Sum(p []byte) []byte {
    // Make a copy of d so that caller can keep writing and summing.
    d0 := *h
    hash := d0.checkSum()
    return append(p, hash...)
}

func (h *Hash) checkSum() (hash []byte) {
    if !h.noMsg {
        // Finalize message.
        h.hashLastBlock()
    }

    return h.appendOutput(h.outLen)
}

// OutputReader returns an io.Reader that can be used to read
// arbitrary-length output of the hash.
// Reading from it doesn't change the underlying hash state.
func (h *Hash) OutputReader() io.Reader {
    return newOutputReader(h)
}

func (h *Hash) hashBlock(b []byte, unpaddedLen uint64) {
    var u [8]uint64

    // Update block counter.
    h.t[0] += unpaddedLen

    u[0] = uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
        uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
    u[1] = uint64(b[8]) | uint64(b[9])<<8 | uint64(b[10])<<16 | uint64(b[11])<<24 |
        uint64(b[12])<<32 | uint64(b[13])<<40 | uint64(b[14])<<48 | uint64(b[15])<<56
    u[2] = uint64(b[16]) | uint64(b[17])<<8 | uint64(b[18])<<16 | uint64(b[19])<<24 |
        uint64(b[20])<<32 | uint64(b[21])<<40 | uint64(b[22])<<48 | uint64(b[23])<<56
    u[3] = uint64(b[24]) | uint64(b[25])<<8 | uint64(b[26])<<16 | uint64(b[27])<<24 |
        uint64(b[28])<<32 | uint64(b[29])<<40 | uint64(b[30])<<48 | uint64(b[31])<<56
    u[4] = uint64(b[32]) | uint64(b[33])<<8 | uint64(b[34])<<16 | uint64(b[35])<<24 |
        uint64(b[36])<<32 | uint64(b[37])<<40 | uint64(b[38])<<48 | uint64(b[39])<<56
    u[5] = uint64(b[40]) | uint64(b[41])<<8 | uint64(b[42])<<16 | uint64(b[43])<<24 |
        uint64(b[44])<<32 | uint64(b[45])<<40 | uint64(b[46])<<48 | uint64(b[47])<<56
    u[6] = uint64(b[48]) | uint64(b[49])<<8 | uint64(b[50])<<16 | uint64(b[51])<<24 |
        uint64(b[52])<<32 | uint64(b[53])<<40 | uint64(b[54])<<48 | uint64(b[55])<<56
    u[7] = uint64(b[56]) | uint64(b[57])<<8 | uint64(b[58])<<16 | uint64(b[59])<<24 |
        uint64(b[60])<<32 | uint64(b[61])<<40 | uint64(b[62])<<48 | uint64(b[63])<<56

    block(&h.k, &h.t, &h.k, &u)

    // Clear first block flag.
    h.t[1] &^= firstBlockFlag
}

// outputReader implements io.Reader and cipher.Stream interfaces.
// It is used for reading arbitrary-length output of Skein.
type outputReader struct {
    Hash
    counter uint64
}

// newOutputReader returns a new outputReader initialized with
// a copy of the given hash.
func newOutputReader(h *Hash) *outputReader {
    // Initialize with the copy of h.
    r := &outputReader{Hash: *h}
    if !r.noMsg {
        // Finalize message.
        r.hashLastBlock()
    }
    // Set buffer position to end.
    r.nx = BlockSize
    return r
}

// nextBlock puts the next hash output block into the internal buffer.
func (r *outputReader) nextBlock() {
    r.outputBlock(&r.x, r.counter)
    r.counter++ // increment counter
    r.nx = 0
}

// Read puts the next len(p) bytes of hash output into p.
// It never returns an error.
func (r *outputReader) Read(p []byte) (n int, err error) {
    n = len(p)
    left := BlockSize - r.nx

    if len(p) < left {
        r.nx += copy(p, r.x[r.nx:r.nx+len(p)])
        return
    }

    copy(p, r.x[r.nx:])
    p = p[left:]
    r.nextBlock()

    for len(p) >= BlockSize {
        copy(p, r.x[:])
        p = p[BlockSize:]
        r.nextBlock()
    }
    if len(p) > 0 {
        r.nx += copy(p, r.x[:len(p)])
    }
    return
}

// XORKeyStream XORs each byte in the given slice with the next byte from the
// hash output. Dst and src may point to the same memory.
func (r *outputReader) XORKeyStream(dst, src []byte) {
    left := BlockSize - r.nx

    if len(src) < left {
        for i, v := range src {
            dst[i] = v ^ r.x[r.nx]
            r.nx++
        }
        return
    }

    for i, b := range r.x[r.nx:] {
        dst[i] = src[i] ^ b
    }
    dst = dst[left:]
    src = src[left:]
    r.nextBlock()

    for len(src) >= BlockSize {
        for i, v := range src[:BlockSize] {
            dst[i] = v ^ r.x[i]
        }
        dst = dst[BlockSize:]
        src = src[BlockSize:]
        r.nextBlock()
    }
    if len(src) > 0 {
        for i, v := range src {
            dst[i] = v ^ r.x[i]
            r.nx++
        }
    }
}

// addArg adds Skein argument into the hash state.
func (h *Hash) addArg(argType uint64, arg []byte) {
    h.t[0] = 0
    h.t[1] = argType<<56 | firstBlockFlag
    h.update(arg)
    h.hashLastBlock()
}

// addConfig adds configuration block into the hash state.
func (h *Hash) addConfig(outBits uint64) {
    var c [32]byte
    copy(c[:], schemaId)
    c[8] = byte(outBits)
    c[9] = byte(outBits >> 8)
    c[10] = byte(outBits >> 16)
    c[11] = byte(outBits >> 24)
    c[12] = byte(outBits >> 32)
    c[13] = byte(outBits >> 40)
    c[14] = byte(outBits >> 48)
    c[15] = byte(outBits >> 56)
    h.addArg(configArg, c[:])
}

// New returns a new skein.Hash configured with the given arguments. The final
// output length of hash function in bytes is outLen (for example, 64 when
// calculating 512-bit hash). Configuration arguments may be nil.
func New(outLen uint64, args *Args) *Hash {
    h := new(Hash)
    h.outLen = outLen

    if args != nil && args.Key != nil {
        // Key argument comes before configuration.
        h.addArg(keyArg, args.Key)
        // Configuration.
        h.addConfig(outLen * 8)
    } else {
        // Configuration without key.
        // Try using precomputed values for common sizes.
        switch outLen {
            case 224 / 8:
                h.k = iv224
            case 256 / 8:
                h.k = iv256
            case 384 / 8:
                h.k = iv384
            case 512 / 8:
                h.k = iv512
            default:
                h.addConfig(outLen * 8)
        }
    }

    // Other arguments, in specified order.
    if args != nil {
        h.noMsg = args.NoMsg

        if args.Person != nil {
            h.addArg(personArg, args.Person)
        }
        if args.PublicKey != nil {
            h.addArg(publicKeyArg, args.PublicKey)
        }
        if args.KeyId != nil {
            h.addArg(keyIDArg, args.KeyId)
        }
        if args.Nonce != nil {
            h.addArg(nonceArg, args.Nonce)
        }
    }

    // Init tweak to first message block.
    h.t[0] = 0
    h.t[1] = messageArg<<56 | firstBlockFlag

    // Save a copy of initial chain value for Reset.
    h.ik = h.k
    return h
}

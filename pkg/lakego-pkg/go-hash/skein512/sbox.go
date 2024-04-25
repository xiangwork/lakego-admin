package skein512

// Argument types (in the order they must be used).
const (
    keyArg       uint64 = 0
    configArg    uint64 = 4
    personArg    uint64 = 8
    publicKeyArg uint64 = 12
    keyIDArg     uint64 = 16
    nonceArg     uint64 = 20
    messageArg   uint64 = 48
    outputArg    uint64 = 63
)

const (
    firstBlockFlag uint64 = 1 << 62
    lastBlockFlag  uint64 = 1 << 63
)

var schemaId = []byte{'S', 'H', 'A', '3', 1, 0, 0, 0}
var outTweak = [2]uint64{8, outputArg<<56 | firstBlockFlag | lastBlockFlag}

// Precomputed initial values of state after configuration for unkeyed hashing.
var iv224 = [8]uint64{
    0xCCD0616248677224, 0xCBA65CF3A92339EF, 0x8CCD69D652FF4B64, 0x398AED7B3AB890B4,
    0x0F59D1B1457D2BD0, 0x6776FE6575D4EB3D, 0x99FBC70E997413E9, 0x9E2CFCCFE1C41EF7,
}

var iv256 = [8]uint64{
    0xCCD044A12FDB3E13, 0xE83590301A79A9EB, 0x55AEA0614F816E6F, 0x2A2767A4AE9B94DB,
    0xEC06025E74DD7683, 0xE7A436CDC4746251, 0xC36FBAF9393AD185, 0x3EEDBA1833EDFC13,
}

var iv384 = [8]uint64{
    0xA3F6C6BF3A75EF5F, 0xB0FEF9CCFD84FAA4, 0x9D77DD663D770CFE, 0xD798CBF3B468FDDA,
    0x1BC4A6668A0E4465, 0x7ED7D434E5807407, 0x548FC1ACD4EC44D6, 0x266E17546AA18FF8,
}

var iv512 = [8]uint64{
    0x4903ADFF749C51CE, 0x0D95DE399746DF03, 0x8FD1934127C79BCE, 0x9A255629FF352CB1,
    0x5DB62599DF6CA7B0, 0xEABE394CA9D5C3F4, 0x991112C71A75B523, 0xAE18A40B660FCC33,
}

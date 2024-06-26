package xxhash64_test

import (
    "fmt"
    "testing"
    "encoding/binary"

    "github.com/deatil/go-hash/xxhash/xxhash64"
)

type test struct {
    sum             uint64
    data, printable string
}

var testdata = []test{
    {0xef46db3751d8e999, "", ""},
    {0xd24ec4f1a98c6e5b, "a", ""},
    {0x65f708ca92d04a61, "ab", ""},
    {0x44bc2cf5ad770999, "abc", ""},
    {0xde0327b0d25d92cc, "abcd", ""},
    {0x07e3670c0c8dc7eb, "abcde", ""},
    {0xfa8afd82c423144d, "abcdef", ""},
    {0x1860940e2902822d, "abcdefg", ""},
    {0x3ad351775b4634b7, "abcdefgh", ""},
    {0x27f1a34fdbb95e13, "abcdefghi", ""},
    {0xd6287a1de5498bb2, "abcdefghij", ""},
    {0xbf2cd639b4143b80, "abcdefghijklmnopqrstuvwxyz012345", ""},
    {0x64f23ecf1609b766, "abcdefghijklmnopqrstuvwxyz0123456789", ""},
    {0xc5a8b11443765630, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", ""},

    {0x1c330fb2d66be179, "as", ""},
    {0x631c37ce72a97393, "asd", ""},
    {0x415872f599cea71e, "asdf", ""},
    {
        // Exactly 63 characters, which exercises all code paths.
        0x02a2e85470d6fd96,
        "Call me Ishmael. Some years ago--never mind how long precisely-",
        "",
    },
    {
        0x93267f9820452ead,
        "The quick brown fox jumps over the lazy dog http://i.imgur.com/VHQXScB.gif",
        "",
    },
}

func init() {
    for i := range testdata {
        d := &testdata[i]
        if len(d.data) > 20 {
            d.printable = d.data[:20]
        } else {
            d.printable = d.data
        }
    }
}

func TestBlockSize(t *testing.T) {
    xxh := xxhash64.New()
    if s := xxh.BlockSize(); s <= 0 {
        t.Errorf("invalid BlockSize: %d", s)
    }

}

func TestSize(t *testing.T) {
    xxh := xxhash64.New()
    if s := xxh.Size(); s != 8 {
        t.Errorf("invalid Size: got %d expected 8", s)
    }
}

func TestData(t *testing.T) {
    for i, td := range testdata {
        xxh := xxhash64.New()
        data := []byte(td.data)
        xxh.Write(data)
        if h := xxh.Sum64(); h != td.sum {
            t.Errorf("test %d: xxh64(%s)=0x%x expected 0x%x", i, td.printable, h, td.sum)
            t.FailNow()
        }

        if h := xxhash64.Checksum(data); h != td.sum {
            t.Errorf("test %d: xxh64(%s)=0x%x expected 0x%x", i, td.printable, h, td.sum)
            t.FailNow()
        }

        sum := xxhash64.Sum(data)
        sumh := binary.BigEndian.Uint64(sum[:])
        if sumh != td.sum {
            t.Errorf("test %d: Sum(%s)=0x%x expected 0x%x", i, td.printable, sumh, td.sum)
            t.FailNow()
        }

        // =============

        xxh = xxhash64.NewWithSeed(0)
        data = []byte(td.data)
        xxh.Write(data)
        if h := xxh.Sum64(); h != td.sum {
            t.Errorf("test NewWithSeed %d: xxh64(%s)=0x%x expected 0x%x", i, td.printable, h, td.sum)
            t.FailNow()
        }

        if h := xxhash64.ChecksumWithSeed(data, 0); h != td.sum {
            t.Errorf("test ChecksumWithSeed %d: xxh64(%s)=0x%x expected 0x%x", i, td.printable, h, td.sum)
            t.FailNow()
        }

        sum = xxhash64.SumWithSeed(data, 0)
        sumh = binary.BigEndian.Uint64(sum[:])
        if sumh != td.sum {
            t.Errorf("test %d: SumWithSeed(%s, 0)=0x%x expected 0x%x", i, td.printable, sumh, td.sum)
            t.FailNow()
        }

    }
}

func TestSplitData(t *testing.T) {
    for i, td := range testdata {
        xxh := xxhash64.New()
        data := []byte(td.data)
        l := len(data) / 2
        xxh.Write(data[0:l])
        xxh.Write(data[l:])
        h := xxh.Sum64()
        if h != td.sum {
            t.Errorf("test %d: xxh64(%s)=0x%x expected 0x%x", i, td.printable, h, td.sum)
            t.FailNow()
        }
    }
}

func TestSum(t *testing.T) {
    for i, td := range testdata {
        xxh := xxhash64.New()
        data := []byte(td.data)
        xxh.Write(data)
        b := xxh.Sum(data)
        if h := binary.BigEndian.Uint64(b[len(data):]); h != td.sum {
            t.Errorf("test %d: xxh64(%s)=0x%x expected 0x%x", i, td.printable, h, td.sum)
            t.FailNow()
        }
    }
}

func TestReset(t *testing.T) {
    xxh := xxhash64.New()
    for i, td := range testdata {
        xxh.Write([]byte(td.data))
        h := xxh.Sum64()
        if h != td.sum {
            t.Errorf("test %d: xxh64(%s)=0x%x expected 0x%x", i, td.printable, h, td.sum)
            t.FailNow()
        }
        xxh.Reset()
    }
}

type testString struct {
    sum  string
    data string
}

var testdataStrings = []testString{
    {"ef46db3751d8e999", ""},
    {"d24ec4f1a98c6e5b", "a"},
    {"65f708ca92d04a61", "ab"},
    {"44bc2cf5ad770999", "abc"},
    {"de0327b0d25d92cc", "abcd"},
    {"07e3670c0c8dc7eb", "abcde"},
    {"fa8afd82c423144d", "abcdef"},
    {"1860940e2902822d", "abcdefg"},
    {"3ad351775b4634b7", "abcdefgh"},
    {"27f1a34fdbb95e13", "abcdefghi"},
    {"d6287a1de5498bb2", "abcdefghij"},
    {"bf2cd639b4143b80", "abcdefghijklmnopqrstuvwxyz012345"},
    {"64f23ecf1609b766", "abcdefghijklmnopqrstuvwxyz0123456789"},
    {"c5a8b11443765630", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."},

    {"1c330fb2d66be179", "as"},
    {"631c37ce72a97393", "asd"},
    {"415872f599cea71e", "asdf"},
    {
        // Exactly 63 characters, which exercises all code paths.
        "02a2e85470d6fd96",
        "Call me Ishmael. Some years ago--never mind how long precisely-",
    },
    {
        "93267f9820452ead",
        "The quick brown fox jumps over the lazy dog http://i.imgur.com/VHQXScB.gif",
    },
}

func TestDataString(t *testing.T) {
    for i, td := range testdataStrings {
        data := []byte(td.data)

        xxh := xxhash64.New()
        xxh.Write(data)
        h := xxh.Sum(nil)
        if fmt.Sprintf("%x", h) != td.sum {
            t.Errorf("test %d: got %x, want %x", i, h, td.sum)
        }

        sum := xxhash64.Sum(data)
        if fmt.Sprintf("%x", sum) != td.sum {
            t.Errorf("test %d: got %x, want %x", i, sum, td.sum)
        }

    }
}

// =============

var testdata1 = []byte(testdata[len(testdata)-1].data)

func Benchmark_XXH64(b *testing.B) {
    h := xxhash64.New()
    for n := 0; n < b.N; n++ {
        h.Write(testdata1)
        h.Sum64()
        h.Reset()
    }
}

func Benchmark_XXH64_Checksum(b *testing.B) {
    for n := 0; n < b.N; n++ {
        xxhash64.Checksum(testdata1)
    }
}

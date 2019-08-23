package main

import (
    "encoding/binary"
    "errors"
    "hash"
    "net/http"
    "net/url"
    "log"
    "io/ioutil"
    "fmt"
    "os"
    "bufio"
    "math/bits"
    "encoding/hex"
)

const Size = 16

const BlockSize = 64

const (
    init0 = 0x00000001
    init1 = 0xEFCDAB89
    init2 = 0x98BADCFE
    init3 = 0x10000000
    _K0 = 0x5A827999
    _K1 = 0x6ED9EBA1
    _K2 = 0x8F1BBCDC
    _K3 = 0xCA62C1D6

)
func blockGeneric(dig *digest, p []byte) {
    // load state
    a, b, c, d := dig.s[0], dig.s[1], dig.s[2], dig.s[3]

    for i := 0; i <= len(p)-BlockSize; i += BlockSize {
        // eliminate bounds checks on p
        q := p[i:]
        q = q[:BlockSize:BlockSize]

        // save current state
        aa, bb, cc, dd := a, b, c, d

        // load input block
        x0 := binary.LittleEndian.Uint32(q[4*0x0:])
        x1 := binary.LittleEndian.Uint32(q[4*0x1:])
        x2 := binary.LittleEndian.Uint32(q[4*0x2:])
        x3 := binary.LittleEndian.Uint32(q[4*0x3:])
        x4 := binary.LittleEndian.Uint32(q[4*0x4:])
        x5 := binary.LittleEndian.Uint32(q[4*0x5:])
        x6 := binary.LittleEndian.Uint32(q[4*0x6:])
        x7 := binary.LittleEndian.Uint32(q[4*0x7:])
        x8 := binary.LittleEndian.Uint32(q[4*0x8:])
        x9 := binary.LittleEndian.Uint32(q[4*0x9:])
        xa := binary.LittleEndian.Uint32(q[4*0xa:])
        xb := binary.LittleEndian.Uint32(q[4*0xb:])
        xc := binary.LittleEndian.Uint32(q[4*0xc:])
        xd := binary.LittleEndian.Uint32(q[4*0xd:])
        xe := binary.LittleEndian.Uint32(q[4*0xe:])
        xf := binary.LittleEndian.Uint32(q[4*0xf:])

        // round 1
        a = b + bits.RotateLeft32((((c^d)&b)^d)+a+x0+0xd76aa478, 7)
        d = a + bits.RotateLeft32((((b^c)&a)^c)+d+x1+0xe8c7b756, 12)
        c = d + bits.RotateLeft32((((a^b)&d)^b)+c+x2+0x242070db, 17)
        b = c + bits.RotateLeft32((((d^a)&c)^a)+b+x3+0xc1bdceee, 22)
        a = b + bits.RotateLeft32((((c^d)&b)^d)+a+x4+0xf57c0faf, 7)
        d = a + bits.RotateLeft32((((b^c)&a)^c)+d+x5+0x4787c62a, 12)
        c = d + bits.RotateLeft32((((a^b)&d)^b)+c+x6+0xa8304613, 17)
        b = c + bits.RotateLeft32((((d^a)&c)^a)+b+x7+0xfd469501, 22)
        a = b + bits.RotateLeft32((((c^d)&b)^d)+a+x8+0x698098d8, 7)
        d = a + bits.RotateLeft32((((b^c)&a)^c)+d+x9+0x8b44f7af, 12)
        c = d + bits.RotateLeft32((((a^b)&d)^b)+c+xa+0xffff5bb1, 17)
        b = c + bits.RotateLeft32((((d^a)&c)^a)+b+xb+0x895cd7be, 22)
        a = b + bits.RotateLeft32((((c^d)&b)^d)+a+xc+0x6b901122, 7)
        d = a + bits.RotateLeft32((((b^c)&a)^c)+d+xd+0xfd987193, 12)
        c = d + bits.RotateLeft32((((a^b)&d)^b)+c+xe+0xa679438e, 17)
        b = c + bits.RotateLeft32((((d^a)&c)^a)+b+xf+0x49b40821, 22)

        // round 2
        a = b + bits.RotateLeft32((((b^c)&d)^c)+a+x1+0xf61e2562, 5)
        d = a + bits.RotateLeft32((((a^b)&c)^b)+d+x6+0xc040b340, 9)
        c = d + bits.RotateLeft32((((d^a)&b)^a)+c+xb+0x265e5a51, 14)
        b = c + bits.RotateLeft32((((c^d)&a)^d)+b+x0+0xe9b6c7aa, 20)
        a = b + bits.RotateLeft32((((b^c)&d)^c)+a+x5+0xd62f105d, 5)
        d = a + bits.RotateLeft32((((a^b)&c)^b)+d+xa+0x02441453, 9)
        c = d + bits.RotateLeft32((((d^a)&b)^a)+c+xf+0xd8a1e681, 14)
        b = c + bits.RotateLeft32((((c^d)&a)^d)+b+x4+0xe7d3fbc8, 20)
        a = b + bits.RotateLeft32((((b^c)&d)^c)+a+x9+0x21e1cde6, 5)
        d = a + bits.RotateLeft32((((a^b)&c)^b)+d+xe+0xc33707d6, 9)
        c = d + bits.RotateLeft32((((d^a)&b)^a)+c+x3+0xf4d50d87, 14)
        b = c + bits.RotateLeft32((((c^d)&a)^d)+b+x8+0x455a14ed, 20)
        a = b + bits.RotateLeft32((((b^c)&d)^c)+a+xd+0xa9e3e905, 5)
        d = a + bits.RotateLeft32((((a^b)&c)^b)+d+x2+0xfcefa3f8, 9)
        c = d + bits.RotateLeft32((((d^a)&b)^a)+c+x7+0x676f02d9, 14)
        b = c + bits.RotateLeft32((((c^d)&a)^d)+b+xc+0x8d2a4c8a, 20)

        // round 3
        a = b + bits.RotateLeft32((b^c^d)+a+x5+0xfffa3942, 4)
        d = a + bits.RotateLeft32((a^b^c)+d+x8+0x8771f681, 11)
        c = d + bits.RotateLeft32((d^a^b)+c+xb+0x6d9d6122, 16)
        b = c + bits.RotateLeft32((c^d^a)+b+xe+0xfde5380c, 23)
        a = b + bits.RotateLeft32((b^c^d)+a+x1+0xa4beea44, 4)
        d = a + bits.RotateLeft32((a^b^c)+d+x4+0x4bdecfa9, 11)
        c = d + bits.RotateLeft32((d^a^b)+c+x7+0xf6bb4b60, 16)
        b = c + bits.RotateLeft32((c^d^a)+b+xa+0xbebfbc70, 23)
        a = b + bits.RotateLeft32((b^c^d)+a+xd+0x289b7ec6, 4)
        d = a + bits.RotateLeft32((a^b^c)+d+x0+0xeaa127fa, 11)
        c = d + bits.RotateLeft32((d^a^b)+c+x3+0xd4ef3085, 16)
        b = c + bits.RotateLeft32((c^d^a)+b+x6+0x04881d05, 23)
        a = b + bits.RotateLeft32((b^c^d)+a+x9+0xd9d4d039, 4)
        d = a + bits.RotateLeft32((a^b^c)+d+xc+0xe6db99e5, 11)
        c = d + bits.RotateLeft32((d^a^b)+c+xf+0x1fa27cf8, 16)
        b = c + bits.RotateLeft32((c^d^a)+b+x2+0xc4ac5665, 23)

        // round 4
        a = b + bits.RotateLeft32((c^(b|^d))+a+x0+0xf4292244, 6)
        d = a + bits.RotateLeft32((b^(a|^c))+d+x7+0x432aff97, 10)
        c = d + bits.RotateLeft32((a^(d|^b))+c+xe+0xab9423a7, 15)
        b = c + bits.RotateLeft32((d^(c|^a))+b+x5+0xfc93a039, 21)
        a = b + bits.RotateLeft32((c^(b|^d))+a+xc+0x655b59c3, 6)
        d = a + bits.RotateLeft32((b^(a|^c))+d+x3+0x8f0ccc92, 10)
        c = d + bits.RotateLeft32((a^(d|^b))+c+xa+0xffeff47d, 15)
        b = c + bits.RotateLeft32((d^(c|^a))+b+x1+0x85845dd1, 21)
        a = b + bits.RotateLeft32((c^(b|^d))+a+x8+0x6fa87e4f, 6)
        d = a + bits.RotateLeft32((b^(a|^c))+d+xf+0xfe2ce6e0, 10)
        c = d + bits.RotateLeft32((a^(d|^b))+c+x6+0xa3014314, 15)
        b = c + bits.RotateLeft32((d^(c|^a))+b+xd+0x4e0811a1, 21)
        a = b + bits.RotateLeft32((c^(b|^d))+a+x4+0xf7537e82, 6)
        d = a + bits.RotateLeft32((b^(a|^c))+d+xb+0xbd3af235, 10)
        c = d + bits.RotateLeft32((a^(d|^b))+c+x2+0x2ad7d2bb, 15)
        b = c + bits.RotateLeft32((d^(c|^a))+b+x9+0xeb86d391, 21)

        // add saved state
        a += aa
        b += bb
        c += cc
        d += dd
    }

    // save state
    dig.s[0], dig.s[1], dig.s[2], dig.s[3] = a, b, c, d
}



// gE#2Tv

// digest represents the partial evaluation of a checksum.
type digest struct {
    s   [4]uint32
    x   [BlockSize]byte
    nx  int
    len uint64
}

func (d *digest) Reset() {
    d.nx = 0
    d.len = 0
}

const (
    magic         = "hug\x01"
    marshaledSize = len(magic) + 4*4 + BlockSize + 8
)

func (d *digest) MarshalBinary() ([]byte, error) {
    b := make([]byte, 0, marshaledSize)
    b = append(b, magic...)
    b = appendUint32(b, d.s[0])
    b = appendUint32(b, d.s[1])
    b = appendUint32(b, d.s[2])
    b = appendUint32(b, d.s[3])
    b = append(b, d.x[:d.nx]...)
    b = b[:len(b)+len(d.x)-d.nx] // already zero
    b = appendUint64(b, d.len)
    return b, nil
}

func (d *digest) UnmarshalBinary(b []byte) error {
    if len(b) < len(magic) || string(b[:len(magic)]) != magic {
        return errors.New("UberCrypt: invalid state")
    }
    if len(b) != marshaledSize {
        return errors.New("UberCrypt: invalid size")
    }
    b = b[len(magic):]
    b, d.s[0] = consumeUint32(b)
    b, d.s[1] = consumeUint32(b)
    b, d.s[2] = consumeUint32(b)
    b, d.s[3] = consumeUint32(b)
    b = b[copy(d.x[:], b):]
    b, d.len = consumeUint64(b)
    d.nx = int(d.len % BlockSize)
    return nil
}

func appendUint64(b []byte, x uint64) []byte {
    var a [8]byte
    binary.BigEndian.PutUint64(a[:], x)
    return append(b, a[:]...)
}

func appendUint32(b []byte, x uint32) []byte {
    var a [4]byte
    binary.BigEndian.PutUint32(a[:], x)
    return append(b, a[:]...)
}

func consumeUint64(b []byte) ([]byte, uint64) {
    return b[8:], binary.BigEndian.Uint64(b[0:8])
}

func consumeUint32(b []byte) ([]byte, uint32) {
    return b[4:], binary.BigEndian.Uint32(b[0:4])
}

func New() hash.Hash {
    d := new(digest)
    return d
}

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return BlockSize }
func (d *digest) Write(p []byte) (nn int, err error) {
    nn = len(p)
    d.len += uint64(nn)
    if d.nx > 0 {
        n := copy(d.x[d.nx:], p)
        d.nx += n
        if d.nx == BlockSize {
            blockGeneric(d, d.x[:])
            d.nx = 0
        }
        p = p[n:]
    }
    if len(p) >= BlockSize {
        n := len(p) &^ (BlockSize - 1)
        blockGeneric(d, p[:n])
        p = p[n:]
    }
    if len(p) > 0 {
        d.nx = copy(d.x[:], p)
    }
    return
}

func (d *digest) Sum(in []byte) []byte {
    d0 := *d
    hash := d0.checkSum()
    return append(in, hash[:]...)
}

func (d *digest) checkSum() [Size]byte {
    tmp := [1 + 63 + 8]byte{0x80}
    pad := (55 - d.len) % 64
    binary.LittleEndian.PutUint64(tmp[1+pad:], d.len<<3)
    d.Write(tmp[:1+pad+8])

    if d.nx != 0 {
        panic("d.nx != 0")
    }

    var digest [Size]byte
    binary.LittleEndian.PutUint32(digest[0:], d.s[0])
    binary.LittleEndian.PutUint32(digest[4:], d.s[1])
    binary.LittleEndian.PutUint32(digest[8:], d.s[2])
    binary.LittleEndian.PutUint32(digest[12:], d.s[3])
    return digest
}

func Sum(data []byte, key []byte) []byte {
    var d digest
    d.Reset()
    fmt.Println("Setting up some constants")
    d.s[0] = init0 + uint32(key[0]) << 24 + uint32(key[1]) << 16 + uint32(key[2]) << 8
    d.s[1] = init1
    d.s[2] = init2
    d.s[3] = init3 + uint32(key[3]) << 16 + uint32(key[4]) << 8 + uint32(key[5])
    fmt.Println(fmt.Sprintf("s0: %x, s1: %x, s2: %x, s3: %x", d.s[0],d.s[1],d.s[2],d.s[3]))
    d.Write(data)
    arr := d.checkSum()
    slice := arr[:]
    return slice
}

func decode_packet(packet []byte) ([]byte) {

    for i, j := 0, len(packet)-1; i < j; i, j = i+1, j-1 {
        packet[i], packet[j] = packet[j], packet[i]
    }

    b := []byte("");

    for i:=0; i < len(packet); i++ {

        b = append(b, packet[i] ^ 0x42)

    }
    return b

}

func check_user_key(user_input []byte) (bool) {

    if len(user_input) != 6 {
        fmt.Println("Incorrect initialization")
    }

    checksum := hex.EncodeToString(Sum([]byte("custom"), user_input))

    if checksum == "8b9035807842a4e4dbe009f3f1478127" {
        return true
    }
    return false

}

func get_flag(user_input string) ([]byte) {

    resp, err := http.PostForm("http://svieg.com:5000/UberCrypt/NotMalware", url.Values{"key": {user_input}})

    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()

    packet, err := ioutil.ReadAll(resp.Body)

    return decode_packet(packet)


}

func main() {

    fmt.Print("Hey, do you have the key?: ");
    user_input,_,err := bufio.NewReader(os.Stdin).ReadLine();
    if err != nil {
            fmt.Println("Invalid input :/ , ",err);
    }
    valid := check_user_key(user_input)

    if valid {
        flag := get_flag(string(user_input))
        fmt.Println(fmt.Sprintf("Flag: %s\n", string(flag)))
    } else {
        fmt.Println("Install Flash Player to get access to these features.");
    }
}

func reverseBits(num uint32) uint32 {
    var res uint32

    for i := range 32 {
        bit := (num >> i) & 1
        res = res | (bit << (31 - i))
    }

    return res
}

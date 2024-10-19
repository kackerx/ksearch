package main

func IsBitOne(n uint64, i int) bool {
	return (n & (1 << (i - 1))) != 0
}

func SetBit(n uint64, i int) uint64 {
	return n | (1 << (i - 1))
}

const (
	MALE = 1 << iota
	VIP
	WeekActive
)

type Candidate struct {
	ID     int
	Gender string
	VIP    bool // 是否vip
	Active int  // 几天内活跃
	Bits   uint64
}

func (c *Candidate) SetMale() {
	c.Gender = "男"
	c.Bits |= MALE
}

func (c *Candidate) SetVIP() {
	c.VIP = true
	c.Bits |= VIP
}

func (c *Candidate) SetAction(day int) {
	c.Active = day
	if day <= 7 {
		c.Bits |= WeekActive
	}
}

type BitMap struct {
	Table uint64
}

func NewBitMap(min int, arr []int) *BitMap {
	bitMap := new(BitMap)
	for _, v := range arr {
		bitMap.Table = SetBit(bitMap.Table, v-min)
	}

	return bitMap
}

func IntersectionOfBitMap(bm1, bm2 *BitMap, min int) []int {
	res := make([]int, 0)
	s := bm1.Table & bm2.Table
	for i := 1; i <= 64; i++ {
		if IsBitOne(s, i) {
			res = append(res, i+min)
		}
	}

	return res
}

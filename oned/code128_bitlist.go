package oned

import "fmt"

type code128_BitList struct {
	cap      uint32
	data     []uint8
}

func NewCode128_BitList() *code128_BitList {
	res := code128_BitList{
		cap: 0,
		data: []uint8{} ,
	}

	return &res
}

func (this *code128_BitList) GetCap() uint32 {
	return this.cap
}

func (this *code128_BitList) GetBytes() []byte {
	l := len (this.data)
	res := make ([]byte, l)
	for i:=0; i<l; i++ {
		res[i] = this.data[i]
	}
	return res
}

func (this *code128_BitList) AddByte(b uint8) {
	for i:=7; i>=0; i-- {
		this.AddBit(((b >> uint(i)) & 0x01) == 0x01)
	}
}

func (this *code128_BitList) AddBit(bit bool) {
	pos := this.cap

	this.cap++
	neededBytes := (this.cap + 7) / 8

	if neededBytes > uint32(len(this.data)) {
		this.data = append(this.data, 0)
	}
	this.SetBit(pos, bit)
}

func (this *code128_BitList) AddBits(bits ... bool) {
	for _, bit := range bits {
		this.AddBit (bit)
	}
}

func (this *code128_BitList) SetBit(index uint32, bit bool) {
	byteIdx :=      index / 8
	bitIdx  := 7 - (index % 8)

	if bit {
		this.data[byteIdx] = this.data[byteIdx] |  (1 << bitIdx)
	} else {
		this.data[byteIdx] = this.data[byteIdx] & ^(1 << bitIdx)
	}
}

func (this *code128_BitList) GetBit(index uint32) bool {
	byteIdx :=      index / 8
	bitIdx  := 7 - (index % 8)

	return (this.data[byteIdx] & (1 << bitIdx)) == (1 << bitIdx)
}

func (this *code128_BitList) String() string {
	tmp := ""
	for _, b := range this.data {
		tmp = tmp + fmt.Sprintf(" %02x", b)
	}
	return fmt.Sprintf("{ %02x [%s ]}", this.cap, tmp)
}

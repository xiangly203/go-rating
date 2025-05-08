package utils

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const saltBit = uint(8)
const saltShift = uint(8)
const increasShift = saltBit + saltShift

type Mist struct {
	sync.Mutex
	increas int64
	saltA   int64
	saltB   int64
}

func NewMist() *Mist {
	mist := Mist{increas: 1}
	return &mist
}

func (c *Mist) Generate() int64 {
	c.Lock()
	c.increas++
	randA, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltA = randA.Int64()
	randB, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltB = randB.Int64()
	// 通过位运算实现自动占位
	mist := int64((c.increas << increasShift) | (c.saltA << saltShift) | c.saltB)
	c.Unlock()
	return mist
}

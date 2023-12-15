package utils

import (
	"fmt"
	"sync"
	"time"
)

const (
	twepoch            = int64(1420041600000)
	workerIdBits       = 5
	datacenterIdBits   = 5
	maxWorkerId        = -1 ^ (-1 << workerIdBits)
	maxDatacenterId    = -1 ^ (-1 << datacenterIdBits)
	sequenceBits       = 12
	workerIdShift      = sequenceBits
	datacenterIdShift  = sequenceBits + workerIdBits
	timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits
	sequenceMask       = -1 ^ (-1 << sequenceBits)
)

type SnowflakeDistributeId struct {
	mu            sync.Mutex
	lastTimestamp int64
	workerId      int64
	datacenterId  int64
	sequence      int64
}

func NewSnowflakeId(workerId int64, datacenterId int64) *SnowflakeDistributeId {
	if workerId > maxWorkerId || workerId < 0 {
		panic(fmt.Sprintf("worker Id can't be greater than %d or less than 0", maxWorkerId))
	}
	if datacenterId > maxDatacenterId || datacenterId < 0 {
		panic(fmt.Sprintf("datacenter Id can't be greater than %d or less than 0", maxDatacenterId))
	}
	return &SnowflakeDistributeId{
		lastTimestamp: -1,
		workerId:      workerId,
		datacenterId:  datacenterId,
	}
}

func (s *SnowflakeDistributeId) NextId() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := timeGen()
	if timestamp < s.lastTimestamp {
		panic(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", s.lastTimestamp-timestamp))
	}

	if s.lastTimestamp == timestamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			timestamp = tilNextMillis(s.lastTimestamp)
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp
	return ((timestamp - twepoch) << timestampLeftShift) | (s.datacenterId << datacenterIdShift) | (s.workerId << workerIdShift) | s.sequence
}

func timeGen() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func tilNextMillis(lastTimestamp int64) int64 {
	timestamp := timeGen()
	for timestamp <= lastTimestamp {
		timestamp = timeGen()
	}
	return timestamp
}

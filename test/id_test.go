package test

import (
	"fmt"
	"go_gin/utils"
	"testing"
	"time"
)

func TestId(t *testing.T) {
	n := 10
	snowflakeID := make([]int64, n)
	mistId := make([]int64, n)
	str := make([]string, n)
	logidStr := make([]string, n)
	mist := utils.NewMist()
	id := utils.NewSnowflakeId(1, 1)
	logid := utils.NewLogid()
	snowStart := time.Now()

	for i := 0; i < n; i++ {
		snowflakeID[i] = id.NextId()
	}
	snowEnd := time.Now()
	mistStart := time.Now()
	for i := 0; i < n; i++ {
		mistId[i] = mist.Generate()
	}
	mistEnd := time.Now()

	strStart := time.Now()
	for i := 0; i < n; i++ {
		str[i] = utils.RandStr(64)
	}
	strEnd := time.Now()

	logidStart := time.Now()
	for i := 0; i < n; i++ {
		logidStr[i] = logid.NextLogid()
	}
	logidEnd := time.Now()

	fmt.Println(snowEnd.Sub(snowStart))
	fmt.Println(mistEnd.Sub(mistStart))
	fmt.Println(strEnd.Sub(strStart))
	fmt.Println(logidEnd.Sub(logidStart))
}

package utils

import (
	"go_gin/config"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex

type Logid struct {
	sync.Mutex
}

func NewLogid() *Logid {
	return &Logid{}
}
func (c *Logid) NextLogid() string {
	c.Lock()
	defer c.Unlock()
	now := time.Now()
	nowStr := now.Format("200601021504")
	rStr := RandStr(config.LenLogid - len(nowStr))
	var sb strings.Builder
	sb.WriteString(nowStr)
	sb.WriteString(rStr)
	return sb.String()
}

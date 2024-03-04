package test

import (
	"fmt"
	"go_gin/utils"
	"testing"
)

func TestLogid(t *testing.T) {
	logid := utils.NewLogid()
	fmt.Println(logid.NextLogid())
}

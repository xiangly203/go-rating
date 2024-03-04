package test

import (
	"go_gin/utils"
	"testing"
)

func TestUpload(t *testing.T) {
	n, _ := utils.NewBucketBasics()
	err := n.UploadFile("xyi-next-app", "testa.png", "test.png")
	if err != nil {
		return
	}
}

package test

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	// 创建
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)

	// 解析
	u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}

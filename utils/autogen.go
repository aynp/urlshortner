package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func TimestampString() string {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	return timestamp
}

func UUIDString() string {
	return uuid.New().String()
}

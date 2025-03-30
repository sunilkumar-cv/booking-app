package utils

import (
	"fmt"
	"time"
)

func Log(action string) {
	fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05"), action)
}

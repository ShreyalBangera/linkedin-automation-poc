package stealth

import (
	"math/rand"
	"time"
)

func HumanDelay() {
	delay := rand.Intn(2000) + 1000 // 1–3 seconds
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

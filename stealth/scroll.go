package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// HumanLikeScroll simulates natural scrolling using mouse wheel
func HumanLikeScroll(page *rod.Page) {
	mouse := page.Mouse

	scrollSteps := rand.Intn(5) + 5 // 5–10 scrolls

	for i := 0; i < scrollSteps; i++ {
		// scroll down like a human
		scroll := rand.Intn(300) + 200
		mouse.Scroll(0, float64(scroll), rand.Intn(5)+5)

		// pause as if reading
		time.Sleep(time.Duration(rand.Intn(2000)+1000) * time.Millisecond)

		// sometimes scroll up a bit
		if rand.Intn(10) > 7 {
			mouse.Scroll(0, -100, rand.Intn(3)+3)
			time.Sleep(800 * time.Millisecond)
		}
	}
}

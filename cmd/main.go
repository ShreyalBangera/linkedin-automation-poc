package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/joho/godotenv"

	"linkedin-automation/search"
	"linkedin-automation/stealth"
	"linkedin-automation/storage"
)

func main() {
	godotenv.Load()

	email := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	u := launcher.New().
		Leakless(false).
		Headless(false).
		MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()
	page := browser.MustPage("https://www.linkedin.com")
	page.MustWaitLoad()

	// Load cookies
	if storage.LoadCookies(page, "cookies.json") == nil {
		fmt.Println("🍪 Cookies loaded")
		page.MustNavigate("https://www.linkedin.com/feed/")
		page.MustWaitLoad()
	}

	time.Sleep(4 * time.Second)

	if !strings.Contains(page.MustInfo().URL, "/feed") {
		// Login fallback
		page.MustNavigate("https://www.linkedin.com/login")
		page.MustWaitLoad()

		stealth.HumanDelay()
		stealth.HumanType(page.MustElement("#username"), email)
		stealth.HumanDelay()
		stealth.HumanType(page.MustElement("#password"), password)
		stealth.HumanDelay()
		page.MustElement(`button[type="submit"]`).MustClick()

		time.Sleep(8 * time.Second)

		if strings.Contains(page.MustInfo().URL, "/feed") {
			fmt.Println("✅ Login successful, saving cookies")
			storage.SaveCookies(page, "cookies.json")
		}
	}

	// 🔍 SEARCH PROFILES
	profiles := search.SearchPeople(page, "software engineer")

	fmt.Println("👤 Profiles found:")
	for i, p := range profiles {
		if i >= 10 {
			break
		}
		fmt.Println(p)
	}

	select {}
}

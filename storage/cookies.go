package storage

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// Simple serializable cookie
type SimpleCookie struct {
	Name     string  `json:"name"`
	Value    string  `json:"value"`
	Domain   string  `json:"domain"`
	Path     string  `json:"path"`
	Expires  float64 `json:"expires"`
	Secure   bool    `json:"secure"`
	HTTPOnly bool    `json:"httpOnly"`
}

// Save cookies to JSON
func SaveCookies(page *rod.Page, path string) error {
	cookies := page.MustCookies()
	var simple []SimpleCookie

	for _, c := range cookies {
		simple = append(simple, SimpleCookie{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  float64(c.Expires), // ✅ FIX
			Secure:   c.Secure,
			HTTPOnly: c.HTTPOnly,
		})
	}

	data, err := json.MarshalIndent(simple, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// Load cookies from JSON
func LoadCookies(page *rod.Page, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var simple []SimpleCookie
	if err := json.Unmarshal(data, &simple); err != nil {
		return err
	}

	var params []*proto.NetworkCookieParam
	for _, c := range simple {
		exp := proto.TimeSinceEpoch(c.Expires) // ✅ FIX

		params = append(params, &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  exp,
			Secure:   c.Secure,
			HTTPOnly: c.HTTPOnly,
		})
	}

	return page.SetCookies(params)
}

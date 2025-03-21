package shared

import "net/url"

const (
	WebAppURL = "https://activity-tracker-93db9.web.app"
)

// IsValidURL checks if the link is a valid URL
func IsValidURL(link string) bool {
	parsedURL, err := url.Parse(link)

	return err == nil && parsedURL != nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}

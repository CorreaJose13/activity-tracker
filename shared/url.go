package shared

import "net/url"

// IsValidURL checks if the link is a valid URL
func IsValidURL(link string) bool {
	parsedURL, err := url.Parse(link)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}

	return true
}

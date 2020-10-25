package helpers

import (
	"net/url"
)

func RedirectHttps(rawUrl string) string {
	url, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	url.Scheme = "https"

	return url.String()
}

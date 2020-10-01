package helper

import (
	"net/url"
)

func JoinURL(baseUrl *url.URL, pathUrlStr string) (joinedUrlStr *url.URL, err error) {
	pathUrl, err := url.Parse(pathUrlStr)
	if err != nil {
		return
	}

	joinedUrlStr = baseUrl.ResolveReference(pathUrl)
	return
}

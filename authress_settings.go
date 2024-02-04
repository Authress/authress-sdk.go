package authress

import (
	"net/url"
)

type AuthressSettings struct {
	AuthressApiUrl         *url.URL
	ServiceClientAccessKey string
	UserAgent              string
}

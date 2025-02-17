package constants

import "net/textproto"

var (
	XServiceName = textproto.CanonicalMIMEHeaderKey("X-service-name")
	XApiKey = textproto.CanonicalMIMEHeaderKey("X-api-key")
	XRequestAt = textproto.CanonicalMIMEHeaderKey("X-request-at")
	Authorization = textproto.CanonicalMIMEHeaderKey("authorization")
)
package http

import (
	"net"
	"net/http"
	"time"
)

// DefaultHTTPHeaders returns a secure by default collection of recommended
// HTTP response headers intended for user privacy and security.
//
// This is a good starting point for applications. A client can then choose
// to override specific headers as necessary.
//
// These are the headers that are currently returned by this function:
//
// Content-Type: text/plain; charset=utf-8
// Cache-Control: no-store, max-age=0
// X-Content-Type-Options: nosniff
// Strict-Transport-Security: max-age=63072000; includeSubDomains; preload
// X-Frame-Options: DENY
// X-Robots-Tag: noindex
// Content-Security-Policy: default-src 'self'; base-uri 'none'; form-action 'self'; frame-ancestors 'none'; object-src 'none'; upgrade-insecure-requests
// Referrer-Policy: no-referrer
// X-XSS-Protection: 1; mode=block
// Cross-Origin-Embedder-Policy: require-corp
// Cross-Origin-Opener-Policy: same-origin
// Cross-Origin-Resource-Policy: same-origin
// Permissions-Policy: accelerometer=(),autoplay=(),camera=(),display-capture=(),document-domain=(),encrypted-media=(),fullscreen=(),geolocation=(),gyroscope=(),magnetometer=(),microphone=(),midi=(),payment=(),picture-in-picture=(),publickey-credentials-get=(),screen-wake-lock=(),sync-xhr=(self),usb=(),web-share=(),xr-spatial-tracking=()
func DefaultHTTPHeaders() http.Header {
	headers := make(http.Header)
	headers.Set("Content-Type", "text/plain; charset=utf-8")
	headers.Set("Cache-Control", "no-store, max-age=0")
	headers.Set("X-Content-Type-Options", "nosniff")
	headers.Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	headers.Set("X-Frame-Options", "DENY")
	headers.Set("X-Robots-Tag", "noindex")
	headers.Set("Content-Security-Policy", "default-src 'self'; base-uri 'none'; form-action 'self'; frame-ancestors 'none'; object-src 'none'; upgrade-insecure-requests")
	headers.Set("Referrer-Policy", "no-referrer")
	headers.Set("X-XSS-Protection", "1; mode=block")
	headers.Set("Cross-Origin-Embedder-Policy", "require-corp")
	headers.Set("Cross-Origin-Opener-Policy", "same-origin")
	headers.Set("Cross-Origin-Resource-Policy", "same-origin")
	headers.Set("Permissions-Policy", "accelerometer=(),autoplay=(),camera=(),display-capture=(),document-domain=(),encrypted-media=(),fullscreen=(),geolocation=(),gyroscope=(),magnetometer=(),microphone=(),midi=(),payment=(),picture-in-picture=(),publickey-credentials-get=(),screen-wake-lock=(),sync-xhr=(self),usb=(),web-share=(),xr-spatial-tracking=()")

	return headers
}

// DefaultHTTPErrorHeaders returns a secure by default collection of
// recommended HTTP response headers intended for user privacy and security.
//
// DefaultHTTPHeaders() are used and modified as such:
//
// Content-Type: text/html; charset=utf-8
//
// All headers from DefaultHTTPHeaders are returned along with the
// modifications above.
func DefaultHTTPErrorHeaders() http.Header {
	headers := DefaultHTTPHeaders()
	headers.Set("Content-Type", "text/html; charset=utf-8")

	return headers
}

// DefaultHTTPErrorHeadersForJSON returns a secure by default collection of
// recommended HTTP response headers intended for user privacy and security.
//
// DefaultHTTPHeaders() are used and modified as such:
//
// Content-Type: application/json; charset=utf-8
//
// All headers from DefaultHTTPHeaders are returned along with the
// modifications above.
func DefaultHTTPErrorHeadersForJSON() http.Header {
	headers := DefaultHTTPHeaders()
	headers.Set("Content-Type", "application/json; charset=utf-8")

	return headers
}

type Option struct {
	transportDialContextTimeout    time.Duration
	transportDialContextKeepAlive  time.Duration
	transportDialContextDeadline   time.Time
	transportForceAttemptHTTP2     bool
	transportMaxIdleConns          int
	transportIdleConnTimeout       time.Duration
	transportTLSHandshakeTimeout   time.Duration
	transportResponseHeaderTimeout time.Duration
	transportExpectContinueTimeout time.Duration
	clientTimeout                  time.Duration
}

func WithConnectTimeout(t time.Duration) func(o *Option) {
	return func(o *Option) {
		o.transportDialContextTimeout = t
	}
}

func WithKeepAlive(t time.Duration) func(o *Option) {
	return func(o *Option) {
		o.transportDialContextKeepAlive = t
	}
}

func WithDeadline(t time.Time) func(o *Option) {
	return func(o *Option) {
		o.transportDialContextDeadline = t
	}
}

func WithForceAttemptHTTP2(t bool) func(o *Option) {
	return func(o *Option) {
		o.transportForceAttemptHTTP2 = t
	}
}

func WithMaxIdleConns(t int) func(o *Option) {
	return func(o *Option) {
		o.transportMaxIdleConns = t
	}
}

func WithIdleConnTimeout(t time.Duration) func(o *Option) {
	return func(o *Option) {
		o.transportIdleConnTimeout = t
	}
}

func WithTLSHandshakeTimeout(t time.Duration) func(o *Option) {
	return func(o *Option) {
		o.transportTLSHandshakeTimeout = t
	}
}

func WithResponseHeaderTimeout(t time.Duration) func(o *Option) {
	return func(o *Option) {
		o.transportResponseHeaderTimeout = t
	}
}

func WithExpectContinueTimeout(t time.Duration) func(o *Option) {
	return func(o *Option) {
		o.transportExpectContinueTimeout = t
	}
}

func WithWholeRequestTimeout(t time.Duration) func(o *Option) {
	return func(o *Option) {
		o.clientTimeout = t
	}
}

// NewClient returns a http.Client configured with reasonable defaults for all
// timeouts. The default standard library http.Client has no timeouts set,
// potentially causing requests to hang indefinitely.
func NewClient(options ...func(*Option)) *http.Client {
	op := Option{
		transportDialContextTimeout:    30 * time.Second,
		transportDialContextKeepAlive:  30 * time.Second,
		transportDialContextDeadline:   time.Time{},
		transportForceAttemptHTTP2:     true,
		transportMaxIdleConns:          100,
		transportIdleConnTimeout:       90 * time.Second,
		transportTLSHandshakeTimeout:   5 * time.Second,
		transportResponseHeaderTimeout: 10 * time.Second,
		transportExpectContinueTimeout: 1 * time.Second,
		clientTimeout:                  60 * time.Second,
	}

	for _, o := range options {
		o(&op)
	}

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			// Timeout for establishing a new connection
			Timeout:   op.transportDialContextTimeout,
			KeepAlive: op.transportDialContextKeepAlive,
			// Deadline is the absolute point in time after which dials
			// will fail
			Deadline: op.transportDialContextDeadline,
		}).DialContext,
		ForceAttemptHTTP2: op.transportForceAttemptHTTP2,
		MaxIdleConns:      op.transportMaxIdleConns,
		IdleConnTimeout:   op.transportIdleConnTimeout,
		// Timeout for the TLS handshake
		TLSHandshakeTimeout: op.transportTLSHandshakeTimeout,
		// Timeout for reading the headers of the response
		ResponseHeaderTimeout: op.transportResponseHeaderTimeout,
		ExpectContinueTimeout: op.transportExpectContinueTimeout,
	}

	client := &http.Client{
		Timeout:   op.clientTimeout, // Timeout of the whole request
		Transport: transport,
	}

	return client
}

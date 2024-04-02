package apigateway

import (
	"encoding/base64"
	"html/template"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"

	httputil "github.com/Vitality-South/goutil/http"
)

// Common HTTP error response body content
const (
	Error400Content = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Bad Request</title></head><body><h1>Bad Request</h1><p>The server could not parse, decode or understand the input.</p></body></html>`
	Error401Content = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Unauthorized</title></head><body><h1>Unauthorized</h1><p>The server could not authenticate the request.</p></body></html>`
	Error403Content = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Forbidden</title></head><body><h1>Forbidden</h1><p>The server will not allow access to this resource.</p></body></html>`
	Error404Content = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Page Not Found</title></head><body><h1>Page Not Found</h1><p>Sorry, but the page you were trying to view does not exist.</p></body></html>`
	Error405Content = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Method Not Allowed</title></head><body><h1>Method Not Allowed</h1><p>The server will not allow this HTTP request method.</p></body></html>`
	Error410Content = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Gone</title></head><body><h1>Gone</h1><p>This resource has been permanently deleted from the server.</p></body></html>`
	Error413Content = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Payload Too Large</title></head><body><h1>Payload Too Large</h1><p>The input is too large.</p></body></html>`
	Error429Content = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Too Many Requests</title></head><body><h1>Too Many Requests</h1><p>The user has sent too many requests in a given amount of time ("rate limiting"). Please try again.</p></body></html>`

	Error500Content = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Internal Server Error</title></head><body><h1>Internal Server Error</h1><p>The server has encountered a situation it doesn't know how to handle.</p></body></html>`
	Error502Content = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Bad Gateway</title></head><body><h1>Bad Gateway</h1><p>The server has encountered an invalid response from a backend server. Please try again later.</p></body></html>`
	Error503Content = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Service Unavailable</title></head><body><h1>Service Unavailable</h1><p>The server is temporarily unavailable. Please try again later.</p></body></html>`
)

// ProxyResponse returns an API Gateway Proxy HTTP response with the
// provided status, headers, and body.
//
// The HTTP response body is always base64 encoded and the appropriate flag is
// set for API Gateway to handle it. If using REST APIs, make sure to set the
// Binary Media Types to "*/*". This is not necessary for HTTP APIs.
//
// If headers is nil or empty, DefaultHTTPHeaders() will be used.
func ProxyResponse(status int, headers http.Header, body []byte) events.APIGatewayProxyResponse {
	// determine HTTP headers to use; use DefaultHTTPHeaders if user input is
	// nil or empty
	h := func() http.Header {
		if len(headers) == 0 {
			return httputil.DefaultHTTPHeaders()
		}

		return headers
	}()

	return events.APIGatewayProxyResponse{
		Body:              base64.StdEncoding.EncodeToString(body),
		StatusCode:        status,
		MultiValueHeaders: h,
		IsBase64Encoded:   true,
	}
}

// V2HTTPResponse returns an API Gateway HTTP response with the
// provided status, headers, and body.
//
// The HTTP response body is always base64 encoded and the appropriate flag is
// set for API Gateway to handle it. If using REST APIs, make sure to set the
// Binary Media Types to "*/*". This is not necessary for HTTP APIs.
//
// If headers is nil or empty, DefaultHTTPHeaders() will be used.
func V2HTTPResponse(status int, headers http.Header, body []byte) events.APIGatewayV2HTTPResponse {
	// determine HTTP headers to use; use DefaultHTTPHeaders if user input is
	// nil or empty
	h := func() http.Header {
		if len(headers) == 0 {
			return httputil.DefaultHTTPHeaders()
		}

		return headers
	}()

	return events.APIGatewayV2HTTPResponse{
		Body:              base64.StdEncoding.EncodeToString(body),
		StatusCode:        status,
		MultiValueHeaders: h,
		IsBase64Encoded:   true,
	}
}

// RedirectBody returns a useful html body for 30x response status codes.
func RedirectBody(url string) (string, error) {
	html := `<!doctype html><html lang="en"><head><meta charset="utf-8">
	<title>Redirecting</title></head>
	<body><h1>Redirecting</h1>
	<p>Redirecting to <a href="{{.Url}}">{{.Url}}</a></p>
	</body>
	</html>`

	tmpl, err := template.New("RedirectBody").Parse(html)
	if err != nil {
		return "", err
	}

	data := struct {
		Url string
	}{
		Url: url,
	}

	var sb strings.Builder

	err = tmpl.Execute(&sb, data)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

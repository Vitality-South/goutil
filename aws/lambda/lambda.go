package lambda

import (
	"encoding/base64"

	"github.com/aws/aws-lambda-go/events"
)

// WebsocketAPILambdaResponse is a type suitable for use as a return type in a
// Lambda function for API Gateway Websocket API.
//
// The API Gateway Websocket API expects a 200 series status code on success
// and non-200 series status (such as 40x or 50x) for errors.
type WebsocketAPIResponse struct {
	StatusCode int `json:"statusCode"`
}

// RequestBody returns the body and properly base64 decodes it if necessary.
func RequestBody(isBase64Encoded bool, body string) (string, error) {
	if isBase64Encoded {
		data, err := base64.StdEncoding.DecodeString(body)

		if err != nil {
			return body, err
		}

		return string(data), nil
	}

	return body, nil
}

// RequestBodyFromApiGatewayWebsocket returns the body and properly base64
// decodes it if necessary.
func RequestBodyFromApiGatewayWebsocket(request *events.APIGatewayWebsocketProxyRequest) (string, error) {
	return RequestBody(request.IsBase64Encoded, request.Body)
}

// RequestBodyFromApiGatewayV2HTTP returns the body and properly base64
// decodes it if necessary.
func RequestBodyFromFromApiGatewayV2HTTP(request *events.APIGatewayV2HTTPRequest) (string, error) {
	return RequestBody(request.IsBase64Encoded, request.Body)
}

// RequestBodyFromFromAPIGatewayProxy returns the body and properly base64
// decodes it if necessary.
func RequestBodyFromFromAPIGatewayProxy(request *events.APIGatewayProxyRequest) (string, error) {
	return RequestBody(request.IsBase64Encoded, request.Body)
}

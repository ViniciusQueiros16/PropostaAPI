package response

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func ApiResponse(status int, body interface{}) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Headers": "*",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
		},
		StatusCode: status,
	}

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return resp, nil
}

package main

import "context"

type HTTPParameters struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type WebEvent struct {
	Parameters HTTPParameters `json:"http"`
}

type ResponseHeaders struct {
	ContentType string `json:"Content-Type"`
}

type Response struct {
	Body       string          `json:"body"`
	StatusCode string          `json:"statusCode"`
	Headers    ResponseHeaders `json:"headers"`
}

func Main(ctx context.Context, evt WebEvent) Response {
	switch evt.Parameters.Method {
	case "GET":
		return handleGet()
	case "POST":
		return handlePost()
	}
	return handleError()
}

func handleGet() Response {
	// example 1x1 GIF
	gif := "R0lGODlhAQABAAD/ACwAAAAAAQABAAACADs="
	return Response{
		Body:       gif,
		StatusCode: "200",
		Headers: ResponseHeaders{
			ContentType: "image/gif",
		},
	}
}

func handlePost() Response {
	return handleError()
}

func handleError() Response {
	return Response{
		StatusCode: "400",
	}
}

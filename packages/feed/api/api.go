package main

type ResponseHeaders struct {
	ContentType string `json:"Content-Type"`
}

type Response struct {
	Body       string          `json:"body"`
	StatusCode string          `json:"statusCode"`
	Headers    ResponseHeaders `json:"headers"`
}

func Main() Response {
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

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type HTTPParameters struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type WebEvent struct {
	Parameters HTTPParameters `json:"http"`
}

type FeedItem struct {
	Created     string `json:"created"`
	DownloadURL string `json:"download"`
}

type FeedViewModel struct {
	UploadURL string     `json:"upload_url"`
	Items     []FeedItem `json:"items"`
}

type ResponseHeaders struct {
	ContentType string `json:"Content-Type"`
}

type Response struct {
	Body       FeedViewModel   `json:"body"`
	StatusCode string          `json:"statusCode"`
	Headers    ResponseHeaders `json:"headers"`
}

var (
	key, secret, bucket, region string
)

func init() {
	key = os.Getenv("SPACES_KEY")
	if key == "" {
		panic("no key provided")
	}
	secret = os.Getenv("SPACES_SECRET")
	if secret == "" {
		panic("no secret provided")
	}
	bucket = os.Getenv("BUCKET")
	if bucket == "" {
		panic("no bucket provided")
	}
	region = os.Getenv("REGION")
	if region == "" {
		panic("no bucket provided")
	}
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
	// Configure blob storage.
	config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String(fmt.Sprintf("%s.digitaloceanspaces.com:443", region)),
		Region:      aws.String("us-east-1"), // Must be us-east-1 for compatibility with Spaces.
	}
	sess, err := session.NewSession(config)
	if err != nil {
		return handleError()
	}

	// Sign the upload URL.
	uploadURL, err := uploadURL(sess, bucket, uuid.NewString(), time.Hour)
	if err != nil {
		return handleError()
	}

	// Retrieve existing feed items.
	// TODO

	return Response{
		Body: FeedViewModel{
			UploadURL: uploadURL,
		},
		StatusCode: "200",
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

func uploadURL(sess *session.Session, bucket string, filename string, duration time.Duration) (string, error) {
	client := s3.New(sess)
	req, _ := client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	url, err := req.Presign(duration)
	if err != nil {
		return "", err
	}
	return url, nil
}

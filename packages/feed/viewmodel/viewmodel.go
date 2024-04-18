package main

import (
	"fmt"
	"log"
	"net/http"
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

type FeedViewModel struct {
	UploadURL    string   `json:"upload_url"`
	DownloadURLs []string `json:"download_urls"`
}

type Response struct {
	StatusCode *int               `json:"statusCode,omitempty"`
	Headers    *map[string]string `json:"headers,omitempty"`
	Body       *FeedViewModel     `json:"body,omitempty"`
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
		panic("no region provided")
	}
}

func Main(evt WebEvent) Response {
	switch evt.Parameters.Method {
	case "GET":
		return handleGet()
	}
	return handleError(fmt.Errorf("unexpected request method %q", evt.Parameters.Method))
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
		return handleError(err)
	}
	client := s3.New(sess)

	// Sign the upload URL.
	uploadURL, err := presignUploadURL(client, bucket, uuid.NewString(), time.Hour)
	if err != nil {
		return handleError(err)
	}

	// Retrieve existing feed items.
	downloadURLs := []string{}
	out, err := client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		handleError(err)
	}
	for _, obj := range out.Contents {
		req, _ := client.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(*obj.Key)},
		)
		url, err := req.Presign(time.Hour)
		if err != nil {
			handleError(err)
		}
		downloadURLs = append(downloadURLs, url)
	}

	return Response{
		Body: &FeedViewModel{
			UploadURL:    uploadURL,
			DownloadURLs: downloadURLs,
		},
	}
}

func handleError(err error) Response {
	log.Println(err.Error())
	statusCode := http.StatusBadRequest
	return Response{
		StatusCode: &statusCode,
	}
}

func presignUploadURL(client *s3.S3, bucket string, key string, duration time.Duration) (string, error) {
	req, _ := client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	url, err := req.Presign(duration)
	if err != nil {
		return "", err
	}
	return url, nil
}

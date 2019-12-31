package s3

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"grpc-poc/controller/vendors/http"
)

func client() *s3.S3 {
	config := &aws.Config{Region: aws.String("sa-east-1")}
	sess := session.Must(session.NewSession(config))

	s3Client := s3.New(sess)

	return s3Client
}

// io.Writer
func Download(bucket string, file string) (*int64, error) {
	downloader := s3manager.NewDownloaderWithClient(client())

	f, err := os.Create("audio_temp.mp3")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create file %s, %v", "audio_temp.mp3", err))
	}

	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(file),
	})

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to download file, %v", err))
	}

	fmt.Print(n)
	return &n, nil
}

func Stream(from string, bucket string, to string, acl string, contentType string) error {
	uploader := s3manager.NewUploaderWithClient(client())

	requestedRecord, err := http.Request("GET", from)

	if err != nil {
		return err
	}

	upParams := &s3manager.UploadInput{
		Bucket:          aws.String(bucket),
		Key:             aws.String(to),
		Body:            *requestedRecord,
		ACL:             aws.String(acl),
		ContentEncoding: aws.String(contentType),
	}

	_, err = uploader.Upload(upParams)

	return err
}

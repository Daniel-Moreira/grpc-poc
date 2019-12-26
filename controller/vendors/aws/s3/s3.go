package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3manager"
)

func client() *s3.S3 {
	config := &aws.Config{Region: aws.String("sa-east-1")}
	sess := session.Must(session.NewSession(config))

	s3Client := s3.New(sess)

	return s3Client
}

func Stream(from, bucket, to, acl, contentType) error {
	uploader := s3manager.NewUploaderWithClient(client())

  requestedRecord := requestRecord(from)

	upParams := &s3manager.UploadInput{
    Bucket: aws.String(bucket),
    Key: aws.String(to),
    Body: requestedRecord,
    ACL: aws.String(acl),
    ContentEncoding: aws.String(contentType)
  }

	uploader.Upload(upParams)

	return nil
}

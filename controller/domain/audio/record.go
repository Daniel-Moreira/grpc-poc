package audio

import (
	"errors"
	"fmt"
	"grpc-poc/controller/vendors/aws/s3"
	"grpc-poc/rpc"
)

func GetRecord(id int32) (bool, error) {
	bucket := "recordings-backup"
	filePath := fmt.Sprintf("teste/%d.mp3", id)

	_, err := s3.Download(bucket, filePath)

	if err != nil {
		return false, err
	}

	return true, nil
}

func BackupRecord(record *rpc.Record) (bool, error) {
	if record.GetStatus() != rpc.Record_completed {
		return false, errors.New("audio not finished")
	}

	toBucket := "recordings-backup"
	toPath := fmt.Sprintf("teste/%d.mp3", record.GetId())
	fromUri := fmt.Sprintf("%s.mp3", record.GetUrl())
	acl := "public-read"
	contentType := "audio/mpeg"

	err := s3.Stream(fromUri, toBucket, toPath, acl, contentType)

	if err != nil {
		return false, err
	}

	return true, nil
}

package audio

import (
	"errors"
	"grpc-poc/rpc"
	"grpc-poc/controller/vendors/aws/s3"
)

func GetRecord(id int32) (bool, error) {
  bucket := "recordings-backup"
  file := fmt.Sprintf("teste/%d.mp3", id)

  file, err := s3.Download(bucket, file)

  if err != nil {
    false, err
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
    false, err
  }

	return true, nil
}

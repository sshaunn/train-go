package config

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func init() {
	fmt.Println("init config")
}

type S3Service struct {
	Session *session.Session
}

const (
	BUCKET_NAME = "shaun-golang-bucket"
	REGION      = "us-east-1"
)

func (s *S3Service) NewS3Service() *S3Service {
	s.Session, _ = session.NewSession(&aws.Config{Region: aws.String(REGION)})
	return s
}

func (s *S3Service) Save(filename string, f multipart.File) error {
	_, err := s3.New(s.Session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(filename),
		Body:   f,
	})
	fmt.Println(filename)
	return err
}

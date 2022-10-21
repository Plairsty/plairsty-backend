package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func S3Uploader(apiId, apiKey, region string) *s3manager.Uploader {
	cred := credentials.NewStaticCredentials(
		apiId, apiKey, "",
	)
	s3session := session.Must(session.NewSession(
		&aws.Config{
			Region:      aws.String(region),
			Credentials: cred,
		}))
	return s3manager.NewUploader(s3session)
}

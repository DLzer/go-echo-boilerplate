package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// DO S3 Client Constructor
func NewS3Client(accessKey string, accessSecret string, s3Endpoint string, s3Region string) (*s3.S3, error) {
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
		Endpoint:    aws.String(s3Endpoint),
		Region:      aws.String(s3Region),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}

	return s3.New(newSession), nil
}

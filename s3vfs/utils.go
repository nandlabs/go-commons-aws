package s3vfs

import (
	"bytes"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

type UrlOpts struct {
	Host   string
	Bucket string
	Key    string
}

func parseUrl(url *url.URL) (*UrlOpts, error) {

	err := validateUrl(url)
	if err != nil {
		return nil, err
	}

	host := url.Host
	pathParams := strings.Split(url.Path, "/")
	bucket := pathParams[0]
	var b bytes.Buffer
	for _, item := range pathParams[1:] {
		b.WriteString("/")
		b.WriteString(item)
	}
	key := b.String()
	return &UrlOpts{
		Host:   host,
		Bucket: bucket,
		Key:    key,
	}, nil
}

func (urlOpts *UrlOpts) CreateS3Service() (*s3.S3, error) {
	awsSession, err := GetSession(urlOpts.Host, urlOpts.Bucket)
	if err != nil {
		return nil, err
	}
	svc := s3.New(awsSession)
	return svc, nil
}

func keyExists(bucket, key string, svc *s3.S3) (bool, error) {
	_, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "NotFound": // s3.ErrCodeNoSuchKey does not work, aws is missing this error code so we hardwire a string
				return false, nil
			default:
				return false, err
			}
		}
		return false, err
	}
	return true, nil
}

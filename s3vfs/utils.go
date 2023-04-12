package s3vfs

import (
	"bytes"
	"context"
	"errors"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
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

func validateUrl(u *url.URL) error {
	pathElements := strings.Split(u.Path, "/")
	if len(pathElements) == 1 {
		//Only Bucket provided
		return nil
	} else if len(pathElements) >= 2 {
		//Bucket and object path provided
		return nil
	} else { //path elements==0
		//return error as it's not a valid url with bucket missing
		return errors.New("invalid url with bucket missing")
	}
}

func (urlOpts *UrlOpts) CreateS3Client() (client *s3.Client, err error) {
	var awsSession *aws.Config

	awsSession, err = GetSession(urlOpts.Host, urlOpts.Bucket)
	if err != nil {
		return
	}
	client = s3.NewFromConfig(*awsSession)
	return
}

func keyExists(bucket, key string, svc *s3.Client) (bool, error) {
	_, err := svc.HeadObject(context.Background(), &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		var nsk *types.NoSuchKey
		if errors.As(err, &nsk) {
			// handle NoSuchKey error
			return false, err
		}
	}
	return true, nil
}

func getS3Object(url *url.URL) (result *s3.GetObjectOutput, err error) {
	var urlOpts *UrlOpts
	var svc *s3.Client

	urlOpts, err = parseUrl(url)
	svc, err = urlOpts.CreateS3Client()
	if err != nil {
		return
	}
	result, err = svc.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(urlOpts.Bucket),
		Key:    aws.String(urlOpts.Key),
	})
	return
}

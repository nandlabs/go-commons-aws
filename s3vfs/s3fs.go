package s3vfs

import (
	"context"
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.nandlabs.io/commons/vfs"
)

const (
	fileScheme = "s3"
)

var localFsSchemes = []string{fileScheme}

type S3Fs struct {
	*vfs.BaseVFS
}

// Create : creating a file in the s3 bucket, can create both object and bucket
func (o *S3Fs) Create(u *url.URL) (file vfs.VFile, err error) {
	var urlOpts *UrlOpts
	var svc *s3.Client
	var found bool

	urlOpts, err = parseUrl(u)
	if err != nil {
		return
	}
	svc, err = urlOpts.CreateS3Service()
	if err != nil {
		return
	}
	// check if the same path already exist on the s3 or not
	found, err = keyExists(urlOpts.Bucket, urlOpts.Key, svc)
	if !found {
		return
	}

	// create the folder structure or an empty file
	_, err = svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(urlOpts.Bucket),
		Key:    aws.String(urlOpts.Key),
	})
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}
	return
}

// Open location provided of the S3 bucket
func (o *S3Fs) Open(u *url.URL) (file vfs.VFile, err error) {
	var urlOpts *UrlOpts
	var svc *s3.Client

	urlOpts, err = parseUrl(u)
	if err != nil {
		return
	}
	svc, err = urlOpts.CreateS3Service()
	if err != nil {
		return
	}

	_, err = svc.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(urlOpts.Bucket),
		Key:    aws.String(urlOpts.Key),
	})
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
	if err == nil {
		file = &S3File{
			Location: u,
			fs:       o,
		}
	}
	return
}

func (o *S3Fs) Schemes() []string {
	return localFsSchemes
}

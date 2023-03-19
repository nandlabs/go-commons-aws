package s3vfs

import (
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.nandlabs.io/commons/vfs"
)

const (
	fileScheme = "s3"
)

var localFsSchemes = []string{fileScheme}

type S3Fs struct {
	*vfs.BaseVFS
}

// creating a file in the s3 bucket, can create both object and bucket
func (o *S3Fs) Create(u *url.URL) (file vfs.VFile, err error) {
	urlOpts, err := parseUrl(u)
	if err != nil {
		return nil, err
	}
	svc, err := urlOpts.CreateS3Service()
	if err != nil {
		return nil, err
	}
	// check if the same path already exist on the s3 or not
	found, existError := keyExists(urlOpts.Bucket, urlOpts.Key, svc)
	if !found {
		return nil, existError
	}

	// create the folder structure or an empty file
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(urlOpts.Bucket),
		Key:    aws.String(urlOpts.Key),
	})
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return nil, err
	}
	return
}

// s3://abc/11/test/file.txt
// abc -> abc
// 11/test -> folder hierarchy
func (o *S3Fs) Open(u *url.URL) (file vfs.VFile, err error) {
	urlOpts, err := parseUrl(u)
	if err != nil {
		return nil, err
	}
	svc, err := urlOpts.CreateS3Service()
	if err != nil {
		return nil, err
	}

	resp, openError := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(urlOpts.Bucket),
		Key:    aws.String(urlOpts.Key),
	})
	if openError != nil {
		fmt.Println("Error downloading file:", openError)
		return nil, openError
	}
	if err == nil {
		file = &S3File{
			file:     resp,
			Location: u,
			fs:       o,
		}
	}
	return
}

func (o *S3Fs) Schemes() []string {
	return localFsSchemes
}

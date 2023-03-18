package s3vfs

import (
	"fmt"
	"net/url"
	"strings"

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
	// TODO: repetitive code, can be moved to a generic function
	err = validateUrl(u)
	if err != nil {
		return nil, err
	}

	awsSession, err := GetSession(u.Host, u.Path)
	if err != nil {
		return nil, err
	}
	svc := s3.New(awsSession)

	pathParams := strings.Split(u.Path, "/")
	bucket := pathParams[0]
	key := parseKeyFromPath(pathParams)
	// check if the same path already exist on the s3 or not
	found, existError := keyExists(bucket, key, svc)
	if !found {
		return nil, existError
	}

	// create the folder structure or an empty file
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
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

	pathParams := strings.Split(u.Path, "/")
	bucket := pathParams[0]
	key := parseKeyFromPath(pathParams)

	awsSession, err := GetSession(u.Host, u.Path)
	if err != nil {
		return nil, err
	}
	svc := s3.New(awsSession)

	resp, openError := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
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

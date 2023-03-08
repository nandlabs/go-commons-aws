package s3vfs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	go_commons_aws "go.nandlabs.io/commons-aws"
	"go.nandlabs.io/commons/vfs"
)

const (
	fileScheme = "s3"
)

var (
	svc = s3.New(go_commons_aws.AwsSession)
)

var localFsSchemes = []string{fileScheme}

type S3Fs struct {
	*vfs.BaseVFS
}

func (o S3Fs) Create(u *url.URL) (file vfs.VFile, err error) {
	fileBytes, readErr := os.ReadFile(u.Path)
	if readErr != nil {
		fmt.Println("Error reading file:", readErr)
		return nil, readErr
	}
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("BUCKET"),
		Key:    aws.String(u.Path),
		Body:   bytes.NewReader(fileBytes),
	})
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return nil, err
	}
	return
}

func (o S3Fs) Open(u *url.URL) (file vfs.VFile, err error) {
	resp, openError := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("BUCKET"),
		Key:    aws.String(u.Path),
	})
	if openError != nil {
		fmt.Println("Error downloading file:", openError)
		return nil, openError
	}
	defer resp.Body.Close()

	// the logic can be improved to read the response file
	fileBytes, readError := ioutil.ReadAll(resp.Body)
	if readError != nil {
		fmt.Println("Error reading file contents:", readError)
		return nil, readError
	}

	tempFile, _ := ioutil.TempFile("", "example")
	_, _ = tempFile.Write(fileBytes)

	var f *os.File
	f, err = os.Open(tempFile.Name())
	if err == nil {
		file = &S3File{
			file:     f,
			Location: u,
			fs:       o,
		}
	}
	return
}

func (o S3Fs) Schemes() []string {
	return localFsSchemes
}

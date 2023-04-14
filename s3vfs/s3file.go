package s3vfs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"go.nandlabs.io/commons/vfs"
)

type S3File struct {
	*vfs.BaseFile
	Location *url.URL
	fs       vfs.VFileSystem
	closers  []io.Closer
}

// Read - s3Object read the body
func (s3File *S3File) Read(b []byte) (body int, err error) {
	var result *s3.GetObjectOutput

	result, err = getS3Object(s3File.Location)
	s3File.closers = append(s3File.closers, result.Body)
	defer s3File.Close()
	return result.Body.Read(b)
}

func (s3File *S3File) Write(b []byte) (numBytes int, err error) {
	var urlOpts *UrlOpts
	var svc *s3.Client

	urlOpts, err = parseUrl(s3File.Location)
	if err != nil {
		numBytes = 0
		return
	}
	svc, err = urlOpts.CreateS3Client()
	if err != nil {
		numBytes = 0
		return
	}

	// if key exists in s3 then the key will be overwritten else the new key with input body is created
	_, err = svc.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(urlOpts.Bucket),
		Key:    aws.String(urlOpts.Key),
		Body:   bytes.NewReader(b),
	})
	if err != nil {
		fmt.Println("Error writing file:", err)
		numBytes = 0
		return
	}
	numBytes = len(b)
	return
}

func (s3File *S3File) ListAll() (files []vfs.VFile, err error) {
	var urlOpts *UrlOpts
	var svc *s3.Client
	var result *s3.ListObjectsV2Output

	urlOpts, err = parseUrl(s3File.Location)
	if err != nil {
		return
	}
	svc, err = urlOpts.CreateS3Client()
	if err != nil {
		return
	}

	result, err = svc.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
		Bucket: aws.String(urlOpts.Bucket),
	})
	var contents []types.Object
	if err != nil {
		log.Printf("Couldn't list objects in bucket %v. Here's why: %v\n", urlOpts.Bucket, err)
	} else {
		contents = result.Contents
	}

	for _, item := range contents {
		u, _ := url.Parse(*item.Key)
		files = append(files, &S3File{
			Location: u,
		})
	}
	return
}

func (s3File *S3File) Info() (file vfs.VFileInfo, err error) {
	var result *s3.GetObjectOutput

	result, err = getS3Object(s3File.Location)
	s3File.closers = append(s3File.closers, result.Body)
	defer s3File.Close()
	file = &s3FileInfo{
		key:          result.Metadata["key"],
		size:         result.ContentLength,
		lastModified: *result.LastModified,
	}
	return
}

func (s3File *S3File) AddProperty(name, value string) (err error) {
	var urlOpts *UrlOpts
	var svc *s3.Client

	urlOpts, err = parseUrl(s3File.Location)
	if err != nil {
		return
	}
	svc, err = urlOpts.CreateS3Client()
	if err != nil {
		return
	}
	// Create an input object for the CopyObject API operation.
	copyInput := &s3.CopyObjectInput{
		Bucket:            aws.String(urlOpts.Bucket),
		CopySource:        aws.String(fmt.Sprintf("%s/%s", urlOpts.Bucket, urlOpts.Key)),
		Key:               aws.String(urlOpts.Key),
		MetadataDirective: "REPLACE",
		Metadata: map[string]string{
			name: value,
		},
	}
	// Call the CopyObject API operation to create a copy of the object with the new metadata.
	_, err = svc.CopyObject(context.Background(), copyInput)
	if err != nil {
		return
	}
	return
}

func (s3File *S3File) GetProperty(name string) (value string, err error) {
	var urlOpts *UrlOpts
	var svc *s3.Client
	var result *s3.HeadObjectOutput

	urlOpts, err = parseUrl(s3File.Location)
	if err != nil {
		return
	}
	svc, err = urlOpts.CreateS3Client()
	if err != nil {
		return
	}
	// Create an input object for the HeadObject API operation.
	input := &s3.HeadObjectInput{
		Bucket: aws.String(urlOpts.Bucket),
		Key:    aws.String(urlOpts.Key),
	}
	// Call the HeadObject API operation to retrieve the object metadata.
	result, err = svc.HeadObject(context.Background(), input)
	if err != nil {
		return
	}
	value = result.Metadata[name]

	return
}

func (s3File *S3File) Url() *url.URL {
	return s3File.Location
}

func (s3File *S3File) Delete() (err error) {
	var urlOpts *UrlOpts
	var svc *s3.Client
	var result *s3.DeleteObjectOutput

	urlOpts, err = parseUrl(s3File.Location)
	if err != nil {
		return
	}
	svc, err = urlOpts.CreateS3Client()
	if err != nil {
		return
	}

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(urlOpts.Bucket),
		Key:    aws.String(urlOpts.Key),
	}

	result, err = svc.DeleteObject(context.Background(), input)
	if err != nil {
		return
	}
	logger.Info(result)
	return
}

func (s3File *S3File) Close() (err error) {
	if len(s3File.closers) > 0 {
		for _, closable := range s3File.closers {
			err = closable.Close()
		}
	}
	return
}

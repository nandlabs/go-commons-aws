package s3vfs

import (
	"net/url"

	"github.com/aws/aws-sdk-go/service/s3"
	"go.nandlabs.io/commons/vfs"
)

type S3File struct {
	*vfs.BaseFile
	file     *s3.GetObjectOutput // this file will be the s3Object instead of the os.File
	Location *url.URL
	fs       vfs.VFileSystem
}

// Read - s3Object read the body
func (s3File *S3File) Read(b []byte) (int, error) {}

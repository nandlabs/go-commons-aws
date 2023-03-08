package s3vfs

import (
	"go.nandlabs.io/commons/vfs"
	"net/url"
	"os"
)

type S3File struct {
	*vfs.BaseFile
	file     *os.File
	Location *url.URL
	fs       vfs.VFileSystem
}

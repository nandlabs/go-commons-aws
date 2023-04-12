package s3vfs

import (
	"os"
	"time"
)

type s3FileInfo struct {
	key          string
	size         int64
	lastModified time.Time
}

func (f *s3FileInfo) Name() string {
	return f.key
}

func (f *s3FileInfo) Size() int64 {
	return f.size
}

func (f *s3FileInfo) Mode() os.FileMode {
	// Not applicable for S3 Objects, return default value
	return 0
}
func (f *s3FileInfo) ModTime() time.Time {
	return f.lastModified
}

func (f *s3FileInfo) IsDir() bool {
	// Not applicable for S3 Objects, return default value
	return false
}

func (f *s3FileInfo) Sys() interface{} {
	// Not applicable for S3 Objects, return default value
	return nil
}

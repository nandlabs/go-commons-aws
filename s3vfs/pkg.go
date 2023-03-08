package s3vfs

import "go.nandlabs.io/commons/vfs"

func init() {
	s3Fs := &S3Fs{}
	vfs.Register(s3Fs)
}

# Virtual File System (VFS) for S3

VFS for S3 allows you to abstract away the underlying file system, and provide a uniform interface for accessing files and directories, regardless of where they are physically located.

---
- [Installation](#installation)
- [Features](#features)
- [Usage](#usage)
---

### Installation

```bash
go get go.nandlabs.io/commons-aws/s3vfs
```

### Features

// TODO

### Usage

1. Register your provider
    ```go
    package main
    
    import (
        "context"
    
        "github.com/aws/aws-sdk-go-v2/aws"
        "github.com/aws/aws-sdk-go-v2/config"
        "go.nandlabs.io/commons-aws/s3vfs"
    )
    
    type S3SessionProvider struct {
        region string
        bucket string
    }
    
    func (s3SessionProvider *S3SessionProvider) Get() (*aws.Config, error) {
        sess, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(s3SessionProvider.region))
        return &sess, err
    }
    
    func main() {
        pvd := &S3SessionProvider{
            region: "us-east-1",
            bucket: "dummy",
        }
        s3vfs.AddSessionProvider(pvd.region, pvd.bucket, pvd)
    }
    ```

2. Create a bucket/file in S3
3. Read a file from S3
4. Delete a file in S3
5. Write a file in S3
6. List all the files in S3 bucket
7. Get File Info of an S3 object
8. Get metadata of an S3 object
9. Add metadata to an S3 object
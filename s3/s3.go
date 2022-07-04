package s3

import "io"

type S3 interface {
	GetObject(key string) (io.ReadCloser, error)
	PutObject(key string, out io.WriteCloser) error
	DeleteObjects(key ...string) error
	// ListObject(key string) ([]ObjectEntry, error)
}

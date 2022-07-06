package dfs

import "io"

type DfsOptions struct {
	ContentType string
	// Content-Disposition: attachment; filename="filename.jpg"
	FileName      string
	Headers       map[string]string
	ContentLength int64
	AllowOrigin   string
	AllowMethod   string
}

type Stat struct {
	Exist         bool
	ContentType   string
	FileName      string
	ContentLength int64
	Headers       map[string]string
	AllowOrigin   string
	AllowMethod   string
}

type Dfs interface {
	Stat(key string) (Stat, error)
	Get(key string) (io.ReadCloser, Stat, error)
	Put(key string, out io.Reader, options DfsOptions) error
	Delete(keys ...string) error
	GetBytes(key string) ([]byte, Stat, error)
	PutBytes(key string, out []byte, options DfsOptions) error
	Close() error
}

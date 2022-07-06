package hwobs

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/daqiancode/gocommons/dfs"
	"github.com/daqiancode/gocommons/dfs/hwobs/obs"
)

type OBS struct {
	client *obs.ObsClient
	bucket string
}

func NewOBS(accessKey, secretKey, endpoint, bucketName string) (*OBS, error) {
	obsClient, err := obs.New(accessKey, secretKey, endpoint)
	if err != nil {
		return nil, err
	}
	r := &OBS{
		client: obsClient,
		bucket: bucketName,
	}
	return r, nil
}

func (s *OBS) cleanKey(key string) string {
	return strings.TrimLeft(filepath.Clean(key), "/")

}

func (s *OBS) Stat(key string) (dfs.Stat, error) {
	inp := &obs.GetObjectMetadataInput{}
	inp.Bucket = s.bucket
	inp.Key = s.cleanKey(key)
	out, err := s.client.GetObjectMetadata(inp)
	var r dfs.Stat
	r.Exist = true
	if err != nil {
		if IsNotFountError(err) {
			r.Exist = false
			return r, nil
		}
		return r, err
	}

	r.ContentType = out.ContentType
	r.Headers = out.Metadata
	r.ContentLength = out.ContentLength
	r.AllowOrigin = out.AllowOrigin
	r.AllowMethod = out.AllowMethod
	r.FileName = s.getFileNameFromHeader(out.Metadata)
	return r, nil
}

func IsNotFountError(err error) bool {
	if e, ok := err.(obs.ObsError); ok {
		fmt.Println(e.StatusCode)
		return e.StatusCode == 404
	}
	return false
}

func (s *OBS) getFileNameFromHeader(header map[string]string) string {
	if header == nil {
		return ""
	}
	if v, ok := header["content-disposition"]; ok {
		m := regexp.MustCompile("filename=\"(.*?)\"").FindStringSubmatch(v)
		if len(m) > 1 {
			return m[1]
		}
	}
	return ""
}
func (s *OBS) Get(key string) (io.ReadCloser, dfs.Stat, error) {
	inp := &obs.GetObjectInput{}
	inp.Bucket = s.bucket
	inp.Key = s.cleanKey(key)
	out, err := s.client.GetObject(inp)
	var stat dfs.Stat
	stat.Exist = true
	if err != nil {
		if IsNotFountError(err) {
			stat.Exist = false
			return nil, stat, nil
		}
		return nil, stat, err
	}

	stat.AllowMethod = out.AllowMethod
	stat.AllowOrigin = out.AllowOrigin
	stat.ContentLength = out.ContentLength
	stat.ContentType = out.ContentType
	stat.FileName = s.getFileNameFromHeader(out.Metadata)
	stat.Headers = out.Metadata
	return out.Body, stat, nil
}

func (s *OBS) Put(key string, data io.Reader, options dfs.DfsOptions) error {
	inp := &obs.PutObjectInput{}
	inp.Bucket = s.bucket
	inp.Key = s.cleanKey(key)
	inp.Body = data
	inp.ContentType = options.ContentType

	headers := make(map[string]string)
	for k, v := range options.Headers {
		headers[k] = v
	}
	if options.ContentType != "" {
		headers["ContentType"] = options.ContentType
	}
	if options.FileName != "" {
		headers["Content-Disposition"] = "attachment; filename=\"" + options.FileName + "\""
	}
	if options.ContentLength > 0 {
		headers["ContentLength"] = strconv.FormatInt(options.ContentLength, 10)
	}
	inp.Metadata = headers
	_, err := s.client.PutObject(inp)
	if err != nil {
		return err
	}
	// if out.StatusCode >= 200 && out.StatusCode < 300 {
	// 	return errors.New("obs put object failed, http response status code: " + strconv.Itoa(out.StatusCode))
	// }
	return nil
}
func (s *OBS) Delete(keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	inp := &obs.DeleteObjectsInput{}
	inp.Bucket = s.bucket
	inp.Objects = make([]obs.ObjectToDelete, len(keys))
	for i, key := range keys {
		inp.Objects[i] = obs.ObjectToDelete{Key: key}
	}
	_, err := s.client.DeleteObjects(inp)
	return err
}

func (s *OBS) GetBytes(key string) ([]byte, dfs.Stat, error) {
	src, stat, err := s.Get(key)
	if err != nil {
		return nil, stat, err
	}
	defer src.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, src)
	if err != nil {
		return nil, stat, err
	}
	return buf.Bytes(), stat, nil
}
func (s *OBS) PutBytes(key string, out []byte, options dfs.DfsOptions) error {
	buf := bytes.NewBuffer(out)
	return s.Put(key, buf, options)
}

func (s *OBS) Close() error {
	s.client.Close()
	return nil
}

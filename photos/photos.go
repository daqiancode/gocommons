package photos

import (
	"bytes"
	"io"
	"path/filepath"
	"strconv"

	"github.com/daqiancode/gocommons/dfs"
	"github.com/daqiancode/gocommons/utils/imageutils"
	"github.com/daqiancode/gocommons/utils/pathutils"
)

type Photos struct {
	fs        dfs.Dfs
	PathMaker func(ext string) string
}

func NewPhotos(fs dfs.Dfs) (*Photos, error) {
	return &Photos{
		fs:        fs,
		PathMaker: pathutils.MakeDateRandPath,
	}, nil
}

type UploadPhotoResult struct {
	Raw     string
	Cropped string
	Resized []string
}

//Upload upload image & size, return [raw , resizes images]
func (s *Photos) Upload(file io.ReadSeeker, root, fileName string, contentType string, crop []int, sizes [][]int) (UploadPhotoResult, error) {
	var r UploadPhotoResult
	var err error

	//1. raw
	_, err = file.Seek(0, 0)
	if err != nil {
		return r, err
	}
	rawBuffer := bytes.NewBuffer(nil)
	_, err = io.Copy(rawBuffer, file)
	if err != nil {
		return r, err
	}
	//2. crop
	var cropBuffer *bytes.Buffer
	_, err = file.Seek(0, 0)
	if err != nil {
		return r, err
	}
	var t *imageutils.ImageTransform
	if len(crop) == 4 {
		t, err = imageutils.NewImageTransformExt(file, fileName)
		if err != nil {
			return r, err
		}
		t = t.Crop4(crop)

		cropBuffer, err = t.Buffer()
		if err != nil {
			return r, err
		}
	}
	resizedBuffers := make([]*bytes.Buffer, len(sizes))
	//3. resize
	for i, size := range sizes {
		_, err = file.Seek(0, 0)
		if err != nil {
			return r, err
		}
		if cropBuffer != nil {
			t, err = imageutils.NewImageTransformExt(bytes.NewBuffer(cropBuffer.Bytes()), fileName)

		} else {
			t, err = imageutils.NewImageTransformExt(bytes.NewBuffer(rawBuffer.Bytes()), fileName)
		}
		if err != nil {
			return r, err
		}
		t.ResizeMax(size[0], size[1])
		resizedBuffers[i], err = t.Buffer()
		if err != nil {
			return r, err
		}
	}

	//4. upload to obs
	p := filepath.Join(root, s.PathMaker(filepath.Ext(fileName)))
	opts := dfs.DfsOptions{ContentType: contentType}
	r.Raw = p
	err = s.fs.Put(r.Raw, rawBuffer, opts)
	if err != nil {
		return r, err
	}

	if cropBuffer != nil {
		r.Cropped = pathutils.TagFilename(p, "0")
		err = s.fs.Put(r.Cropped, cropBuffer, opts)
		if err != nil {
			return r, err
		}
	}

	r.Resized = make([]string, len(sizes))
	for i := range sizes {
		r.Resized[i] = pathutils.TagFilename(p, strconv.Itoa(i+1))
		err = s.fs.Put(r.Resized[i], resizedBuffers[i], opts)
		if err != nil {
			return r, err
		}
	}
	return r, nil

}

func (s *Photos) Resize(file io.Reader, fileName string, size []int) (*bytes.Buffer, error) {
	t, err := imageutils.NewImageTransformExt(file, fileName)
	if err != nil {
		return nil, err
	}
	t.ResizeMax(size[0], size[1])
	return t.Buffer()
}

func (s *Photos) Delete(keys ...string) error {
	return s.fs.Delete(keys...)
}

func (s *Photos) Close() error {
	return s.fs.Close()
}

package photos_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/daqiancode/gocommons/dfs/hwobs"
	"github.com/daqiancode/gocommons/env"
	"github.com/daqiancode/gocommons/photos"
	"github.com/stretchr/testify/assert"
)

func TestPhotos(t *testing.T) {
	fs, err := hwobs.NewOBS(env.GetenvMust("obs_access_token"), env.GetenvMust("obs_secret_token"), env.GetenvMust("obs_endpoint"), env.GetenvMust("obs_bucket"))
	assert.Nil(t, err)
	photos := photos.NewPhotos(fs)
	defer photos.Close()

	file, err := os.OpenFile(filepath.Join(env.Getwd(), "test.jpg"), os.O_RDONLY, 0644)
	assert.Nil(t, err)
	defer file.Close()
	r, err := photos.Upload(file, "photos", "test.jpg", "image/jpeg", []int{10, 10, 100, 100}, [][]int{{300, 200}, {100, 100}, {40, 20}})
	fmt.Println(err)
	assert.Nil(t, err)
	fmt.Println(r)
	r1, err := photos.Upload(file, "photos", "test.jpg", "image/jpeg", nil, [][]int{{300, 200}, {100, 100}, {40, 20}})
	fmt.Println(err)
	assert.Nil(t, err)
	fmt.Println(r1)
	r2, err := photos.Upload(file, "photos", "test.jpg", "image/jpeg", nil, nil)
	fmt.Println(err)
	assert.Nil(t, err)
	fmt.Println(r2)
	err = photos.Delete(r.Raw, r.Cropped)
	assert.Nil(t, err)
	err = photos.Delete(r.Resized...)
	assert.Nil(t, err)
	err = photos.Delete(r1.Raw, r1.Cropped)
	assert.Nil(t, err)
	err = photos.Delete(r1.Resized...)
	assert.Nil(t, err)
	err = photos.Delete(r2.Raw, r2.Cropped)
	assert.Nil(t, err)
	err = photos.Delete(r2.Resized...)
	assert.Nil(t, err)

}

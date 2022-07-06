package hwobs_test

import (
	"testing"

	"github.com/daqiancode/gocommons/dfs"
	"github.com/daqiancode/gocommons/dfs/hwobs"
	"github.com/daqiancode/gocommons/env"
	"github.com/stretchr/testify/assert"
)

func TestObs(t *testing.T) {
	ak := env.Getenv("hwobs_ak")
	sk := env.Getenv("hwobs_sk")
	endpoint := env.Getenv("hwobs_endpoint")
	bucket := env.Getenv("hwobs_bucket")

	fs, err := hwobs.NewOBS(ak, sk, endpoint, bucket)
	defer fs.Close()
	assert.Nil(t, err)
	test(t, fs)

}

func test(t *testing.T, d dfs.Dfs) {
	key := "/test/hello.txt"
	content := "hello world"
	filename := "file.txt"
	err := d.PutBytes(key, []byte(content), dfs.DfsOptions{FileName: filename})
	assert.Nil(t, err)
	_, err = d.Stat(key)
	assert.Nil(t, err)
	stat, err := d.Stat(key + "_not_exist")
	assert.Nil(t, err)
	assert.False(t, stat.Exist)

	bs, stat1, err := d.GetBytes(key)
	assert.Nil(t, err)
	assert.Equal(t, content, string(bs))
	assert.True(t, stat1.Exist)
	assert.Equal(t, stat1.FileName, filename)

}

package requests_test

import (
	"fmt"
	"testing"

	"github.com/daqiancode/gocommons/requests"
	"github.com/stretchr/testify/assert"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func TestCaller(t *testing.T) {
	caller := requests.NewCaller("http://localhost:8080/api/iam/", nil, nil)
	var token Token
	_, _, err := caller.PostForm("/token", map[string]string{"name": "jim", "password": "123"}, &token)
	// fmt.Println(string(bs), resp, err)
	assert.Nil(t, err)
	fmt.Println(token.AccessToken)
	var r interface{}
	caller.SetBearer(token.AccessToken)
	_, resp, err := caller.Get("/protected", nil, &r)
	fmt.Println(resp.StatusCode)
	fmt.Println(r)
}
